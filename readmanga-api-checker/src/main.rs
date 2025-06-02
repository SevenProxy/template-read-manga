use readmanga_api_checker::{ fmt, start_server, web, App, AppState, HttpServer };
use readmanga_api_checker::{ checker_lain };

const PORT: u16 = 3002;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
  fmt()
    .with_max_level(tracing::Level::INFO)
    .with_target(false)
    .pretty()
    .init();
  start_server();
  HttpServer::new(|| {
  App::new()
    .app_data(web::Data::new(AppState {
      app_name: String::from("Checker API"),
    }))
    .service(
      web::scope("/api")
        .service(checker_lain)
    )
  })
  .bind(("0.0.0.0", PORT))?
  .run()
  .await
}
