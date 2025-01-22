pub mod request;
use std::{ fs, io::Write, path::Path };
use std::fs::{ OpenOptions, ReadDir };
pub use std::string::String;
use chrono::Utc;
use mio::net::TcpStream;
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

use crate::{ remove_prefix, remove_suffix, Config, Redirection };

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
    pub cgi_file_format: String,
    pub upload_limit: u32,
    pub accepted_methods: Vec<String>,
    pub directory_listing: bool,
    pub redirections: Vec<Redirection>,
    pub exclusion: Vec<String>
}

impl Server {
    pub fn new(
        ip_addr: String,
        hostname: String,
        ports: Vec<u16>,
        root_directory: String,
        error_path: String,
        default_file: String,
        cgi_file_format: String,
        upload_limit: u32,
        accepted_methods: Vec<String>,
        directory_listing: bool,
        redirections: Vec<Redirection>,
        exclusion: Vec<String>,
    ) -> Self {
        Self {
            ip_addr,
            hostname,
            ports,
            root_directory,
            error_path,
            default_file,
            cgi_file_format,
            upload_limit,
            accepted_methods,
            directory_listing,
            redirections,
            exclusion
        }
    }

    pub fn access_log(&self, request: Request, config: &Config, status_code: u16, cookie: &String) {
        // Log request
        let mut tera = Tera::default();
        tera.add_raw_template("access_log", &config.http.access_log_format).unwrap();
        let mut context = Context::new();
        // "{{remote_addr}} - {{remote_user}} [{{time_local}}] - {{method}} - {{status}} {{bytes_sent}}"

        let id_session = if let Some(p1) = cookie.split(";").into_iter().next() {
            let parts = p1.split("=").collect::<Vec<&str>>();
            if parts.len() == 2 {
                parts[1]
            } else {
                ""
            }
        } else {
            ""
        };

        let addr = format!("{}:{}{}", &request.host, &request.port, &request.location);
        context.insert("remote_addr", &addr);
        context.insert("remote_user", id_session);
        context.insert("time_local", &format!("{}", Utc::now().format("%d-%m-%Y %H:%M:%S")));
        context.insert("method", &request.method);
        context.insert("status", &status_code);
        context.insert("bytes_sent", &((request.bytes as f64) / 1000.0));

        if let Ok(str) = tera.render("access_log", &context) {
            match OpenOptions::new().append(true).open(&config.log_files.access_log) {
                Ok(mut log_file) => {
                    let log_result = log_file.write((str + "\n").as_bytes());
                    match log_result {
                        Err(e) => eprintln!("Writing error. Err: {}", e),
                        Ok(_) => (),
                    }
                }
                Err(_) => (),
            }
        }
    }

    pub fn handle_request(
        &self,
        mut stream: &mut TcpStream,
        request: Request,
        cookie: String,
        config: &Config
    ) {
        // Vérification de la méthode
        if !self.accepted_methods.iter().any(|m| m.to_uppercase() == request.method.to_uppercase()) {
            Self::send_error_response(
                &self,
                &mut stream,
                request.clone(),
                config,
                405,
                "Method Not Allowed",
                &cookie
            );
            return;
        }

        let location_path;
        // Chemin réel du fichier
        let mut root = self.root_directory.clone();
        root = remove_suffix(root, "/");

        let location = "./".to_string() + &root + &request.location;

        let discover = fs::read_dir(&location);
        let entries: ReadDir;
        let all;
        let mut dir_path;

        if !request.location.contains(".") {
            if !Path::new(&location).exists() {
                Self::send_error_response(
                    &self,
                    &mut stream,
                    request.clone(),
                    config,
                    404,
                    "Not Found",
                    &cookie
                );
                return;
            }
            location_path = "/index.html".to_string();
            dir_path = "src/static_files".to_string();
        } else {
            location_path = Self::check_and_clean_path(&request.location);
            dir_path = self.root_directory.clone();
        }

        if location_path.contains("/image") || location_path.contains("/css") {
            dir_path = "src/static_files".to_string();
        }

        let path = format!(
            "./{}/{}",
            remove_suffix(dir_path, "/"),
            remove_prefix(location_path, "/")
        ); // Chemin relatif au dossier public

        if !discover.is_err() {
            entries = discover.unwrap();
            all = entries
                .map(|entry| {
                    let el = entry.unwrap().path();
                    let name = el.to_str().unwrap().strip_prefix(&location).unwrap().to_string();

                    let entry_name = remove_prefix(name.clone(), "/");

                    DirectoryElement {
                        entry: entry_name.clone(),
                        entry_type: match el.is_dir() {
                            true => "folder".to_string(),
                            _ =>
                                match entry_name.strip_suffix(".rb") {
                                    Some(_) => "ruby".to_string(),
                                    None => "file".to_string(),
                                }
                        },
                        link: request.location.clone() + &name,
                        is_directory: el.is_dir(),
                    }
                })
                .collect::<Vec<DirectoryElement>>();

            self.handle_listing_directory(&mut stream, all, cookie, request, config);
            return;
        }

        if Path::new(&path).exists() {
            // Servir un fichier statique
            self.handle_static_file(request, config, &mut stream, &path, cookie);
        } else {
            // Ressource introuvable
            Self::send_error_response(
                &self,
                &mut stream,
                request,
                config,
                404,
                "Not Found",
                &cookie
            );
        }
    }

