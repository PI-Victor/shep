extern crate clap;

use clap::{App, AppSettings, Arg};

mod core;
mod vcs;
mod util;

use core::{start};
use util::{config};

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
    let m = App::new("shep")
            .author("Cloudflavor Org")
            .version(VERSION)
            .about(ASCIIART)
            .setting(AppSettings::SubcommandRequiredElseHelp)
            .arg(Arg::with_name("start")
                .help("Start the application"))
            .arg(Arg::with_name("config")
                .help("Configuration location"))
            .get_matches();
    
    let c = config{
        ip_address: String()
    };

    start(c);
}