// main.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
	fmt.Fprintf(w, "Current Date & Time: %s\n", currentTime)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
