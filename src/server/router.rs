pub use super::{Server, Session};
pub use crate::{Request, Response};

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