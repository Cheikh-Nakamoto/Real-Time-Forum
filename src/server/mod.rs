pub mod request;

use std::{
    fs,
    io::{self, Read, Write},
    path::Path,
    process::{Command, Stdio},
};

use mio::net::TcpStream;
pub use request::*;
pub mod response;
pub use response::*;
pub mod router;
pub use router::*;
pub mod session;
use serde::Deserialize;
pub use session::*;
pub mod cgi;
pub use cgi::*;

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
        directory_listing: bool,
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

    pub fn handle_request(&self, mut stream:&mut TcpStream,request : Request) {
        // Vérifier si le chemin correspond à un fichier statique ou à un script CGI
        // Déterminer le chemin de la ressource demandée
        let location = if request.location == "/" {
            "/index.html".to_string()
        } else {
            request.location
        };
        let path = format!("./src/static_files{}", location); // Chemin relatif au dossier public
        println!("if path exist {}", Path::new(&path).exists() );
        println!("Chemin vérifié : {}", path);
        if Path::new(&path).exists() {
            // Servir un fichier statique
            self.handle_static_file(&mut stream, &path);
            println!("Handle static function");
        } else {
            // Ressource introuvable
            println!("Handle static function error");
            Self::send_error_response(&mut stream, 404, "Not Found");
        }
    }

    /// Gère une requête pour un fichier statique.
    fn handle_static_file(&self, stream: &mut TcpStream, path: &str) {
        let path = Path::new(path);

        // Déterminer le type de contenu en fonction de l'extension du fichier
        let content_type = match path.extension().and_then(|ext| ext.to_str()) {
            Some("html") => "text/html",
            Some("css") => "text/css",
            Some("js") => "application/javascript",
            Some("png") => "image/png",
            Some("jpg") | Some("jpeg") => "image/jpeg",
            Some("gif") => "image/gif",
            Some("json") => "application/json",
            _ => "text/plain", // Type par défaut
        };

        match fs::read(path) {
            Ok(content) => {
                let response = format!(
                    "HTTP/1.1 200 OK\r\nContent-Type: {}\r\nContent-Length: {}\r\n\r\n",
                    content_type,
                    content.len()
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
