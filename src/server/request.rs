use std::io::Read;

use mio::{net::TcpStream, Token};

// -------------------------------------------------------------------------------------
// REQUEST
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone)]
pub struct Request {
    pub id_session: String,
    pub location: String,
    pub host: String,
    pub port: u16,
    pub method: String,
    pub bytes: usize,
    pub body: String,
}

impl Request {
    pub fn new(
        id_session: String,
        location: String,
        host: String,
        port: u16,
        method: String,
        bytes: usize,
        body: String,
    ) -> Self {
        Self {
            id_session,
            location,
            host,
            port,
            method,
            bytes,
            body,
        }
    }
    /// Lit une requête HTTP à partir d'un TcpStream.
    pub fn read_request(stream: &mut TcpStream, token: Token) -> Self {
        // Lire les données du client
        let mut buffer = [0; 1024];
        let n = match stream.read(&mut buffer) {
            Ok(n) => n,
            Err(_) => 0,
        };

        if n == 0 {
            // Connexion fermée par le client et doit etre logger dans server.log
            eprintln!("Client({}) deconnecte", token.0);
        }

        // Convertir les données reçues en une chaîne de caractères
        let request_str = String::from_utf8_lossy(&buffer[..n]).to_string();

        // Parser la requête HTTP pour créer une instance de `Request`
        // On utilise `catch_unwind` pour capturer les paniques de `parse_http_request`
        Request::parse_http_request(&request_str, n)
    }
    /// Parse une requête HTTP et crée une instance de `Request`.
    pub fn parse_http_request(request_str: &str, n: usize) -> Self {
        let mut location = String::new();
        let mut host = String::new();
        let mut port: u16 = 0;
        let mut method = String::new();
        let mut body = String::new();

        // Diviser la requête en lignes
        let lines: Vec<&str> = request_str.lines().collect();

        // Parser la première ligne (ex: "GET /index.html HTTP/1.1")
        if lines.len() > 3 {
            let parts: Vec<&str> = lines[0].split_whitespace().collect();
            if parts.len() >= 2 {
                method = parts[0].to_string(); // Méthode (GET, POST, etc.)
                location = parts[1].to_string(); // URL (/index.html)
            }
            let raw_host = lines[1].strip_prefix("Host: ");
            if let Some(h) = raw_host {
                let host_parts: Vec<&str> = h.split(":").collect();
                host = host_parts[0].to_string();
                port = host_parts[1].parse::<u16>().unwrap();
            }
        }

        // Extraire les cookies des en-têtes
        let cookies = Self::extract_cookies(&lines);

        // Parser le corps de la requête (s'il existe)
        let mut is_body = false;
        for line in lines.iter().skip(1) {
            if line.is_empty() {
                // Une ligne vide sépare les en-têtes du corps
                is_body = true;
                continue;
            }
            if is_body {
                body.push_str(line);
                body.push('\n');
            }
        }

        // Créer une instance de `Request`
        Request::new(
            cookies,
            location,
            host,
            port,
            method,
            n,
            body,
        )
    }

    fn extract_cookies(headers: &[&str]) -> String {
        let mut cookies = String::new();

        for line in headers {
            if line.starts_with("Cookie:") {
                let cookie_str = line.trim_start_matches("Cookie:").trim();
                for cookie in cookie_str.split(';') {
                    let mut parts = cookie.trim().splitn(2, '=');
                    if let (Some(_), Some(value)) = (parts.next(), parts.next()) {
                        cookies= value.to_string();
                    }
                }
            }
        }

        cookies
    }
}
// -------------------------------------------------------------------------------------
