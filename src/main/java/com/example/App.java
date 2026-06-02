package com.example;
import com.sun.net.httpserver.HttpServer;
import java.net.InetSocketAddress;
import java.nio.charset.StandardCharsets;
public class App { public static void main(String[] args) throws Exception { int port = Integer.parseInt(System.getenv().getOrDefault("PORT", "8080")); HttpServer s = HttpServer.create(new InetSocketAddress("0.0.0.0", port), 0); s.createContext("/", ex -> { byte[] body = "axhub java maven ok".getBytes(StandardCharsets.UTF_8); ex.sendResponseHeaders(200, body.length); ex.getResponseBody().write(body); ex.close(); }); s.start(); } }
