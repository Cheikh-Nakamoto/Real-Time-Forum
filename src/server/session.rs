use chrono::Utc;
use uuid::Uuid;

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