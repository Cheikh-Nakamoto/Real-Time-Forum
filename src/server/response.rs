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
    /// Crée une nouvelle réponse.
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
        format!("HTTP/1.1 {}\r\n{}\r\n{}", self.status, headers, self.body)
    }

    // -------------------------------------------------------------------------------------
    // MÉTHODES D'ERREUR
    // -------------------------------------------------------------------------------------

    /// Renvoie une réponse 400 Bad Request.
    pub fn bad_request() -> Self {
        Self {
            id_session: String::new(), // Pas de session pour les erreurs
            status: "400 Bad Request".to_string(),
            content_type: "text/plain".to_string(),
            body: "400 Bad Request: The request could not be understood by the server.".to_string(),
        }
    }

    /// Renvoie une réponse 404 Not Found.
    pub fn not_found() -> Self {
        Self {
            id_session: String::new(),
            status: "404 Not Found".to_string(),
            content_type: "text/plain".to_string(),
            body: "404 Not Found: The requested resource was not found.".to_string(),
        }
    }

    /// Renvoie une réponse 500 Internal Server Error.
    pub fn internal_server_error() -> Self {
        Self {
            id_session: String::new(),
            status: "500 Internal Server Error".to_string(),
            content_type: "text/plain".to_string(),
            body: "500 Internal Server Error: The server encountered an unexpected condition.".to_string(),
        }
    }

    /// Renvoie une réponse 405 Method Not Allowed.
    pub fn method_not_allowed() -> Self {
        Self {
            id_session: String::new(),
            status: "405 Method Not Allowed".to_string(),
            content_type: "text/plain".to_string(),
            body: "405 Method Not Allowed: The requested method is not allowed for this resource.".to_string(),
        }
    }

    /// Renvoie une réponse 401 Unauthorized.
    pub fn unauthorized() -> Self {
        Self {
            id_session: String::new(),
            status: "401 Unauthorized".to_string(),
            content_type: "text/plain".to_string(),
            body: "401 Unauthorized: Authentication is required to access this resource.".to_string(),
        }
    }

    /// Renvoie une réponse 403 Forbidden.
    pub fn forbidden() -> Self {
        Self {
            id_session: String::new(),
            status: "403 Forbidden".to_string(),
            content_type: "text/plain".to_string(),
            body: "403 Forbidden: You do not have permission to access this resource.".to_string(),
        }
    }
}