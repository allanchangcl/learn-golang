package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	// ...
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, s)
	w.Write([]byte("hello"))

}

// func mylog(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Before")
// 		h.ServeHTTP(w, r) // call original
// 		log.Println("After")

// 	})
// }

// func hello(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "hello\n")
// }

// func (h myHandler) ServeHttp(w http.ResponseWriter, r *http.Response) {
// 	w.Write([]byte("hello"))
// }

func logger(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		f.ServeHTTP(w, r) // call original
		log.Println("After")

	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	// http.Handle("/hello", server{})
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/log", logger(hello))

	http.ListenAndServe(":8090", nil)
}
