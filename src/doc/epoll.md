## **Qu'est-ce que `epoll` ?**

`epoll` est une interface du noyau Linux qui permet de surveiller plusieurs descripteurs de fichiers (sockets, fichiers, etc.) et d'être notifié lorsqu'ils sont prêts à être lus ou écrits. Il est plus performant que des mécanismes plus anciens comme `select` ou `poll`, surtout lorsqu'il y a un grand nombre de descripteurs de fichiers à surveiller.

### **Avantages de `epoll`**
1. **Scalabilité** : Capable de gérer des milliers de descripteurs de fichiers.
2. **Efficacité** : Ne nécessite pas de parcourir tous les descripteurs de fichiers à chaque appel.
3. **Faible latence** : Notifie immédiatement votre programme lorsqu'un événement se produit.

---

## **Utilisation de `epoll` en Rust avec `mio`**

`mio` est une crate Rust qui fournit une API de haut niveau pour `epoll` (ainsi que pour d'autres mécanismes comme `kqueue` sur macOS). Elle simplifie l'utilisation de `epoll` tout en restant performante.

### **Installation de `mio`**

Ajoutez `mio` à votre `Cargo.toml` :

```toml
[dependencies]
mio = "0.8"
```

---

## **Exemple de serveur TCP avec `epoll`**

Voici un exemple complet de serveur TCP utilisant `epoll` avec `mio` :

```rust
use mio::net::{TcpListener, TcpStream};
use mio::{Events, Interest, Poll, Token};
use std::collections::HashMap;
use std::io::{Read, Write};
use std::str;

// Définir des tokens pour identifier les sockets
const SERVER: Token = Token(0); // Token pour le listener
const CLIENT_START: Token = Token(1); // Token de départ pour les clients

struct Server {
    listener: TcpListener,
    clients: HashMap<Token, TcpStream>,
    next_token: usize,
}

impl Server {
    fn new(addr: &str) -> Self {
        let listener = TcpListener::bind(addr.parse().unwrap()).unwrap();
        Self {
            listener,
            clients: HashMap::new(),
            next_token: CLIENT_START.0,
        }
    }

    fn run(&mut self) {
        // Créer un Poll pour surveiller les événements
        let mut poll = Poll::new().unwrap();
        let mut events = Events::with_capacity(1024);

        // Enregistrer le listener pour surveiller les nouvelles connexions
        poll.registry()
            .register(&mut self.listener, SERVER, Interest::READABLE)
            .unwrap();

        println!("Server running on localhost:8080");

        // Boucle principale du serveur
        loop {
            // Attendre des événements
            poll.poll(&mut events, None).unwrap();

            // Traiter les événements
            for event in events.iter() {
                match event.token() {
                    SERVER => {
                        // Accepter une nouvelle connexion
                        let (mut stream, _) = self.listener.accept().unwrap();
                        println!("New connection");

                        // Générer un token pour le nouveau client
                        let token = Token(self.next_token);
                        self.next_token += 1;

                        // Enregistrer le client pour surveiller les événements
                        poll.registry()
                            .register(&mut stream, token, Interest::READABLE)
                            .unwrap();

                        // Ajouter le client à la HashMap
                        self.clients.insert(token, stream);
                    }
                    token => {
                        // Gérer les données reçues d'un client
                        if let Some(mut stream) = self.clients.get_mut(&token) {
                            let mut buffer = [0; 1024];
                            match stream.read(&mut buffer) {
                                Ok(0) => {
                                    // Connexion fermée par le client
                                    println!("Client disconnected");
                                    self.clients.remove(&token);
                                }
                                Ok(n) => {
                                    // Lire les données et renvoyer une réponse
                                    let request = str::from_utf8(&buffer[..n]).unwrap();
                                    println!("Received: {}", request);

                                    let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
                                    stream.write(response.as_bytes()).unwrap();
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

fn main() {
    let mut server = Server::new("127.0.0.1:8080");
    server.run();
}
```

---

### **Explication du code**

1. **Tokens** :
   - Chaque socket (client ou serveur) est identifié par un `Token`. Le `Token` est un identifiant unique utilisé par `mio` pour savoir quel socket a déclenché un événement.

2. **Poll** :
   - `Poll` est l'objet principal qui surveille les événements sur les sockets. Il utilise `epoll` en interne sur Linux.

3. **Enregistrement des sockets** :
   - Le `TcpListener` (le socket du serveur) est enregistré avec `Interest::READABLE` pour être notifié lorsqu'une nouvelle connexion est prête à être acceptée.
   - Les `TcpStream` (sockets des clients) sont enregistrés avec `Interest::READABLE` pour être notifiés lorsqu'il y a des données à lire.

4. **Boucle d'événements** :
   - La boucle principale appelle `poll.poll()` pour attendre des événements. Lorsqu'un événement se produit, il est traité en fonction du `Token`.

5. **Gestion des clients** :
   - Lorsqu'un client envoie des données, elles sont lues dans un buffer, et une réponse simple est renvoyée.
   - Si un client se déconnecte, son socket est retiré de la `HashMap`.

---

### **Fonctionnement de `epoll` avec `mio`**

1. **Création du `Poll`** :
   - `Poll::new()` initialise l'instance de `Poll`, qui utilise `epoll` en interne.

2. **Enregistrement des sockets** :
   - `poll.registry().register()` enregistre un socket avec un `Token` et un ensemble d'intérêts (`Interest::READABLE`, `Interest::WRITABLE`, etc.).

3. **Attente des événements** :
   - `poll.poll(&mut events, None)` attend que des événements se produisent sur les sockets enregistrés. Les événements sont stockés dans `events`.

4. **Traitement des événements** :
   - Chaque événement contient un `Token` qui identifie le socket concerné. Vous pouvez ensuite lire ou écrire des données sur ce socket.

---

### **Avantages de cette approche**

1. **Scalabilité** :
   - Un seul thread peut gérer des milliers de connexions simultanément.

2. **Efficacité** :
   - `epoll` notifie uniquement votre programme lorsqu'un événement se produit, évitant ainsi le busy-waiting.

3. **Simplicité** :
   - `mio` abstrait la complexité de `epoll` tout en restant performant.

---

### **Limites**

1. **Linux uniquement** :
   - `epoll` est spécifique à Linux. Si vous avez besoin de portabilité, utilisez une crate comme `tokio` qui fonctionne sur plusieurs systèmes d'exploitation.

2. **Complexité accrue** :
   - Comparé à un serveur multi-thread, la gestion des connexions dans un seul thread peut être plus complexe à déboguer.

---

### **Alternative : `tokio`**

Si vous préférez une approche plus haut niveau et portable, vous pouvez utiliser `tokio`, une crate asynchrone qui utilise `epoll` en interne sur Linux. Voici un exemple simple avec `tokio` :

```rust
use tokio::net::TcpListener;
use tokio::io::{AsyncReadExt, AsyncWriteExt};

#[tokio::main]
async fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").await.unwrap();
    println!("Server running on 127.0.0.1:8080");

    loop {
        let (mut socket, _) = listener.accept().await.unwrap();
        tokio::spawn(async move {
            let mut buffer = [0; 1024];
            let n = socket.read(&mut buffer).await.unwrap();

            let request = String::from_utf8_lossy(&buffer[..n]);
            println!("Received: {}", request);

            let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
            socket.write_all(response.as_bytes()).await.unwrap();
        });
    }
}
```

---

### **Conclusion**

`epoll` est un outil puissant pour créer des serveurs performants en Rust, surtout sur Linux. Avec `mio`, vous pouvez facilement utiliser `epoll` pour gérer des milliers de connexions avec un seul thread.