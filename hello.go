package main

import (
	"fmt"
	"net/http"
	"os"
    "runtime"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! Here's some info:\n\n")
	time := time.Now()
    numCPU := runtime.NumCPU()
    version := runtime.Version()
	hostname, _ := os.Hostname()
	environ := os.Environ()
	uid := os.Getuid()
	gid := os.Getgid()
	fmt.Fprintf(w, "Time: %v\n", time)
	fmt.Fprintf(w, "CPUs: %v\n", numCPU)
	fmt.Fprintf(w, "Go version: %v\n", version)
	fmt.Fprintf(w, "Hostname: %v\n", hostname)
	fmt.Fprintf(w, "Environment: %v\n", environ)
	fmt.Fprintf(w, "UID: %v GID: %v\n", uid, gid)
}

func main() {
	fmt.Println("Listening on port 8080...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
