use super::{Request, Response};
pub use super::{Server, Session};
use mio::net::{TcpListener, TcpStream};
use mio::{Events, Interest, Poll, Token};
use uuid::Uuid;
use std::collections::HashMap;
use std::io::{self, Read, Write};
use std::net::ToSocketAddrs;

// -------------------------------------------------------------------------------------
// ROUTER
// -------------------------------------------------------------------------------------
const CLIENT_START: Token = Token(1000); // Token de départ pour les clients

#[derive(Debug)]
pub struct Router {
    pub servers: Vec<Server>,
    pub sessions: Vec<Session>,
    pub listeners: HashMap<Token, TcpListener>, // Associe un token à un TcpListener
    pub clients: HashMap<Token, TcpStream>,     // Associe un token à un TcpStream
    pub next_token: usize,
}

impl Router {
    pub fn new() -> Self {
        Self {
            servers: vec![],
            sessions: vec![],
            listeners: HashMap::new(),
            clients: HashMap::new(),
            next_token: CLIENT_START.0,
        }
    }

    /// Ajoute un serveur et démarre l'écoute sur ses ports.
    pub fn add_server(&mut self, server: Server) -> io::Result<()> {
        for &port in &server.ports {
            let addr = format!("{}:{}", server.ip_addr, port)
                .to_socket_addrs()?
                .next()
                .ok_or_else(|| {
                    io::Error::new(io::ErrorKind::Other, "Impossible de résoudre l'adresse")
                })?;

            let listener = TcpListener::bind(addr)?;
            let token = Token(self.next_token-1000);
            self.next_token += 1;
            self.listeners.insert(token, listener);
        }
        self.servers.push(server);
        Ok(())
    }

    pub fn remove_server(&mut self, server: Server) -> io::Result<()> {
        // Filtrer les serveurs pour supprimer celui qui correspond
        self.servers
            .retain(|s| s.ip_addr != server.ip_addr && s.hostname != server.hostname);

        // Fermer les listeners associés à ce serveur
        for &port in &server.ports {
            // Trouver le token associé à ce port
            let token = self
                .listeners
                .iter()
                .find(|(_, listener)| {
                    listener.local_addr().ok().map(|addr| addr.port()) == Some(port)
                })
                .map(|(token, _)| *token);
            if let Some(token) = token {
                // Fermer le listener en le retirant de la HashMap
                self.listeners.remove(&token);
                println!("Listener fermé pour le port {}", port);
            }
        }

        Ok(())
    }

    pub fn add_session(&mut self, session: Session) {
        self.sessions.push(session);
    }
    pub fn remove_session(mut self, session_id: String) {
        self.sessions = self
            .sessions
            .into_iter()
            .filter(|s| s.id != session_id)
            .collect();
    }
    /// Démarre le Router et commence à écouter les événements.
    pub fn run(&mut self) -> io::Result<()> {
        let mut poll = Poll::new()?;
        let mut server_tokens = HashMap::new();

        // Enregistrer chaque listener avec un token unique
        for (token, listener) in &mut self.listeners {
            poll.registry()
                .register(listener, *token, Interest::READABLE)?;
            server_tokens.insert(*token, listener.local_addr()?);
        }

        let mut events = Events::with_capacity(128);
        loop {
            poll.poll(&mut events, None)?;

            for event in events.iter() {
                if let Some(&addr) = server_tokens.get(&event.token()) {
                    // Nouvelle connexion sur un TcpListener
                    self.accept_connection(event.token(), &poll)?;
                    println!("Nouvelle connexion sur le port {}", addr.port());
                } else {
                    // Données reçues sur un TcpStream
                    self.handle_client(event.token())?;
                    println!("Nouvelle requête")
                }
            }
        }
    }

    /// Accepte une nouvelle connexion et l'ajoute à la liste des clients.
    fn accept_connection(&mut self, token: Token, poll: &Poll) -> io::Result<()> {
        if let Some(listener) = self.listeners.get_mut(&token) {
            let (mut stream, _) = listener.accept()?;
            let client_token = Token(self.next_token);
            self.next_token += 1;
            poll.registry()
                .register(&mut stream, client_token, Interest::READABLE)?;
            self.clients.insert(client_token, stream);
            println!("Nouveau client connecté avec le token: {:?}", client_token);
        }
        Ok(())
    }

    /// Gère les données d'un client existant.
    fn handle_client(&mut self, token: Token) -> io::Result<()> {
        let request = {
            // Lire les données du client
            if let Some(stream) = self.clients.get_mut(&token) {
                let mut buffer = [0; 1024];
                let n = stream.read(&mut buffer)?;

                if n == 0 {
                    // Connexion fermée par le client
                    self.clients.remove(&token);
                    println!("Client déconnecté : {:?}", token);
                    return Ok(());
                }

                // Convertir les données reçues en une chaîne de caractères
                let request_str = String::from_utf8_lossy(&buffer[..n]).to_string();

                // Parser la requête HTTP pour créer une instance de `Request`
                self.parse_http_request(&request_str)
            } else {
                return Ok(());
            }
        };

        // Traiter la requête et générer une réponse
        //Creation et ajout de la session a la reponse si necessiare
        let response = Self::route_request(request);
        // Envoyer la réponse au client
        if let Some(stream) = self.clients.get_mut(&token) {
            Self::send_response(stream, response)?;
        }

        Ok(())
    }

    /// Parse une requête HTTP et crée une instance de `Request`.
    fn parse_http_request(&self, request_str: &str) -> Request {
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
        Request::new(Uuid::new_v4().to_string() ,location, host, port, method, body)
    }

    /// Envoie une réponse HTTP au client.
    fn send_response(stream: &mut TcpStream, response: Response) -> io::Result<()> {
        let response_str = response.to_http_response();
        stream.write_all(response_str.as_bytes())?;
        stream.flush()?;
        Ok(())
    }

    // Route une requête HTTP et génère une réponse.
    pub fn route_request(req: Request) -> Response {
        // On récupère le hostname, l'adresse ip et le port de la requête
        // On parcoure la liste des serveurs et on vérifie lequel a le hostname, le port et l'ip correspondant

        // Exemple de logique de routage
        if req.location == "/" {
            Response::new(
                String::new(), // id_session
                "HTTP/1.1 200 OK".to_string(),
                "text/html".to_string(),
                "<h1>Bienvenue !</h1>".to_string(),
            )
        } else {
            Response::new(
                String::new(), // id_session
                "HTTP/1.1 404 Not Found".to_string(),
                "text/html".to_string(),
                "<h1>Page non trouvée</h1>".to_string(),
            )
        }
    }
}
