use axum::{routing::get, Router};
use std::net::SocketAddr;
#[tokio::main]
async fn main() {
    let port: u16 = std::env::var("PORT").ok().and_then(|v| v.parse().ok()).unwrap_or(8080);
    let app = Router::new().route("/", get(|| async { "axhub rust axum ok" }));
    let listener = tokio::net::TcpListener::bind(SocketAddr::from(([0, 0, 0, 0], port))).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
