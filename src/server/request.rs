use mio::net::TcpStream;
use regex::Regex;
use std::{ collections::HashMap, io::{ BufReader, Read } };

use crate::{ get_boundary, remove_prefix, remove_suffix };

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
        headers: HashMap<String, String>
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

    pub fn read_request(stream: &mut TcpStream) -> Self {
        let mut buffer = [0; 8192]; // Buffer de 8 Ko
        let mut request_str = String::new();
        let mut headers_end = None;
        let mut is_post = false;
        let mut byte_reader = 0;
        let new_line_pattern = "\r\n\r\n";

        // ---------------------------------------------
        // Autre manière de lire à tester
        // ---------------------------------------------
        // let mut reader = BufReader::new(stream);
        // let mut data = Vec::new();
        // reader.read_to_end(&mut data);
        // request_str.push_str(&String::from_utf8_lossy(&data));
        // ---------------------------------------------

        loop {
            match stream.read(&mut buffer) {
                Ok(n) => {
                    let buff = String::from_utf8_lossy(&buffer[..n]);
                    if buff.starts_with("POST") {
                        is_post = true;
                    }
                    request_str.push_str(&buff);
                    byte_reader += n;

                    // if let Some(pos) = request_str.find(&new_line_pattern) {
                    //     headers_end = Some(pos);
                    // }
                }
                Err(_) => {
                    break;
                }
            }
        }

        // Vérification de la présence des 2 parties de la requête
        match request_str.find(new_line_pattern) {
            None => {
                eprintln!("Requête incomplète : fin des en-têtes non trouvée");
                return Request::new(
                    String::new(),
                    String::new(),
                    String::new(),
                    0,
                    String::new(),
                    0,
                    String::new(),
                    String::new(),
                    0,
                    String::new(),
                    HashMap::new()
                );
            }
            Some(header_limit) => {
                let headers_end = headers_end.unwrap();
                let headers = &request_str[..headers_end];

                let mut request = Request::parse_http_request(headers, headers_end, byte_reader);
                let mut form_data = vec![]; // Chaque HashMap représente un champ du formulaire.

                if is_post {
                    let mut head = request_str.clone();
                    let mut body = head.split_off(header_limit);
                    body = body.strip_prefix(new_line_pattern).unwrap().to_string();
                    let boundary = get_boundary(&request_str).unwrap();
                    let body_parts = body
                        .split(boundary.as_str())
                        .map(|s|
                            remove_suffix(remove_prefix(s.to_string(), "\r\n"), "\r\n--").replace(
                                new_line_pattern,
                                "; value="
                            )
                        )
                        .collect::<Vec<String>>();

                    // Tu peux jeter un coup d'œil sur la docu pour comprendre la syntaxe
                    // https://docs.rs/regex/latest/regex/index.html
                    let re = Regex::new(
                        r#"(?xs)
                    Content-Disposition:\s*
                    (?<content_disposition>[^;]+);\s*
                    name="(?<name>[^"]+)"\s*
                    (?:\s*;\s*
                        (?:filename="(?<filename>[^"]+)"\s*)?
                        (?:Content-Type:\s*(?<content_type>[^;]+)\s*)?
                    )*
                    ;\s*value=(?<value>.*)
                    "#
                    ).unwrap();

                    // Ici on parcourt les différentes parties du body pour voir si les champs recherchés sont là
                    body_parts.iter().for_each(|s| {
                        if let Some(caps) = re.captures(&s) {
                            let mut values = HashMap::new();
                            values.insert(
                                "content_disposition",
                                Some(caps["content_disposition"].to_string())
                            );
                            values.insert("name", Some(caps["name"].to_string()));
                            values.insert(
                                "filename",
                                caps.name("filename").map_or(None, |m| Some(m.as_str().to_string()))
                            );
                            values.insert(
                                "content_type",
                                caps
                                    .name("content_type")
                                    .map_or(None, |m| Some(m.as_str().to_string()))
                            );
                            values.insert("value", Some(caps["value"].to_string()));
                            form_data.push(values);
                        }
                    });

                    // A partir d'ici tu peux placer la fonction qui permet d'utiliser les données collectées
                    // Par exemple enregistrer l'image, pour le cas de la création de dossier tu auras ici
                    // Le nom du dossier à créer
                    // Tu sauras comment mettre à jour la variable request avec ces données collectées.
                    println!("{:#?}", form_data);
                }
                request
            }
        }

        // ----------------------------------------------------------------------------------------
        // Ancien bloc `if is_post {...}`
        // ----------------------------------------------------------------------------------------
        //     if is_post {
        //         let body_start = headers_end + 4;
        //         let body_already_not_read = request.length - body_start;

        //         let mut body = vec![0; request.length];

        //         if body_already_not_read > 0 {
        //             body = request_str.as_bytes()[body_start..].to_vec();
        //             if let Some(pos) = body.windows(4).position(|elem| elem == b"\r\n\r\n") {
        //                 // println!("trouvé dexieme delimiteur a l'index : {}", pos);
        //                 let sec_header = String::from_utf8_lossy(&body[..pos]).into_owned();
        //                 body = body[pos + 4..].to_vec();
        //                 let liste = sec_header.lines().collect::<Vec<&str>>();
        //                 let filename = Self::extract_header_value(&liste, "Content-Disposition");
        //                 println!("filename1 {} \n\n", filename);
        //                 request.filename = filename.replace('"', "").to_string();
        //             }
        //             if request.filename.is_empty() {
        //                 let liste = headers.lines().collect::<Vec<&str>>();
        //                 let filename = Self::extract_header_value(&liste, "Content-Disposition");
        //                 println!("filename2 {} \n\n", filename);
        //                 request.filename = filename.replace('"', "").to_string();
        //             }
        //         }

        //         let tmp = String::from_utf8_lossy(&body);
        //         // let liste = &tmp.to_owned().lines().collect::<Vec<&str>>();
        //         // let content_disposition = Self::extract_header_value(liste, "Content-Disposition");
        //         // println!("value {}", content_disposition);
        //         //    println!("body {} \n\n",tmp);

        //         request.body = tmp.to_string();
        //     }
        // ----------------------------------------------------------------------------------------
    }

    pub fn parse_http_request(request_str: &str, header_end: usize, n: usize) -> Self {
        let mut location = String::new();
        let mut host = String::new();
        let mut port: u16 = 0;
        let mut method = String::new();
        let body = String::new();
        let mut filename = String::new();
        let mut length = header_end;
        let mut headers = HashMap::new();

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
            } else if line.starts_with("Content-Length:") {
                // Extraire la taille du fichier
                length += line
                    .split(":")
                    .nth(1)
                    .and_then(|s| s.trim().parse::<usize>().ok())
                    .unwrap_or(0);
            } else if line.starts_with("Content-Length:") {
                // Extraire la taille du fichier
                length += line
                    .split(":")
                    .nth(1)
                    .and_then(|s| s.trim().parse::<usize>().ok())
                    .unwrap_or(0);
            } else if line.contains(":") {
                let mut parts = line.splitn(2, ":");
                if let (Some(key), Some(value)) = (parts.next(), parts.next()) {
                    let key = key.trim().trim_matches('"').to_string(); // Supprimer les espaces et les guillemets
                    let value = value.trim().to_string(); // Supprimer les espaces
                    if key == "Content-Disposition" {
                        filename = Self::extract_filename(&value);
                    }
                    // Ignorer les en-têtes vides
                    if !key.is_empty() && !value.is_empty() {
                        headers.insert(key, value);
                    }
                }
            }
        }

        let binding = Self::extract_header_value(&lines, "Referer:");
        let referer = binding.split(":").nth(1).unwrap_or_default();

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
            headers
        )
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

    fn extract_filename(content_disposition: &str) -> String {
        for segment in content_disposition.split(';') {
            let segment = segment.trim(); // Supprimer les espaces autour

            if segment.starts_with("filename=") {
                let filename = segment["filename=".len()..].trim();
                return filename.trim_matches('"').to_string();
            }
        }
        String::new()
    }
}
