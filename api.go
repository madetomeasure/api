package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	welcome()
	http.ListenAndServe(":"+port(), r)
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func welcome() {
	fmt.Println("=> Booting...")
	fmt.Println("=> Made to Measure Alpha application starting in development on http://0.0.0.0:" + port())
	fmt.Println("=> Ctrl-C to shutdown Server")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "Hello from Made to Measure")
}
