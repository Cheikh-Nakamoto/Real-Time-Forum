mod server;  // Importe le module server (mod.rs)
use server::{Router, Server};

fn main() -> std::io::Result<()> {
    // Crée un serveur
    let server = Server::new(
        "localhost-01".to_string(),         // ip_addr
        "localhost".to_string(),         // hostname
        vec![8080, 8081],                // ports
        "./www".to_string(),             // root_directory
        "./errors".to_string(),          // error_path
        "index.html".to_string(),        // default_file
        "access.log".to_string(),        // access_log
        "php".to_string(),               // cgi_file_format
        10_000_000,                      // upload_limit (10 Mo)
        vec!["GET".to_string(), "POST".to_string()], // accepted_methods
        false,                           // directory_listing
    );

    // Crée un routeur et ajoute le serveur
    let mut router = Router::new();
    router.add_server(server)?;

    // Démarre le routeur
    println!("Serveur en écoute sur les ports 8080 et 8081...");
    router.run()?;

    Ok(())
}