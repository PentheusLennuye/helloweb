package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	httpPort     = 8080
	readTimeout  = 60
	writeTimeout = 60
	idleTimeout  = 60
)

/*
 * helloweb is a ground-breaking technology that prints a message on a website
 */
func main() {
	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/", helloWebOut)
	runServer(muxHandler)
}

func helloWebOut(rw http.ResponseWriter, r *http.Request) {
	header := rw.Header()
	header.Add("Content-Type", "text/html; charset=utf-8")
	rw.Write([]byte("<html><body><p>Hello to almost everyone on the Web!</p></body></html>"))
}

/*
 * runServer(handler) maintains the service until SIG. Hard-coded for now
 */
func runServer(handler http.Handler) {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", httpPort),
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		IdleTimeout:  idleTimeout * time.Second,
		Handler:      handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
