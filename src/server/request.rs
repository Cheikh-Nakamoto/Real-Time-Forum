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