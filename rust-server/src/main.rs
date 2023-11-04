mod adapters {
    pub mod rest {
        pub mod rest_controller;
    }
}

mod domain {
    pub mod model {
        pub mod event;
    }
    pub mod ports {
        pub mod i_event_port;
    }
}

mod infrastructure {
    pub mod configuration {
        pub mod configuration;
    }
    pub mod server {
        pub mod server;
    }
}

use crate::infrastructure::configuration::configuration::Configuration;

fn main() {
    let configuration = Configuration::new();

    if let Err(err) = configuration {
        panic!("Failed to load configuration: {}", err);
    }

    let server = infrastructure::server::server::new_server(configuration.unwrap());
}
