// -------------------------------------------------------------------------------------
// REQUEST
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone)]
pub struct Request {
    pub id_session: String,
    pub location: String,
    pub host: String,
    pub port: u16,
    pub method: String,
    pub body: String,
}

impl Request {
    pub fn new(
        id_session: String,
        location: String,
        host: String,
        port: u16,
        method: String,
        body: String
    ) -> Self {
        Self { id_session, location, host, port, method, body }
    }
}
// -------------------------------------------------------------------------------------
