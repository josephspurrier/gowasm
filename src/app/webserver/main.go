package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	address = flag.String("address", ":80", "listen address")
	dir     = flag.String("dir", "../static", "directory to serve")
)

// Handler will log the HTTP requests
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", Handler(fs))

	log.Printf("Listening on %v", *address)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
