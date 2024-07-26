package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dc := gg.NewContext(800, 800)
	hostname, _ := os.Hostname()
	seed := time.Now().UnixNano() + int64(hostname[0])
	drawRandomArt(dc, seed)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
