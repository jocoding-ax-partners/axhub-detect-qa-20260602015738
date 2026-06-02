FROM rust:1.83-slim AS builder
WORKDIR /app
COPY . .
RUN cargo build --release

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/target/release/axhub-detect-rust /app/axhub-detect-rust
EXPOSE 8080
CMD ["/app/axhub-detect-rust"]
