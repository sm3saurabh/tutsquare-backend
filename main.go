package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "log"
)

func startTutSquare(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("This is the first method of tutsquare, feel proud"))
}

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/start", startTutSquare)

  log.Fatal(http.ListenAndServe(":8080", router))
}
