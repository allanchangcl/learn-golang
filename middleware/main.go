package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret"))
var isDebug = true

func debug(val interface{}, isDebug bool) {
	if isDebug {
		log.Println(val)
	}
}

func main() {
	fmt.Println("vim-go")
	// http.HandleFunc("/", indexServer)
	http.HandleFunc("/login", login)
	http.HandleFunc("/home", Auth(home))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

// protected
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// get access
func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// start authentication

	// set user as authenticated
	session.Values["userID"] = 12345678
	session.Save(r, w)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home, %s!", r.URL.Path[1:])
}

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session, _ := store.Get(r, "session")
		session, _ := store.Get(r, "session")
		userID, ok := session.Values["userID"]
		debug(ok, isDebug)
		debug(userID, isDebug)
		if !ok {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}
		HandlerFunc.ServeHTTP(w, r)
	}
}
