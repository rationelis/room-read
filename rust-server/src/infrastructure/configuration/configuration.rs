use serde::Deserialize;
use serde_yaml;
use std::fs::File;

#[derive(Deserialize, Debug)]
pub struct Configuration {
    pub environment: String,
    pub server: ServerConfiguration,
}

#[derive(Deserialize, Debug)]
pub struct ServerConfiguration {
    pub host: String,
    pub port: u16,
}

pub fn load_configuration() -> Result<Configuration, String> {
    let file_result = File::open("config.yml");

    match file_result {
        Ok(file) => match serde_yaml::from_reader(file) {
            Ok(config) => Ok(config),
            Err(err) => Err(format!("Failed to parse YAML: {}", err)),
        },
        Err(err) => Err(format!("Failed to open the file: {}", err)),
    }
}
