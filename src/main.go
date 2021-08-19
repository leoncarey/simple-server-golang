package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", ServerResponse)
    http.ListenAndServe(":3000", nil)
}
