package com.example;

public class App {
  public static void main(String[] args) {
    String port = System.getenv().getOrDefault("PORT", "8080");
    System.out.println("axhub gradle qa listening on " + port);
  }
}
