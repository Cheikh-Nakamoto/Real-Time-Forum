pub mod request;
use std::io::Error;

pub use request::*;
pub mod response;
pub use response::*;
pub mod router;
pub use router::*;
pub mod session;
pub use session::*;

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

    pub fn start() {}

    pub fn stop() {}

    pub fn handle_request(req: Request) -> Response {
        todo!()
    }
}
// -------------------------------------------------------------------------------------
