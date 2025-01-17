// -------------------------------------------------------------------------------------
// RESPONSE
// -------------------------------------------------------------------------------------
pub struct Response {
    pub id_session: String,
    pub status: String,
    pub content_type: String,
    pub body: String,
}

impl Response {
    pub fn new(id_session: String, status: String, content_type: String, body: String) -> Self {
        Self {
            id_session,
            status,
            content_type,
            body,
        }
    }

    /// Convertit la réponse en une chaîne de caractères HTTP valide.
    pub fn to_http_response(&self) -> String {
        let headers = format!(
            "Content-Type: {}\r\nContent-Length: {}\r\n",
            self.content_type,
            self.body.len()
        );
        format!(
            "{}\r\n{}\r\n{}",
            self.status, headers, self.body
        )
    }
}