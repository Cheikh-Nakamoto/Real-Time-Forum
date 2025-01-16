pub use mio::net::{ TcpListener, TcpStream };
pub use mio::{ Events, Interest, Poll, Token };
pub use std::collections::HashMap;
pub use std::io::{ Read, Write };
pub use std::str;

// Définir des tokens pour identifier les sockets
const SERVER: Token = Token(0); // Token pour le listener
const CLIENT_START: Token = Token(1); // Token de départ pour les clients

pub struct Client {
   pub listener: TcpListener,
   pub clients: HashMap<Token, TcpStream>,
   pub next_token: usize,
}

impl Client {
  pub  fn new(addr: &str) -> Self {
        let listener = TcpListener::bind(addr.parse().unwrap()).unwrap();
        Self {
            listener,
            clients: HashMap::new(),
            next_token: CLIENT_START.0,
        }
    }

    pub fn run(&mut self) -> std::io::Result<()> {
        let mut poll = Poll::new()?; // Utilisez `?` pour propager les erreurs
        let mut events = Events::with_capacity(1024);

        // Enregistrer le listener pour surveiller les nouvelles connexions
        poll.registry().register(&mut self.listener, SERVER, Interest::READABLE)?;

        println!("Server running on localhost:8080");

        loop {
            poll.poll(&mut events, None)?; // Utilisez `?` pour propager les erreurs

            for event in events.iter() {
                match event.token() {
                    SERVER => {
                        let (mut stream, _) = self.listener.accept()?; // Utilisez `?` pour propager les erreurs
                        println!("New connection");

                        let token = Token(self.next_token);
                        self.next_token += 1;

                        poll.registry().register(&mut stream, token, Interest::READABLE)?; // Utilisez `?` pour propager les erreurs

                        self.clients.insert(token, stream);
                    }
                    token => {
                        if let Some(mut stream) = self.clients.get_mut(&token) {
                            let mut buffer = [0; 1024];
                            match stream.read(&mut buffer) {
                                Ok(0) => {
                                    println!("Client disconnected");
                                    self.clients.remove(&token);
                                }
                                Ok(n) => {
                                    let request = str::from_utf8(&buffer[..n]).unwrap();
                                    println!("Received: {}", request);

                                    let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
                                    stream.write(response.as_bytes())?; // Utilisez `?` pour propager les erreurs
                                }
                                Err(e) => {
                                    eprintln!("Error reading from client: {}", e);
                                    self.clients.remove(&token);
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
