// -------------------------------------------------------------------------------------
// RESPONSE
// -------------------------------------------------------------------------------------
pub struct Response {
    pub status_line: String,
    pub content_type: String,
    pub body: String,
}

impl Response {
    pub fn new(status_line: String, content_type: String, body: String) -> Self {
        Self {
            status_line,
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
            self.status_line, headers, self.body
        )
    }
}