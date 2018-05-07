package main

import (
	"io"
	"net/http"
	"log"
	"fmt"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {

	fmt.Println("hello, Vi L. !!")
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}