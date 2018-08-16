extern crate config;

use config::{ConfigError, Config, File, Environment};


const GITHUB_API_URL:  &'static str = "";
const GITLAB_API_URL:  &'static str = "";
const BITBUCKET_API_URL:  &'static str = "";
const DEFAULT_IP_ADDR_BIND: &'static str = "127.0.0.1";


#[derive(Debug, Deserialize)]
pub struct Configuration {
    // path self signed or CA valid SSL certificates to be used communication between the bot and
    // the vcs servers.
    #[serde(default)]
    pub ssl_certificates: String,

    // ip address for the application to bind to; defaults to 127.0.0.1;
    #[serde(default)]
    pub ip_address: String,

    // GitHub specific configuration
    pub github: GitHub,

    // Jenkins specific configuration
    pub jenkins: JenkinsCI,
}

#[derive(Debug, Deserialize)]
pub struct GitHub {
    #[serde(default)]
    pub token: String,
}

#[derive(Debug, Deserialize)]
pub struct JenkinsCI {
    #[serde(default)]
    pub uri: String,
    #[serde(default)]
    pub token: String,
}

impl Configuration {
    pub fn new(path: &str) -> Result<Self, ConfigError> {
        let mut c = Config::new();
        c.merge(File::with_name(path))?;
        c.merge(Environment::with_prefix("SHEP"))?;

        c.try_into()
    }
}
