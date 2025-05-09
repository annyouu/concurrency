package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, custom handler!")
}

func main() {
	http.ListenAndServe(":8080", HelloHandler{})
}

