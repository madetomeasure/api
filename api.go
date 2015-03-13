package main

import (
	"github.com/madetomeasure/api/endpoints/subscribers"
        //"github.com/madetomeasure/api/endpoints/lists"
        //"github.com/madetomeasure/api/endpoints/campaigns"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	consoleStatus()

        r := startRouter()
	http.ListenAndServe(":"+port(), r)
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func startRouter() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./public/")))
	r.HandleFunc("/subscribers", subscriber.RequestHandler)
	r.HandleFunc("/welcome", NotImplementedHandler)
        return r
}

func consoleStatus() {
	fmt.Println("=> Booting...")
	fmt.Println("=> Made to Measure Alpha application starting in development on http://0.0.0.0:" + port())
	fmt.Println("=> Ctrl-C to shutdown Server")
}

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
