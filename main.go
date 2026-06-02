package main
import (
  "fmt"
  "net/http"
  "os"
  "github.com/go-chi/chi/v5"
)
func main(){ port:=os.Getenv("PORT"); if port==""{port="8080"}; r:=chi.NewRouter(); r.Get("/", func(w http.ResponseWriter, r *http.Request){fmt.Fprint(w,"axhub go chi ok")}); http.ListenAndServe("0.0.0.0:"+port,r)}
