package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, str string) {
	// Anything that implements io.Writer can be passed in as w
	// e.g. os.Stdout, http.ResponseWriter, bytes.Buffer
	// This is dependency injection somehow
	fmt.Fprintf(w, "hello %s", str)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "potato")
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(MyGreeterHandler)))
}
