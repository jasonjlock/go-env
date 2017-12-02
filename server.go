package main

import (
    "fmt"
    "net/http"
)

const port = "8080"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Go is running...")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + port, nil)
}
