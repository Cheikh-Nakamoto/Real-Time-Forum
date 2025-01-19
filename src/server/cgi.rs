// -------------------------------------------------------------------------------------
// CGI
// -------------------------------------------------------------------------------------
const FILETYPE: &str = ".py";
pub struct CGI {
    pub file_type: String
}

impl CGI {
    pub fn new() -> Self {
        Self {
            file_type: FILETYPE.to_string()
        }
    }

    pub fn execute_file(_filename: String) -> String {
        todo!()
    }
}
// -------------------------------------------------------------------------------------