package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Say Hello
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse arguments
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Samuel!") // Send data to client side
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil) // set listenening port
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
