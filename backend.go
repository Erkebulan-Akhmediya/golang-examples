package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myHandler struct{}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/hello", myHandler{})
	http.HandleFunc("/home", home)
	http.HandleFunc("/google", google)
	http.HandleFunc("/v2/google", googlev2)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is home page")
}

func index(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, err = io.Copy(w, file); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func google(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://google.com", 200)
}

func googlev2(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("google.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, err = io.Copy(w, file); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
