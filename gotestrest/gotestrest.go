package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/thisisdavidbell/go-examples/queue"
)

type messagestruct struct {
  Message string
}

func pushhandler(w http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var m messagestruct
    err := decoder.Decode(&m);

    if err != nil {
    		panic(err)
    }

    var message = m.Message

    if (message == "") {
      //TODO ERROR case
      message="DIDNTGETMESSAGEFROMBODY"
    }

    queue.Push(message)
    fmt.Fprintf(w, "Pushed %s", message)
}

func pullhandler(w http.ResponseWriter, req *http.Request) {
    var message = queue.Pull()
    jsonmessage := messagestruct{Message: message}
    enc := json.NewEncoder(w)
    enc.Encode(jsonmessage)

//    fmt.Fprintf(w, message)
}

func main() {
   queue.Start();
   http.HandleFunc("/push", pushhandler)
   http.HandleFunc("/pull", pullhandler)
    // prevent accept incoming requests popup on mac os for dev: http.ListenAndServe("localhost:8080", nil)
   http.ListenAndServe(":8080", nil)
}
