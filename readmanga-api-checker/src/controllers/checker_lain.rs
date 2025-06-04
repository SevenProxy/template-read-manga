use actix_web::HttpRequest;

use crate::{adapter::{ Response }, dto::JsonResponse};

pub struct Checker {
  // usecase
}

impl Checker {
  pub async fn lain(&self, _req: HttpRequest) -> Response {
    let message_response: JsonResponse = JsonResponse {
      status: true,
      message: Some(String::from("hello worlds")),
      data: None,
    };
    Response::ok(message_response)
  }
}
