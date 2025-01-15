// -------------------------------------------------------------------------------------
// RESPONSE
// -------------------------------------------------------------------------------------
pub struct Response {
    pub url: String,
    pub method: String,
    pub body: String,
}

impl Response {
    pub fn new(url: String, method: String, body: String) -> Self {
        Self { url, method, body }
    }
}
// -------------------------------------------------------------------------------------
