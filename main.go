package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/hello", handler2)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<p>Go is fast %s ....%s simple</p>", "can", "<strong>variables</strong>")

}

func handler1(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "welcome")
}
