package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "hello %s", msg)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "potato")
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(MyHandler)))
}
