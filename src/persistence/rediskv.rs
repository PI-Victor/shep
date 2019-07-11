extern crate redis;

#[derive(Deserialize, Debug)]
pub struct Redis {
    uri: Option<String>,
    user: Option<String>,
    password: Option<String>,
}

impl Default for Redis {
    fn default() -> Self {
        Redis{
            uri: Some("127.0.0.1:2397".to_string()),
            user: Some("root".to_string()),
            password: Some("root".to_string()),
        }
    }
}
