extern crate ctrlc;

use std::sync::Arc;
use std::sync::atomic::{AtomicBool, Ordering};

use util::Configuration;

pub fn start(c: Configuration) {
    let running = Arc::new(AtomicBool::new(true));
    let r = running.clone();

    ctrlc::set_handler(move || {
        r.store(false, Ordering::SeqCst)
    }).expect("Error setting CTRL-C handler");

    while running.load(Ordering::SeqCst) {

    }
    println!("Received disconnect! Quitting")
}
