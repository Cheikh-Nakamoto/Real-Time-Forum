
## **Qu'est-ce que `TcpListener` ?**

`TcpListener` est une structure fournie par la bibliothèque standard de Rust (`std::net`) qui permet de créer un serveur TCP. Il écoute les connexions entrantes sur une adresse IP et un port spécifiés, et permet d'accepter ces connexions pour communiquer avec des clients.

En résumé, `TcpListener` est utilisé pour :
1. **Écouter les connexions entrantes** sur un port donné.
2. **Accepter les connexions** et obtenir un `TcpStream` pour chaque client.
3. **Gérer les connexions** de manière synchrone ou asynchrone.

---

## **Comment utiliser `TcpListener` ?**

### **1. Créer un `TcpListener`**

Pour créer un `TcpListener`, vous devez spécifier une adresse IP et un port. La méthode `bind` est utilisée pour cela.

#### Exemple :
```rust
use std::net::TcpListener;

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
    println!("Server listening on 127.0.0.1:8080");
}
```

- `bind("127.0.0.1:8080")` : Crée un `TcpListener` qui écoute sur l'adresse IP `127.0.0.1` (localhost) et le port `8080`.
- `unwrap()` : Détecte les erreurs (par exemple, si le port est déjà utilisé).

---

### **2. Accepter les connexions entrantes**

Une fois le `TcpListener` créé, vous pouvez accepter les connexions entrantes en utilisant la méthode `incoming()`. Cette méthode retourne un itérateur sur les connexions entrantes.

#### Exemple :
```rust
use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
    println!("Server listening on 127.0.0.1:8080");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                handle_connection(stream);
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

    let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
    stream.write(response.as_bytes()).unwrap();
    stream.flush().unwrap();
}
```

- `listener.incoming()` : Retourne un itérateur sur les connexions entrantes.
- `stream` : Chaque connexion est représentée par un `TcpStream`, qui permet de lire et d'écrire des données.

---

### **3. Gérer les connexions de manière concurrente**

Pour gérer plusieurs clients simultanément, vous pouvez utiliser des threads. Chaque connexion est gérée dans un thread séparé.

#### Exemple :
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

    let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
    stream.write(response.as_bytes()).unwrap();
    stream.flush().unwrap();
}
```

- `thread::spawn` : Crée un nouveau thread pour chaque connexion.
- Cela permet de gérer plusieurs clients en parallèle.

---

### **4. Gérer les erreurs**

Il est important de gérer les erreurs correctement pour éviter que le serveur ne plante. Utilisez `Result` et `match` pour gérer les erreurs de manière élégante.

#### Exemple :
```rust
use std::net::{TcpListener, TcpStream};
use std::io::{Read, Write};

fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").unwrap();
    println!("Server listening on 127.0.0.1:8080");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                if let Err(e) = handle_connection(stream) {
                    eprintln!("Error handling connection: {}", e);
                }
            }
            Err(e) => {
                eprintln!("Failed to establish a connection: {}", e);
            }
        }
    }
}

fn handle_connection(mut stream: TcpStream) -> std::io::Result<()> {
    let mut buffer = [0; 1024];
    stream.read(&mut buffer)?;

    let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
    stream.write(response.as_bytes())?;
    stream.flush()?;

    Ok(())
}
```

- `?` : Propage les erreurs sans faire paniquer le programme.

---

### **5. Utiliser `TcpListener` de manière asynchrone**

Pour une meilleure performance, vous pouvez utiliser `TcpListener` de manière asynchrone avec des bibliothèques comme `tokio` ou `async-std`.

#### Exemple avec `tokio` :
```rust
use tokio::net::TcpListener;
use tokio::io::{AsyncReadExt, AsyncWriteExt};

#[tokio::main]
async fn main() {
    let listener = TcpListener::bind("127.0.0.1:8080").await.unwrap();
    println!("Server listening on 127.0.0.1:8080");

    loop {
        let (mut stream, _) = listener.accept().await.unwrap();
        tokio::spawn(async move {
            let mut buffer = [0; 1024];
            let n = stream.read(&mut buffer).await.unwrap();

            let request = String::from_utf8_lossy(&buffer[..n]);
            println!("Received: {}", request);

            let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
            stream.write_all(response.as_bytes()).await.unwrap();
        });
    }
}
```

- `tokio::spawn` : Crée une tâche asynchrone pour chaque connexion.
- `await` : Attend que les opérations asynchrones se terminent.

---

### **6. Méthodes utiles de `TcpListener`**

Voici quelques méthodes couramment utilisées de `TcpListener` :

| Méthode                          | Description                                                                 |
|----------------------------------|-----------------------------------------------------------------------------|
| `bind(addr: SocketAddr)`         | Crée un `TcpListener` qui écoute sur l'adresse et le port spécifiés.        |
| `incoming()`                     | Retourne un itérateur sur les connexions entrantes.                         |
| `local_addr()`                   | Retourne l'adresse locale sur laquelle le `TcpListener` écoute.             |
| `set_nonblocking(nonblocking: bool)` | Active ou désactive le mode non-bloquant pour le `TcpListener`.          |
| `accept()`                       | Accepte une nouvelle connexion et retourne un `(TcpStream, SocketAddr)`.    |

---

### **Exemple complet**

Voici un exemple complet de serveur TCP synchrone avec `TcpListener` :

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
                thread::spawn(|| {
                    if let Err(e) = handle_connection(stream) {
                        eprintln!("Error handling connection: {}", e);
                    }
                });
            }
            Err(e) => {
                eprintln!("Failed to establish a connection: {}", e);
            }
        }
    }
}

fn handle_connection(mut stream: TcpStream) -> std::io::Result<()> {
    let mut buffer = [0; 1024];
    stream.read(&mut buffer)?;

    let request = String::from_utf8_lossy(&buffer);
    println!("Received request: {}", request);

    let response = "HTTP/1.1 200 OK\r\n\r\nHello, World!";
    stream.write(response.as_bytes())?;
    stream.flush()?;

    Ok(())
}
```

---

### **Conclusion**

`TcpListener` est un outil puissant pour créer des serveurs TCP en Rust. Que vous utilisiez une approche synchrone ou asynchrone, il vous permet de gérer les connexions entrantes de manière efficace. Avec les exemples et explications fournis, vous devriez être en mesure de créer un serveur robuste et performant.
