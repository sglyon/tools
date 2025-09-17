package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Parse command line flags
	port := flag.String("port", "8080", "port to serve on")
	dir := flag.String("dir", ".", "directory to serve")
	flag.Parse()

	// Create file server
	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	// Start server
	addr := fmt.Sprintf(":%s", *port)
	fmt.Printf("Serving tools at http://localhost%s\n", addr)
	fmt.Printf("Serving directory: %s\n", *dir)
	log.Fatal(http.ListenAndServe(addr, nil))
}