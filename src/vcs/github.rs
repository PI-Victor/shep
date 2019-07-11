#[derive(Deserialize, Debug)]
pub struct GitHub {
    pub user: Option<String>,
    pub token: Option<String>,
    pub password: Option<String>,
    pub api_url: Option<String>,
}

impl Default for GitHub {
    fn default() -> Self {
        GitHub{
            api_url: Some("https://api.github.com".to_string()),
            user: None,
            token: None,
            password: None
        }
    }
}
