extern crate ctrlc;
extern crate log;

use std::sync::Arc;
use std::sync::atomic::{AtomicBool, Ordering};
use log::{info};

use util::Configuration;

pub fn start(c: Configuration) {

    let running = Arc::new(AtomicBool::new(true));
    let r = running.clone();

    ctrlc::set_handler(move || {
        r.store(false, Ordering::SeqCst)
    }).expect("Error setting CTRL-C handler");
    
    info!("Starting application");

    while running.load(Ordering::SeqCst) {
        
    }

    info!("Received disconnect! Quitting")
}
