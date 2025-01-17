pub mod server;
use std::{ collections::HashMap, fs };
use serde::Deserialize;

pub use server::*;

#[derive(Debug, Deserialize, Clone)]
pub struct Config {
    pub log_files: LogFilesConfig,
    pub http: HttpConfig,
}

impl Config {
    pub fn new() -> Self {
        Self {
            log_files: LogFilesConfig {
                error_log: String::new(),
                access_log: String::new(),
                events_limit: 0,
            },
            http: HttpConfig {
                access_log_format: String::new(),
                timeout: 0,
                size_limit: 0,
                servers: HashMap::new(),
            },
        }
    }
}

#[derive(Debug, Deserialize, Clone)]
pub struct LogFilesConfig {
    pub error_log: String,
    pub access_log: String,
    pub events_limit: u32,
}

#[derive(Debug, Deserialize, Clone)]
pub struct HttpConfig {
    pub access_log_format: String,
    pub timeout: u64,
    pub size_limit: u64,
    pub servers: HashMap<String, Server>,
}

pub fn load_config() -> Config {
    let content = fs::read_to_string("src/config.toml").unwrap_or(String::new());
    toml::from_str(&content).unwrap()
}