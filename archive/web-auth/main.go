package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func createToken(w http.ResponseWriter, r *http.Request) {
	token := uuid.New().String()
	body := fmt.Sprintf("token: %s \n", token)
	w.Write([]byte(body))
}

func getBookAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	books := map[string]string{
		"1449311601": "Ryan Boyd",
		"148425094X": "Yvonne Wilson",
		"1484220498": "Prabath Siriwarden",
	}
	body := fmt.Sprintf("Author: %s \n", books[id])
	w.Write([]byte(body))
}

func main() {
	// log.Println("Web Auth !!!")
	router := mux.NewRouter()
	router.HandleFunc("/v1/auth/token", createToken).Methods("GET")
	router.HandleFunc("/v1/book/{id}", getBookAuthor).Methods("GET")
	log.Println("Server started and listening on http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", router)
}
