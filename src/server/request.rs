use std::{collections::HashMap, io::Read};
use mio::net::TcpStream;

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
    pub filename: String,
    pub length: usize,
    pub reference: String,
    pub headers: HashMap<String, String>,
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
        filename: String,
        length: usize,
        reference: String,
        headers: HashMap<String, String>,
    ) -> Self {
        Self {
            id_session,
            location,
            host,
            port,
            method,
            bytes,
            body,
            filename,
            length,
            reference,
            headers,
        }
    }

    /// Lit une requête HTTP à partir d'un TcpStream.
    pub fn read_request(stream: &mut TcpStream) -> Self {
        let mut buffer = [0; 8192]; // Buffer de 8 Ko
        let mut request_str = String::new();
        let mut headers_end = None;
        let mut content_length = 0;
        let mut is_delete = false;
        let mut byte_reader = 0;

        // Lire les données du client jusqu'à la fin des en-têtes
        loop {
            let n = match stream.read(&mut buffer) {
                Ok(n) => n,
                Err(_) => break, // Gestion des erreurs de lecture
            };
            if n == 0 {
                break; // Connexion fermée par le client
            }

            // Convertir les données reçues en une chaîne de caractères
            let buff = String::from_utf8_lossy(&buffer[..n]);
            if buff.starts_with("POST /DELETE") {
                is_delete = true;
            }
            request_str.push_str(&buff);
            byte_reader += n;

            // Vérifier si la fin des en-têtes a été atteinte
            if let Some(pos) = request_str.find("\r\n\r\n") {
                headers_end = Some(pos);

                if is_delete {
                    // Extraire la valeur de Content-Length
                    content_length = Self::extract_content_length(&request_str[..pos]);
                }
                break;
            }
        }

        // Extraire les en-têtes
        let headers_end = headers_end.unwrap_or_default();
        let headers = &request_str[..headers_end];

        // Parser les en-têtes pour créer une instance de `Request`
        let mut request = Request::parse_http_request(headers, headers_end, byte_reader);

        // Si Content-Length > 0, lire le corps de la requête
        if content_length > 0 {
            // Calculer la quantité de données déjà lues dans le corps
            let body_start = headers_end + 4; // 4 pour "\r\n\r\n"
            let body_already_read = request_str.len() - body_start;

            // Lire le reste du corps
            let mut body = vec![0; content_length];
            if body_already_read > 0 {
                // Copier les données déjà lues dans le corps
                body[..body_already_read].copy_from_slice(&request_str.as_bytes()[body_start..]);
            }

            // Lire les octets restants du corps avec stream.read_exact
            if body_already_read < content_length {
                stream
                    .read_exact(&mut body[body_already_read..])
                    .unwrap_or_default();
            }

            // Convertir le corps en String
            request.body = String::from_utf8(body).unwrap_or_default();
        }

        request
    }

    /// Extrait la valeur de l'en-tête Content-Length.
    fn extract_content_length(headers: &str) -> usize {
        headers
            .lines()
            .find(|line| line.starts_with("Content-Length:"))
            .and_then(|line| line.split(':').nth(1))
            .and_then(|s| s.trim().parse::<usize>().ok())
            .unwrap_or(0)
    }

    /// Parse une requête HTTP et crée une instance de `Request`.
    pub fn parse_http_request(request_str: &str, header_end: usize, n: usize) -> Self {
        let mut location = String::new();
        let mut host = String::new();
        let mut port: u16 = 0;
        let mut method = String::new();
        let mut filename = String::new();
        let mut body = String::new();
        let mut length = header_end;
        let mut headers = HashMap::new();

        // Diviser la requête en lignes
        let lines: Vec<&str> = request_str.lines().collect();

        // Parser la première ligne (ex: "GET /index.html HTTP/1.1")
        if !lines.is_empty() {
            let parts: Vec<&str> = lines[0].split_whitespace().collect();
            if parts.len() >= 2 {
                method = parts[0].to_string(); // Méthode (GET, POST, etc.)
                location = parts[1].to_string(); // URL (/index.html)
            }
        }

        // Parser les en-têtes
        for line in lines.iter().skip(1) {
            if line.starts_with("Host:") {
                let host_parts: Vec<&str> = line.split(":").collect();
                host = host_parts[1].trim().to_string();
                if host_parts.len() > 2 {
                    port = host_parts[2].parse::<u16>().unwrap_or(80);
                }
            } else if line.starts_with("Content-Disposition:") {
                // Extraire le nom du fichier
                filename = Self::extract_filename(line);
            } else if line.starts_with("Content-Length:") {
                // Extraire la taille du fichier
                length += line
                    .split(":")
                    .nth(1)
                    .and_then(|s| s.trim().parse::<usize>().ok())
                    .unwrap_or(0);
            } else if line.contains(":") {
                // Ajouter l'en-tête à la HashMap
                let mut parts = line.splitn(2, ":");
                if let (Some(key), Some(value)) = (parts.next(), parts.next()) {
                    headers.insert(key.trim().to_string(), value.trim().to_string());
                }
            }
        }

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

        let binding = Self::extract_header_value(&lines, "Referer:");
        let referer = binding.split(":").nth(1).unwrap_or_default();

        // Créer une instance de `Request`
        Request::new(
            String::new(), // id_session (à remplir plus tard)
            location,
            host,
            port,
            method,
            n,
            body,
            filename,
            length,
            referer.to_string(),
            headers,
        )
    }

    fn extract_filename(header: &str) -> String {
        let parts: Vec<&str> = header.split("filename=").collect();
        if parts.len() > 1 {
            parts[1]
                .trim_matches('"')
                .trim_matches(';')
                .trim()
                .to_string()
        } else {
            String::new()
        }
    }

    pub fn extract_header_value(headers: &[&str], pattern: &str) -> String {
        let mut header_value = String::new();

        for line in headers {
            if line.starts_with(pattern) {
                let cookie_str = line.trim_start_matches(pattern).trim();
                for cookie in cookie_str.split(';') {
                    let mut parts = cookie.trim().splitn(2, '=');
                    if let (Some(_), Some(value)) = (parts.next(), parts.next()) {
                        header_value = value.to_string();
                    }
                }
            }
        }
        header_value
    }
}