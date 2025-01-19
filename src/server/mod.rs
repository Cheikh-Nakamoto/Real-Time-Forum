pub mod request;

use std::{ fs, io::Write, path::Path };

use mio::net::TcpStream;
pub use request::*;
pub mod response;
pub use response::*;
pub mod router;
pub use router::*;
pub mod session;
use serde::{ Deserialize, Serialize };
pub use session::*;
pub mod cgi;
pub use cgi::*;
use tera::{ Context, Tera };

// -------------------------------------------------------------------------------------
// DIRECTORY ELEMENT
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DirectoryElement {
    pub entry: String,
    pub link: String,
    pub is_directory: bool,
}
// -------------------------------------------------------------------------------------

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

    pub fn handle_request(&self, mut stream: &mut TcpStream, request: Request) {
        println!("On veut aller dans {}", request.location);
        println!("Le répertoire de base de ce serveur est {}", self.root_directory.clone());

        // Chemin réel du fichier
        let location = self.root_directory.clone() + &request.location;

        let discover = fs::read_dir(&location);

        if discover.is_err() {
            let output = "RESSOURCE NON TROUVÉE";
            Self::send_error_response(stream, 404, &output);
            return;
        }

        let entries = discover.unwrap();
        let all: Vec<DirectoryElement> = entries
            .map(|entry| {
                let el = entry.unwrap().path();
                let name = el.to_str().unwrap().strip_prefix(&location).unwrap().to_string();

                DirectoryElement {
                    entry: name.clone(),
                    link: request.location.clone() + &name,
                    is_directory: el.is_dir(),
                }
            })
            .collect::<Vec<DirectoryElement>>();

        println!("Contenu du répertoire : {:#?}", all);

        let path = "./src/static_files/index.html"; // Chemin vers le fichier statique
        println!("if path exist {}", Path::new(&path).exists());
        println!("Chemin vérifié : {}", path);
        if Path::new(&path).exists() {
            // Servir un fichier statique
            self.handle_static_file(&mut stream, &path, all);
            println!("Handle static function");
        } else {
            // Ressource introuvable
            println!("Handle static function error");
            Self::send_error_response(&mut stream, 404, "Not Found");
        }
    }

    /// Gère une requête pour un fichier statique.
    fn handle_static_file(&self, stream: &mut TcpStream, path: &str, all: Vec<DirectoryElement>) {
        // Chargement du template
        let tera = Tera::new("src/**/*.html").unwrap();
        let mut context = Context::new();
        context.insert("elements", &all);

        match tera.render(&path.strip_prefix("./src/").unwrap(), &context) {
            Ok(content) => {
                let response = format!(
                    "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: {}\r\n\r\n{}",
                    content.len(),
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
