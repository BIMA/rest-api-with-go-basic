package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Routing static assets
	http.Handle(
		"/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)

	// If you want to use http.Server
	//server := new(http.Server)
	//server.Addr = address
	//err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	_, err := w.Write([]byte(message))
	if err != nil {
		return
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	_, err := w.Write([]byte(message))
	if err != nil {
		return
	}
}
