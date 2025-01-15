use chrono::Utc;
use uuid::Uuid;

// -------------------------------------------------------------------------------------
// REQUEST
// -------------------------------------------------------------------------------------
pub struct Request {
    pub url: String,
    pub method: String,
    pub body: String
}

impl Request {
    pub fn new(url: String, method: String, body: String) -> Self {
        Self { url, method, body }
    }
}
// -------------------------------------------------------------------------------------



// -------------------------------------------------------------------------------------
// RESPONSE
// -------------------------------------------------------------------------------------
pub struct Response {
    pub url: String,
    pub method: String,
    pub body: String
}

impl Response {
    pub fn new(url: String, method: String, body: String) -> Self {
        Self { url, method, body }
    }
}
// -------------------------------------------------------------------------------------



// -------------------------------------------------------------------------------------
// SESSION
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone)]
pub struct Session {
    pub id: String,
    pub expiration_time: i64,
}

const SESSION_LIFETIME: i64 = 60 * 60 * 1000;

impl Session {
    pub fn new() -> Self {
        Self {
            id: Uuid::new_v4().into(),
            expiration_time: Utc::now().timestamp_millis() + SESSION_LIFETIME,
        }
    }
}
// -------------------------------------------------------------------------------------

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

    pub fn handle_request(req: Request) -> Response {
        todo!()
    }
}
// -------------------------------------------------------------------------------------

// -------------------------------------------------------------------------------------
// ROUTER
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone)]
pub struct Router {
    pub servers: Vec<Server>,
    pub sessions: Vec<Session>,
}

impl Router {
    pub fn new() -> Self {
        Self { servers: vec![], sessions: vec![] }
    }

    pub fn add_server(&mut self, server: Server) {
        self.servers.push(server);
    }

    pub fn remove_server(mut self, server: Server) {
        self.servers = self.servers
            .into_iter()
            .filter(|s| s.ip_addr != server.ip_addr && s.hostname != server.hostname)
            .collect();
    }

    pub fn add_session(&mut self, session: Session) {
        self.sessions.push(session);
    }

    pub fn remove_session(mut self, session_id: String) {
        self.sessions = self.sessions
            .into_iter()
            .filter(|s| s.id != session_id)
            .collect();
    }

    pub fn route_request(req: Request) -> Response {
        todo!()
    }
}
// -------------------------------------------------------------------------------------
