package main

import "fmt"
import "net/http"

type Hello struct {}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, "Hello!")
	}

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (h String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h)
}

func (h *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h)
}

func (h Hello) Start(host string) {
	http.ListenAndServe(host, h)
}

func (h Hello) Start2(host string) {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"hello", ":", "gophers!"})
	
	http.ListenAndServe(host, nil)
}
