mod infrastructure {
    pub mod configuration {
        pub mod configuration;
    }
}

use infrastructure::configuration::configuration::load_configuration;

fn main() {
    let configuration = load_configuration();

    if configuration.is_err() {
        panic!(
            "Failed to load configuration: {}",
            configuration.err().unwrap()
        );
    }

    println!("A key-value: {:?}", configuration.environment);
}
