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
	http.ListenAndServe(fmt.Sprintf(":%s", port), createRouter())
}

func createRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", indexHandler)


	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request){
		
	w.Write([]byte("welcome people to the world"))
}
