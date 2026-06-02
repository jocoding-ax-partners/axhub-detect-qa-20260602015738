package com.example;

public class App {
  public static void main(String[] args) {
    String msg = System.getenv().getOrDefault("APPHUB_QA_JAVA_MESSAGE", "java ok");
    System.out.println("language=java message=" + msg);
  }
}
