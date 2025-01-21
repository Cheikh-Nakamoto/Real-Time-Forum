use super::Request;
pub use super::{Server, Session};
use mio::net::{TcpListener, TcpStream};
use mio::{Events, Interest, Poll, Token};
use std::collections::HashMap;
use std::io::{self};
use std::net::ToSocketAddrs;

// -------------------------------------------------------------------------------------
// ROUTER
// -------------------------------------------------------------------------------------
const CLIENT_START: Token = Token(1000); // Token de départ pour les clients

#[derive(Debug)]
pub struct Router {
    pub servers: Vec<Server>,
    pub sessions: HashMap<Token, Session>,
    pub listeners: HashMap<Token, TcpListener>, // Associe un token à un TcpListener
    pub clients: HashMap<Token, TcpStream>,     // Associe un token à un TcpStream
    pub next_token: usize,
}

impl Router {
    pub fn new() -> Self {
        Self {
            servers: vec![],
            sessions: HashMap::new(),
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
                println!("Adresse de connexion {}",addr);
            let listener = TcpListener::bind(addr)?;
            let token = Token(self.next_token - 1000);
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
                    let stream = (self.clients.get_mut(&event.token()))
                        .expect("Erreur lors de la recupeartion du canal tcpstream");
                    let req = Request::read_request(stream, event.token());
                    println!("voila la requete {:?}", req);
                    let mut cookie = String::new();
                    println!("<=========================================>");
                    println!("all session : {:?}",self.sessions);
                    println!("<=========================================>");

                    if let Some(session) = self.sessions.get_mut(&event.token()) {
                        cookie = Session::make_cookie("cookie_01", &*session.id,session.expiration_time);
                    }
                        Self::route_request(self.servers.clone(), &req, stream,cookie);
                }
            }
        }
    }

    /// Accepte une nouvelle connexion et l'ajoute à la liste des clients.
    fn accept_connection(&mut self, token: Token, poll: &Poll) -> io::Result<()> {
        if let Some(listener) = self.listeners.get_mut(&token) {
            let (mut stream, _) = listener.accept()?;
            let client_token = Token(self.next_token);
            if let Some(cookie) =  Session::get_cookie_from_stream(&mut stream, "cookie_01"){
                for (old_token, session) in self.sessions.clone().iter() {
                    if session.id == cookie && !session.is_expired() {
                        let mut new_session = session.clone();
                        new_session.expiration_time = Session::SESSION_LIFETIME;
                        self.sessions.remove(&old_token);
                        self.sessions.insert(client_token.clone(), new_session);
                        println!("Connexion retablie avec le clien{:?} associé au nouveau token{:?}", old_token, client_token);
                    }
                }
            }else{
                let new_session = Session::new();
                self.sessions.insert(client_token.clone(), new_session.clone());
               //Session::send_cookie(&mut stream, "cookie_01", &*new_session.id, 1000);
                println!("Connexion etablie avec le nouveau client associé token{:?}",token);
            }
            self.next_token += 1;
            poll.registry()
                .register(&mut stream, client_token, Interest::READABLE)?;
            self.clients.insert(client_token, stream);
            println!("Nouveau client connecté avec le token: {:?}", client_token);
        }
        Ok(())
    }



    /// Gère les données d'un client existant.
    //fn handle_client(&mut self, token: Token) {}

    // Route une requête HTTP et génère une réponse.
    pub fn route_request(servers: Vec<Server>, req: &Request, stream: &mut TcpStream,cookie:String) {
        // On récupère le hostname, l'adresse ip et le port de la requête
        // On parcoure la liste des serveurs et on vérifie lequel a le hostname, le port et l'ip correspondant
        for server in servers.into_iter() {
            if server.ip_addr == req.host && server.ports.contains(&req.port) {
                println!("{} la resoudre", server.ip_addr);

                server.handle_request(stream, req.clone(),cookie.clone());
            }
        }
    }
}
