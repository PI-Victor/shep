extern crate config;

use config::{ConfigError, Config, File, Environment};


const GITHUB_API_URL:  &'static str = "";
const GITLAB_API_URL:  &'static str = "";
const BITBUCKET_API_URL:  &'static str = "";

#[derive(Debug, Deserialize)]
pub struct Configuration {
    // path self signed or CA valid SSL certificates to be used communication between the bot and
    // the vcs servers.
    pub ssl_certificates: String,
    // ip address for the application to bind to; defaults to 127.0.0.1;
    pub ip_address: String,
    // GitHub specific configuration
    pub git_hub: GitHub,
    // Jenkins specific configuration
    pub jenkins: JenkinsCI,
}

#[derive(Debug, Deserialize)]
pub struct GitHub {
    pub token: String,
}

#[derive(Debug, Deserialize)]
pub struct JenkinsCI {
    pub uri: String,
    pub token: String,
}

impl Configuration {
    pub fn new(path: &str) -> Result<Self, ConfigError> {
        let mut c = Config::new();
        c.merge(File::with_name(".shep.yaml"))?;
        c.merge(Environment::with_prefix("SHEP"))?;

        c.try_into()
    }
}
