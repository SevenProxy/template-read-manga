use actix_web::HttpRequest;
use readmanga_api_checker::{ fmt, start_server, web, Data, App, AppState, HttpServer, Checker };

const PORT: u16 = 3002;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
  let checker: Data<Checker> = web::Data::new(Checker {});
  let name_server: Data<AppState> = web::Data::new(AppState {
    app_name: String::from("Checker API"),
  });

  fmt()
    .with_max_level(tracing::Level::INFO)
    .with_target(false)
    .pretty()
    .init();
  start_server();
  
  HttpServer::new(move || {
  App::new()
    .app_data(name_server.clone())
    .app_data(checker.clone())
    .service(
      web::scope("/api")
        .route("/checker-lain", web::get().to(
          | req: HttpRequest, checker: Data<Checker> | async move {
            checker.lain(req).await
          }
        ))
    )
  })
  .bind(("0.0.0.0", PORT))?
  .run()
  .await
}
