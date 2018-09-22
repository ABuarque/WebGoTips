package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// to add to header
	w.Header().Set("VAI", "Tame")

	// return response with string
	//io.WriteString(w, `{"alive": true}`)

	//return response with json
	user := user{Name: "Aurelio", Email: "abmf"}
	json.NewEncoder(w).Encode(&user)

	// define response code at header
	w.WriteHeader(http.StatusOK)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken == "vaca" {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/v1/users", getAll).Methods("GET")

	router.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router))
}

