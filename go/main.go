// main.go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Applicacion Go levantada con Docker")
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Servidor escuchando en http://localhost:8080")
    http.ListenAndServe(":8081", nil)
}
