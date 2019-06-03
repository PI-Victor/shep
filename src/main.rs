#[macro_use]
extern crate log;
extern crate clap;
extern crate config;
extern crate serde;
extern crate env_logger;

#[macro_use]
extern crate serde_derive;

mod core;
mod util;

use clap::{Arg, App, SubCommand, AppSettings};

use core::{start};
use util::{Configuration, set_logger};

const ASCIIART: &str = r#"
 _____ _    _ ______ _____
/ ____| |  | |  ____|  __ \
| (___| |__| | |__  | |__) |
\___ \|  __  |  __| | ___/
____) | |  | | |____| |
|____/|_|  |_|______|_|
"#;

const VERSION: &str = "v0.1";


fn main() {
    
    let matches = App::new("shep")
        .author("Cloudflavor Org")
        .version(VERSION)
        .about(ASCIIART)
        .setting(AppSettings::SubcommandRequiredElseHelp)
        .subcommand(
            SubCommand::with_name("start")
                .about("Start the bot")
                .arg(
                    Arg::with_name("config")
                        .short("c")
                        .long("config")
                        .value_name("JSON, TOML, YAML, HJSON, INI - configuration")
                        .help("Path to config file")
                        .takes_value(true)
                        .required(true)
                )
        )
        .get_matches();

    set_logger();

    if let Some(matches) = matches.subcommand_matches("start") {
        let config = Configuration::new(matches.value_of("config").unwrap());
        debug!("Loaded configuration: {:?}", config);
        start(config.unwrap());
    }
}
