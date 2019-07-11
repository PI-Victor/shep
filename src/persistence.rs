extern crate redis;

struct Redis {
    uri: Option<String>,
    user: Option<String>,
    password: Option<String>,
}