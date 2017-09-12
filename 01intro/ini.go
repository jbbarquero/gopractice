package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", myNameIs)
	http.ListenAndServe("localhost:4000", nil)
}

func myNameIs(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya, you kill my father, prepare to die")
}
