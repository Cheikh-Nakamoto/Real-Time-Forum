pub use serde::{Deserialize, Serialize};

use std::string::String;

// -------------------------------------------------------------------------------------
// DIRECTORY ELEMENT
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DirectoryElement {
    pub entry: String,
    pub link: String,
    pub entry_type: String,
    pub is_directory: bool,
}
// -------------------------------------------------------------------------------------



// -------------------------------------------------------------------------------------
// HTMLError
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct HTMLError {
    pub code: u16,
    pub status: String,
}
// -------------------------------------------------------------------------------------
