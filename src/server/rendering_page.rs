pub use serde::{Deserialize, Serialize};

use std::string::String;

// -------------------------------------------------------------------------------------
// DIRECTORY ELEMENT
// -------------------------------------------------------------------------------------
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DirectoryElement {
    pub entry: String,
    pub link: String,
    pub is_directory: bool,
}
// -------------------------------------------------------------------------------------
