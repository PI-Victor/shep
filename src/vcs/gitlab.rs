#[derive(Deserialize, Debug)]
pub struct GitLab {
    pub user: Option<String>,
    pub token: Option<String>,
    pub password: Option<String>,
    api_url: String,
}

impl Default for GitLab {
    fn default() -> Self {
        GitLab{
            api_url: "https://gitlab.example.com/api/v4/".to_string(),
            user: None,
            token: None,
            password: None,
        }
    }
}
