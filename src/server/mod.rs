pub mod request;
use std::{ fs, io::Write, path::Path };
use std::fs::ReadDir;
pub use std::string::String;
use mio::net::TcpStream;
use mio::Token;
pub use request::*;
pub mod response;
pub use response::*;

pub mod router;
pub use router::*;
pub mod session;
use tera::{ Context, Tera };
pub use session::*;
pub mod cgi;
pub mod rendering_page;

pub use cgi::*;
pub use rendering_page::*;

// -------------------------------------------------------------------------------------
// SERVER
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone, Deserialize)]
pub struct Server {
    pub ip_addr: String,
    pub hostname: String,
    pub ports: Vec<u16>,
    pub root_directory: String,
    pub error_path: String,
    pub default_file: String,
    pub access_log: String,
    pub cgi_file_format: String,
    pub upload_limit: u32,
    pub accepted_methods: Vec<String>,
    pub directory_listing: bool,
}

impl Server {
    pub fn new(
        ip_addr: String,
        hostname: String,
        ports: Vec<u16>,
        root_directory: String,
        error_path: String,
        default_file: String,
        access_log: String,
        cgi_file_format: String,
        upload_limit: u32,
        accepted_methods: Vec<String>,
        directory_listing: bool
    ) -> Self {
        Self {
            ip_addr,
            hostname,
            ports,
            root_directory,
            error_path,
            default_file,
            access_log,
            cgi_file_format,
            upload_limit,
            accepted_methods,
            directory_listing,
        }
    }

    pub fn handle_request(&self, mut stream: &mut TcpStream, request: Request,cookie:String) {
        let mut location_path = "".to_string();
        // Chemin réel du fichier
        let location = self.root_directory.clone() + &request.location;

        let discover = fs::read_dir(&location);
        let mut entries: ReadDir;
        let mut all: Vec<DirectoryElement> = vec![];
        let mut dir_path = "".to_string();
        if !request.location.contains(".") {
            location_path = "/index.html".to_string();
            dir_path = "src/static_files".to_string();
        } else {
            location_path = Self::check_and_clean_path(&request.location);
            dir_path = self.root_directory.clone();
        }
        if location_path.contains("/image") || location_path.contains("/css") {
            dir_path = "src/static_files".to_string();
        }
        let path = format!("./{}{}", dir_path, location_path); // Chemin relatif au dossier public

        if !discover.is_err() {
            entries = discover.unwrap();
            all = entries
                .map(|entry| {
                    let el = entry.unwrap().path();
                    let name = el.to_str().unwrap().strip_prefix(&location).unwrap().to_string();

                    let mut entry_name = name.clone();
                    if let Some(val) = entry_name.strip_prefix("/") {
                        entry_name = val.to_string();
                    }

                    DirectoryElement {
                        entry: entry_name.clone(),
                        entry_type: match el.is_dir() {
                            true => "folder".to_string(),
                            _ => match entry_name.strip_suffix(".rb") {
                                Some(_) => "ruby".to_string(),
                                None => "file".to_string()
                            }
                        },
                        link: request.location.clone() + &name,
                        is_directory: el.is_dir(),
                    }
                })
                .collect::<Vec<DirectoryElement>>();
            println!("Contenu du répertoire : {:#?}", all);
            self.handle_listing_directory(&mut stream, &path, all,cookie);
            return;
        }
        if Path::new(&path).exists() {
            // Servir un fichier statique
            self.handle_static_file(&mut stream, &path,cookie);
            println!("Handle static function");
        } else {
            // Ressource introuvable
            println!("Handle static function error");
            Self::send_error_response(&mut stream, 404, "Not Found");
        }
    }

    fn handle_static_file(&self, stream: &mut TcpStream, path: &str,cookie:String) {
        // Déterminer le type de contenu en fonction de l'extension du fichier
        let mut to_cgi = false;
        let content_type = match
            Path::new(path)
                .extension()
                .and_then(|ext| ext.to_str())
        {
            Some("html") => "text/html",
            Some("css") => "text/css",
            Some("js") => "application/javascript",
            Some("png") => "image/png",
            Some("jpg") | Some("jpeg") => "image/jpeg",
            Some("gif") => "image/gif",
            Some("json") => "application/json",
            Some("rb") => {
                to_cgi = true;
                "text/plain"
            }
            _ => "text/plain", // Type par défaut
        };

        // Lire le fichier
        match fs::read(path) {
            Ok(mut content) => {
                if to_cgi {
                    content = CGI::execute_file(path.to_string()).into();
                }

                let response = format!(
                    "HTTP/1.1 200 OK\r\nContent-Type: {}\r\nContent-Length: {}\r\n{}\r\n",
                    content_type,
                    content.len(),
                    cookie
                );

                if let Err(e) = stream.write_all(response.as_bytes()) {
                    eprintln!("Erreur lors de l'envoi de l'en-tête : {}", e);
                }
                if let Err(e) = stream.write_all(&content) {
                    eprintln!("Erreur lors de l'envoi du contenu : {}", e);
                }
            }
            Err(e) => {
                eprintln!("Erreur lors de la lecture du fichier : {}", e);
                Self::send_error_response(stream, 500, "Internal Server Error");
            }
        }
    }

    /// Gère une requête pour un fichier statique.
    fn handle_listing_directory(
        &self,
        stream: &mut TcpStream,
        path: &str,
        all: Vec<DirectoryElement>,
        cookie :String
    ) {
        // Chargement du template
        let tera = Tera::new("src/**/*.html").unwrap();
        let mut context = Context::new();
        context.insert("elements", &all);
        match tera.render(&path.strip_prefix("./src/").unwrap(), &context) {
            Ok(content) => {
                let response = format!(
                    "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: {}\r\n{}\r\n{}",
                    content.len(),
                    cookie,
                    content
                );

                if let Err(e) = stream.write_all(response.as_bytes()) {
                    eprintln!("Erreur lors de l'envoi de la réponse : {}", e);
                }
            }
            Err(e) => {
                eprintln!("Erreur lors de la lecture du fichier : {}", e);
                Self::send_error_response(stream, 500, "Internal Server Error");
            }
        }
    }

    /// Envoie une réponse d'erreur HTTP.
    fn send_error_response(stream: &mut TcpStream, status_code: u16, status_message: &str) {
        let response = format!(
            "HTTP/1.1 {} {}\r\nContent-Type: text/plain\r\nContent-Length: {}\r\n\r\n{}",
            status_code,
            status_message,
            status_message.len(),
            status_message
        );
        if let Err(e) = stream.write_all(response.as_bytes()) {
            eprintln!("Erreur lors de l'envoi de la réponse d'erreur : {}", e);
        } else {
            eprintln!("{}", status_message);
        }
    }


    fn check_and_clean_path(path: &str) -> String {
        // Trouver l'index du motif "images/" ou "css/"
        if let Some(index) = path.find("/images/").or_else(|| path.find("/css/")) {
            // Supprimer tout ce qui se trouve avant le motif
            let cleaned_path = &path[index..];
            cleaned_path.to_string()
        } else {
            // Retourner le chemin original si aucun motif n'est trouvé
            path.strip_prefix("/").unwrap().to_string()
        }
    }
    // pub fn access_log(&self, req: &Request) {
    //     let mut file = OpenOptions::new()
    //         .append(true)
    //         .create(true)
    //         .open(&self.access_log)
    //         .unwrap();

    //     writeln!(file, "{} {} {}", req.method, req.path, "200 OK").unwrap();
    // }
}
// -------------------------------------------------------------------------------------
