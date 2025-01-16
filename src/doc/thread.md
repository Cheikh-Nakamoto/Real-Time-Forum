## **Qu'est-ce qu'un thread ?**

Un thread est une unité d'exécution indépendante au sein d'un programme. Chaque thread a son propre stack (pile d'exécution) et peut exécuter du code en parallèle avec d'autres threads. Les threads partagent la mémoire du processus parent, ce qui permet une communication facile entre eux.

### **Avantages des threads**
1. **Concurrence** : Exécuter plusieurs tâches en parallèle.
2. **Performance** : Tirer parti des processeurs multi-cœurs.
3. **Simplicité** : Modéliser des tâches indépendantes de manière naturelle.

### **Inconvénients des threads**
1. **Complexité** : La gestion des threads peut être difficile (synchronisation, partage de données).
2. **Surcharge** : Chaque thread consomme des ressources (mémoire, temps de création).
3. **Problèmes de concurrence** : Risque de race conditions, deadlocks, etc.

---

## **Utilisation des threads en Rust**

Rust fournit une API simple et sûre pour travailler avec les threads via le module `std::thread`. Voici les concepts clés :

1. **Créer un thread** : Utilisez `thread::spawn`.
2. **Attendre un thread** : Utilisez `join` pour attendre la fin d'un thread.
3. **Partager des données** : Utilisez `Arc` (Atomic Reference Counting) et `Mutex` pour partager des données entre threads de manière sûre.

---

### **Exemple de base : Créer un thread**

Voici un exemple simple de création d'un thread en Rust :

```rust
use std::thread;
use std::time::Duration;

fn main() {
    // Créer un nouveau thread
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("Thread: {}", i);
            thread::sleep(Duration::from_millis(500));
        }
    });

    // Code exécuté dans le thread principal
    for i in 1..5 {
        println!("Main: {}", i);
        thread::sleep(Duration::from_millis(300));
    }

    // Attendre que le thread se termine
    handle.join().unwrap();
}
```

#### Explication :
- `thread::spawn` : Crée un nouveau thread et exécute la closure passée en argument.
- `handle.join()` : Attend que le thread se termine avant de continuer.

---

### **Partager des données entre threads**

Pour partager des données entre threads, Rust impose des règles strictes pour éviter les problèmes de concurrence. Vous devez utiliser `Arc` (pour le partage de propriété) et `Mutex` (pour la synchronisation).

#### Exemple :

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    // Donnée partagée entre threads
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }

    // Attendre que tous les threads se terminent
    for handle in handles {
        handle.join().unwrap();
    }

    println!("Résultat final: {}", *counter.lock().unwrap());
}
```

#### Explication :
- `Arc` : Permet de partager la propriété de la donnée entre plusieurs threads.
- `Mutex` : Assure que seule une thread à la fois peut accéder à la donnée.

---

### **Utiliser des threads dans un serveur web**

Dans un serveur web, les threads sont souvent utilisés pour gérer plusieurs connexions client simultanément. Voici un exemple de serveur web simple utilisant des threads :

```rust
use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};
use std::thread;

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
    println!("Server listening on 127.0.0.1:8080");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                // Créer un nouveau thread pour chaque connexion
                thread::spawn(|| {
                    handle_connection(stream);
                });
            }
            Err(e) => {
                eprintln!("Failed to establish a connection: {}", e);
            }
        }
    }
}

fn handle_connection(mut stream: TcpStream) {
    let mut buffer = [0; 1024];
    stream.read(&mut buffer).unwrap();

    let request = String::from_utf8_lossy(&buffer);
    println!("Received request: {}", request);

    let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
    stream.write(response.as_bytes()).unwrap();
    stream.flush().unwrap();
}
```

#### Explication :
- `thread::spawn` : Crée un nouveau thread pour chaque connexion client.
- `handle_connection` : Gère la connexion dans un thread séparé.

---

### **Limites des threads**

1. **Surcharge** : Chaque thread consomme de la mémoire et du temps CPU pour sa création et sa gestion.
2. **Scalabilité** : Un grand nombre de threads peut entraîner une surcharge du système.
3. **Complexité** : La synchronisation entre threads peut être difficile à gérer.

---

### **Alternative : Programmation asynchrone**

Pour une meilleure scalabilité, vous pouvez utiliser la programmation asynchrone avec des crates comme `tokio` ou `async-std`. Ces crates utilisent des tâches légères (tasks) au lieu de threads, ce qui permet de gérer des milliers de connexions avec un seul thread.

#### Exemple avec `tokio` :

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

Les threads sont un outil puissant pour gérer des tâches concurrentes en Rust. Ils sont particulièrement utiles pour les serveurs web, mais peuvent devenir complexes à gérer à grande échelle. Pour des applications hautement concurrentes, envisagez d'utiliser la programmation asynchrone avec des crates comme `tokio`.