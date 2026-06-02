package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        msg := os.Getenv("APPHUB_QA_GO_MESSAGE")
        if msg == "" { msg = "go ok" }
        _, _ = fmt.Fprintf(w, "language=go message=%s", msg)
    })
    port := os.Getenv("PORT")
    if port == "" { port = "8080" }
    _ = http.ListenAndServe(":"+port, nil)
}
