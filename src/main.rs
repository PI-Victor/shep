#[macro_use]
extern crate log;
extern crate actix;
extern crate clap;
extern crate config;
extern crate serde;
extern crate env_logger;
#[macro_use]
extern crate serde_derive;

mod util;
mod vcs;
mod persistence;

use clap::{Arg, App, SubCommand, AppSettings};
use config::{ConfigError, Config, File, Environment};

use util::{Configuration};


const ASCIIART: &str = r#"
 _____ _    _ ______ _____
/ ____| |  | |  ____|  __ \
| (___| |__| | |__  | |__) |
\___ \|  __  |  __| | ___/
____) | |  | | |____| |
|____/|_|  |_|______|_|
"#;

const VERSION: &str = "v0.1.0-alpha";

fn main() {
    let matches = App::new("shep")
        .author("Cloudflavor Org")
        .version(VERSION)
        .about(ASCIIART)
        .setting(AppSettings::SubcommandRequiredElseHelp)
        .arg(Arg::with_name("v")
            .help("verbosity level 0-3")
            .short("v")
            .multiple(true))
        .subcommand(
            SubCommand::with_name("start")
                .about("Start the application")
                .arg(Arg::with_name("config")
                        .short("c")
                        .long("config")
                        .value_name("JSON, TOML, YAML, HJSON, INI - configuration")
                        .help("Path to config file")
                        .takes_value(true)
                        .required(true)))
        .get_matches();

    let mut c = Configuration {..Default::default()};

    if let Some(matches) = matches.subcommand_matches("start") {
        let config = Configuration::new(matches.value_of("config").unwrap());
        c = config.unwrap();
    }

    let log_level = match matches.occurrences_of("v") {
        0 => log::LevelFilter::Warn,
        1 => log::LevelFilter::Info,
        2 => log::LevelFilter::Debug,
        _ => log::LevelFilter::Trace,
    };
    info!("{:?}", log_level);

    env_logger::Builder::from_default_env()
        .filter(Some(module_path!()), log_level)
        .init();

    debug!("Loaded configuration: {:?}", c);

    let os = actix::System::new("shep");

    info!("Shep up and running...");
}

impl Configuration {
    pub fn new(path: &str) -> Result<Self, ConfigError> {
        let mut c = Config::new();
        c.merge(File::with_name(path))?;
        c.merge(Environment::with_prefix("SHEP"))?;
        c.try_into()
    }
}
