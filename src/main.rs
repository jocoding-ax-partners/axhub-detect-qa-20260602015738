fn main() {
    let message = std::env::var("APPHUB_QA_RUST_MESSAGE").unwrap_or_else(|_| "rust ok".to_string());
    println!("language=rust message={message}");
}
