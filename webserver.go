/*
 - / - index.html
 - /hello - hellofunc
 -/form - formfunc -> form.html

*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "via Hello func")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404, not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")

	fmt.Fprintf(w, firstName+lastName)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
