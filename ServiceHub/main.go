package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/portainer", getPortainerHandler).Methods("GET")
	r.HandleFunc("/portainer/auth", createPortainerHandler).Methods("POST")
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World!")
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", r)

}
