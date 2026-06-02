FROM rust:1.83-alpine AS build
RUN apk add --no-cache musl-dev
WORKDIR /app
COPY . .
RUN cargo build --release
FROM alpine:3.20
COPY --from=build /app/target/release/axhub-rust-axum /usr/local/bin/app
EXPOSE 8080
CMD ["/usr/local/bin/app"]
