package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/sirupsen/logrus"

)

func main() {
	

	port := "8080"
	logrus.Info("starting host on port: " + port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), createRouter())
	if err != nil {
		logrus.Fatalf("error: httpListenAndServe: %v", err)
	}
}

func createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", indexHandler)


	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request){
		
	n, err := w.Write([]byte("welcome people to the world"))
	if err != nil {
		logrus.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	logrus.Printf("wrote %d bytes to response", n)
}
