package main

import (
  "fmt"
  "net/http"
)

func roothandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Im the root. You should try /hello")
}

func hwhandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World. Im written in Go")
}

func main() {
   http.HandleFunc("/", roothandler)
    http.HandleFunc("/hello", hwhandler)
    // prevent accept incoming requests popup on mac os for dev: http.ListenAndServe("localhost:8080", nil)
    http.ListenAndServe(":8080", nil)
}
