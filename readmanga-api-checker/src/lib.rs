mod log;
mod error;
mod dto;
mod adapter;
mod controllers;

pub use actix_web::{HttpServer, App, web, web::Data};
pub use log::start_server;
pub use tracing_subscriber::fmt;

pub use controllers::checker_lain::Checker;

pub struct AppState {
  pub app_name: String,
}
