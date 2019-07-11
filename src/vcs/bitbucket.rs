
const BITBUCKET_API_URL:  &str = "https://api.bitbucket.org/2.0";

#[derive(Serialize, Deserialize, Debug)]
pub struct BitBucket {
    pub user: Option<String>,
    pub token: Option<String>,
    pub password: Option<String>,
}
