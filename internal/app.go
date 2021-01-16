package internal

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Welcome http.HandlerFunc
}

func WelcomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request... Processing")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Welcome to serverless"}`))
		log.Println("Response has been returned")
	}
}

func New() App {

	w := WelcomeHandler()

	return App{
		Welcome: w,
	}
}

func (a *App) Handler() http.HandlerFunc {
	router := mux.NewRouter()

	router.Handle("/welcome", a.Welcome)
	// router.Handle("/register", )
	h := http.HandlerFunc(router.ServeHTTP)
	return h
}
