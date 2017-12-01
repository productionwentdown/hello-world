package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! Here's some info:\n\n")
	time, _ := time.Now()
	hostname, _ := os.Hostname()
	environ, _ := os.Environ()
	uid, _ := os.Getuid()
	gid, _ := os.Getgid()
	fmt.Fprintf(w, "Time: %v", time)
	fmt.Fprintf(w, "Hostname: %v\n", hostname)
	fmt.Fprintf(w, "Environment: %v\n", environ)
	fmt.Fprintf(w, "UID: %v GID: %v\n", uid, gid)
}

func main() {
	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
