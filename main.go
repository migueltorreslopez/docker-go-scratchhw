package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received call in hellohandler")
	fmt.Fprintln(w, "Hello World from Go in minimal Docker container")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Started, serving at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
