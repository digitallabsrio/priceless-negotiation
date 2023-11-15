package main
import (
"fmt"
"log"
"net/http"
"os"
)

func main() {
  http.HandleFunc("/", indexHandler)
  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
    log.Printf("Defaulting to port %s", port)
  }
  log.Printf("Listening on port %s", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  fmt.Fprint(w, "Hello, World!")
}