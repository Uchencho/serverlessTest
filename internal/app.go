package internal

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	Welcome  http.HandlerFunc
	Register http.HandlerFunc
	Login    http.HandlerFunc
}

func New() App {

	c := Config{
		Region:      "us-east-1",
		UserPoolID:  os.Getenv("USER-POOL-ID"),
		AppClientID: os.Getenv("APP-CLIENT-ID"),
	}

	w := WelcomeHandler()
	reg := RegisterHandler(c)
	login := LoginHandler(c)

	return App{
		Welcome:  w,
		Register: reg,
		Login:    login,
	}
}

func (a *App) Handler() http.HandlerFunc {
	router := mux.NewRouter()

	router.Handle("/welcome", a.Welcome)
	router.Handle("/register", a.Register)
	router.Handle("/login", a.Login)
	h := http.HandlerFunc(router.ServeHTTP)
	return h
}
