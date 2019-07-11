
const GITLAB_API_URL:  &str = "https://gitlab.example.com/api/v4/";

#[derive(Deserialize, Debug)]
pub struct GitLab {
    pub user: Option<String>,
    pub token: Option<String>,
    pub password: Option<String>,
}