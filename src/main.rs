extern crate clap;
extern crate config;
extern crate serde;

#[macro_use]
extern crate serde_derive;

mod core;
mod vcs;
mod util;
mod logger;

use clap::{Arg, App, SubCommand, AppSettings};

use core::{start};
use util::{Configuration};

const ASCIIART: &str = r#"
 _____ _    _ ______ _____
/ ____| |  | |  ____|  __ \
| (___| |__| | |__  | |__) |
\___ \|  __  |  __| | ___/
____) | |  | | |____| |
|____/|_|  |_|______|_|

A cloud aware bot for CI/CD.
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
                    .about("Start Shep CI/CD bot")
                    .arg(
                        Arg::with_name("config")
                            .short("c")
                            .long("config")
                            .value_name(".yaml")
                            .help("Full path to yaml config")
                            .takes_value(true)
                            .required(true)
                    )
            )
            .get_matches();

    if let Some(matches) = matches.subcommand_matches("start") {
        let config = Configuration::new(matches.value_of("config").unwrap());
        println!("this is the config {:?}", config);

    }
}
