use vcs::{GitHub, GitLab, BitBucket};
use persistence::{Redis};


#[derive(Deserialize, Debug)]
pub struct Configuration {
    // GitHub specific configuration
    pub github: Option<GitHub>,

    // GitLab specific configuration
    pub gitlab: Option<GitLab>,

    // Bitbucket specific configuration
    pub bitbucket: Option<BitBucket>,

    // Redis configuration
    pub redis: Option<Redis>,

}

#[derive(Serialize, Deserialize, Debug)]
pub struct CIService {
    pub uri: String,
    pub token: String,
}

impl Default for Configuration {
    fn default() -> Self {
        Configuration{
            github: None,
            gitlab: None,
            bitbucket: None,
            redis: None,
        }
    }
}