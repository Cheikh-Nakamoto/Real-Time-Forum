pub mod request;
use std::{ net::{ TcpListener, TcpStream }, thread };

use mio::Token;
pub use request::*;
pub mod response;
pub use response::*;
pub mod router;
pub use router::*;
pub mod session;
pub use session::*;
pub mod cgi;
pub use cgi::*;
pub mod client;
pub use client::*;


// -------------------------------------------------------------------------------------
// SERVER
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone)]
pub struct Server {
    pub ip_addr: String,
    pub hostname: String,
    pub ports: Vec<u16>,
    pub root_directory: String,
    pub error_path: String,
    pub default_file: String,
    pub access_log: String,
    pub cgi_file_format: String,
    pub upload_limit: u32,
    pub accepted_methods: Vec<String>,
    pub directory_listing: bool,
}

impl Server {
    pub fn new(
        ip_addr: String,
        hostname: String,
        ports: Vec<u16>,
        root_directory: String,
        error_path: String,
        default_file: String,
        access_log: String,
        cgi_file_format: String,
        upload_limit: u32,
        accepted_methods: Vec<String>,
        directory_listing: bool
    ) -> Self {
        Self {
            ip_addr,
            hostname,
            ports,
            root_directory,
            error_path,
            default_file,
            access_log,
            cgi_file_format,
            upload_limit,
            accepted_methods,
            directory_listing,
        }
    }

    pub fn start(&self) {
        for port in &self.ports {
            let mut client : Client = Client::new(format!("{}:{}", self.ip_addr, port).as_str());
            client.run();
        }
    }

    pub fn stop() {}

    // fn handle_request(&self, mut stream: TcpStream) {
    //     let mut buffer = [0; 1024];
    //     stream.read(&mut buffer).unwrap();

    //     let request = Request::from_bytes(&buffer).unwrap();
    //     let response = self.handle_request(request);

    //     stream.write(&response.to_bytes()).unwrap();
    //     stream.flush().unwrap();
    // }
    // pub fn access_log(&self, req: &Request) {
    //     let mut file = OpenOptions::new()
    //         .append(true)
    //         .create(true)
    //         .open(&self.access_log)
    //         .unwrap();

    //     writeln!(file, "{} {} {}", req.method, req.path, "200 OK").unwrap();
    // }
}
// -------------------------------------------------------------------------------------
