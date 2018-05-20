use std::net::{TcpStream, TcpListener};

use util::{config};

pub fn start(c: config) {
    let server = TcpListener::bind(c.ip_address);
}