    fn handle_static_file(
        &self,
        request: Request,
        config: &Config,
        stream: &mut TcpStream,
        path: &str,
        cookie: String
    ) {
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
                } else {
                    // Log request
                    self.access_log(request, config, 200, &cookie);
                }
                if let Err(e) = stream.write_all(&content) {
                    eprintln!("Erreur lors de l'envoi du contenu : {}", e);
                }
            }
            Err(e) => {
                eprintln!("Erreur lors de la lecture du fichier : {}", e);
                Self::send_error_response(
                    &self,
                    stream,
                    request,
                    config,
                    500,
                    "Internal Server Error",
                    &cookie
                );
            }
        }
    }

    /// Gère une requête pour un fichier statique.
    fn handle_listing_directory(
        &self,
        stream: &mut TcpStream,
        all: Vec<DirectoryElement>,
        cookie: String,
        request: Request,
        config: &Config
    ) {
        // Chargement du template
        let tera = Tera::new("src/**/*.html").unwrap();
        let mut context = Context::new();
        context.insert("elements", &all);
        context.insert("hostname", &self.hostname);

        match tera.render(&self.default_file.strip_prefix("src/").unwrap(), &context) {
            Ok(content) => {
                let response = format!(
                    "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\nContent-Length: {}\r\n{}\r\n{}",
                    content.len(),
                    cookie,
                    content
                );

                if let Err(e) = stream.write_all(response.as_bytes()) {
                    eprintln!("Erreur lors de l'envoi de la réponse : {}", e);
                } else {
                    // Log request
                    self.access_log(request, config, 200, &cookie);
                }
            }
            Err(e) => {
                eprintln!("Erreur lors de la lecture du fichier : {}", e);
                Self::send_error_response(
                    &self,
                    stream,
                    request,
                    config,
                    500,
                    "Internal Server Error",
                    &cookie
                );
            }
        }
    }

    /// Envoie une réponse d'erreur HTTP.
    fn send_error_response(
        &self,
        stream: &mut TcpStream,
        request: Request,
        config: &Config,
        status_code: u16,
        status_message: &str,
        cookie: &String
    ) {
        // Chargement du template
        let tera = Tera::new("src/**/*.html").unwrap();
        let mut context = Context::new();
        context.insert(
            "error",
            &(HTMLError { code: status_code, status: status_message.to_string() })
        );

        match tera.render(&self.error_path.strip_prefix("src/").unwrap(), &context) {
            Ok(content) => {
                let response = format!(
                    "HTTP/1.1 {} {}\r\nContent-Type: text/html\r\nContent-Length: {}\r\n\r\n{}",
                    status_code,
                    status_message,
                    content.len(),
                    content
                );
                if let Err(e) = stream.write_all(response.as_bytes()) {
                    eprintln!("Erreur lors de l'envoi de la réponse d'erreur : {}", e);
                } else {
                    self.access_log(request, config, status_code, &cookie);
                    eprintln!("{}", status_message);
                }
            }
            Err(e) => {
                eprintln!("{}", e);
            }
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
}
// -------------------------------------------------------------------------------------
