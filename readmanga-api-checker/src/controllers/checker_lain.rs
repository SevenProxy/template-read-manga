use actix_web::{get, HttpRequest};

use crate::{adapter::{ Response }, dto::JsonResponse};

#[get("/hello")]
pub async fn checker_lain(_req: HttpRequest) -> Response {
  let message_response: JsonResponse = JsonResponse {
    status: true,
    message: Some(String::from("hello worlds")),
    data: None,
  };
  Response::ok(message_response)
}
