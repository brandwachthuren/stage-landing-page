package main

import (
    "fmt"
	"net/http"
	"os"
	"log"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not set, default 9999")
		port = "9999"
	}
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "bwh stage %s", r.URL.Path[1:])
}