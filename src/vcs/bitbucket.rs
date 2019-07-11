#[derive(Deserialize, Debug)]
pub struct BitBucket {
    pub user: Option<String>,
    pub token: Option<String>,
    pub password: Option<String>,
    pub api_url: Option<String>,
}

impl Default for BitBucket {
    fn default() -> Self {
        BitBucket{
            api_url: Some("https://api.bitbucket.org/2.0".to_string()),
            user: None,
            password: None,
            token: None,
        }
    }
}
