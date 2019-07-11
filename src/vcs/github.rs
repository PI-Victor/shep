const GITHUB_API_URL:  &str = "https://api.github.com";

#[derive(Deserialize, Debug)]
pub struct GitHub {
    pub user: Option<String>,
    pub token: Option<String>,
    pub password: Option<String>,
}