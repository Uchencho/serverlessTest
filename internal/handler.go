package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func WelcomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request... Processing")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Welcome to serverless"}`))
		log.Println("Response has been returned")
	}
}

func RegisterHandler(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var pl RegisterUserRequest
		err := json.NewDecoder(req.Body).Decode(&pl)
		if err != nil {
			log.Println("Handler - ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error" : "malinformed json"}`))
			return
		}

		reg := Register(c)
		err = reg(pl.Email, pl.Password)

		if err != nil {
			log.Println("Error in registering user, ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error" : "Something went wrong"}`))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(`{"message" : "Successful registeration"}`))
	}
}

func LoginHandler(c Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var pl RegisterUserRequest
		err := json.NewDecoder(req.Body).Decode(&pl)
		if err != nil {
			log.Println("Handler - ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error" : "malinformed json"}`))
			return
		}

		loginUser := Login(c)
		token, err := loginUser(pl.Email, pl.Password)

		if err != nil {
			log.Println("Error in loggin in user, ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error" : "Something went wrong"}`))
			return
		}

		bb, err := json.Marshal(token)
		if err != nil {
			log.Println("Handler - ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error" : "could not marshal token"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bb)
	}
}
