use std::process::Command;

// -------------------------------------------------------------------------------------
// CGI
// -------------------------------------------------------------------------------------
// const FILETYPE: &str = ".rb";
pub struct CGI;

impl CGI {
    // pub fn new() -> Self {
    //     Self {
    //         file_type: FILETYPE.to_string(),
    //     }
    // }

    pub fn execute_file(filename: String) -> String {
        let mut input = Command::new("ruby");

        input.arg(filename);
        let res = input.output();

        if res.is_ok() {
            let check_txt = String::from_utf8(res.unwrap().stdout);
            if check_txt.is_ok() {
                check_txt.unwrap()
            } else {
                String::new()
            }
        } else {
            String::new()
        }
    }
}
// -------------------------------------------------------------------------------------
