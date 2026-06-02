package main
import (
  "fmt"
  "net/http"
  "os"
)
func main(){ port:=os.Getenv("PORT"); if port==""{port="8080"}; http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){fmt.Fprint(w,"axhub go http ok")}); http.ListenAndServe("0.0.0.0:"+port,nil)}
