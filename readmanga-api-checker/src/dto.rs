use serde::Serialize;
use serde_json::Value;

#[derive(Serialize)]
pub struct JsonResponse {
  pub status: bool,
  #[serde(skip_serializing_if = "Option::is_none")]
  pub message: Option<String>,

  #[serde(skip_serializing_if = "Option::is_none")]
  pub data: Option<Value>,
}
