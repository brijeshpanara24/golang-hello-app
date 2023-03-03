package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
 
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "hello world")
    }) 
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    } 
    fmt.Printf("Starting server at %s\n", port)
    http.ListenAndServe(":"+port, nil)
}