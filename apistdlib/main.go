package main

import (
	"apistdlib/aboutpage"
	"apistdlib/homepage"
	"apistdlib/server"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// const message = "Hello"

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// godotenv.Load()
		// number := 12

	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://localhost:4343"+r.RequestURI, http.StatusMovedPermanently)
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// fmt.Println(message)
// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	_, err := w.Write([]byte(message))
// 	// w.Write([]byte(message))
// 	if err != nil {
// 		log.Printf("Error: %s", err)
// 		// fmt.Println("Error")
// 	}
// }

func main() {
	// const message = "Hello"
	// number := 12
	var (
		ServiceAddr    = os.Getenv("SERVICE_ADDR")
		ServiceAddrSSL = os.Getenv("SERVICE_ADDR_SSL")
		CertFile       = os.Getenv("CERT_FILE")
		KeyFile        = os.Getenv("KEY_FILE")
	)

	// fmt.Println(ServiceAddr)

	logger := log.New(os.Stdout, "apistdlib", log.LstdFlags|log.Lshortfile)

	// pass dependency to HomeHandler
	h := homepage.NewHandlers(logger)
	a := aboutpage.NewHandlers(logger)

	mux := http.NewServeMux()
	// mux.HandleFunc("/", HomeHandler)
	// mux.HandleFunc("/", homepage.Home)

	// mux.HandleFunc("/", h.Home)
	h.SetupRoutes(mux)
	a.SetupRoutes(mux)
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Println(message)
	// 	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// 	w.WriteHeader(http.StatusOK)
	// 	_, err := w.Write([]byte(message))
	// 	// w.Write([]byte(message))
	// 	if err != nil {
	// 		log.Printf("Error: %s", err)
	// 		// fmt.Println("Error")
	// 	}
	// })

	// srv := NewServer(mux, ServiceAddr)
	// srv := NewServer(mux, ServiceAddr)
	// srv := NewServer(mux, ServiceAddrSSL)
	srv := server.New(mux, ServiceAddrSSL)
	// srvSSL := NewServer(mux, ServiceAddrSSL)
	// srvSSL := NewServer(mux, ServiceAddrSSL)

	// go func() {
	// 	err := srv.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatalf("Server failed to start: %v", err)
	// 		// log.Printf("Error: %s", err)
	// 	}
	// }()

	go func() {
		err := http.ListenAndServe(ServiceAddr, http.HandlerFunc(redirectTLS))
		if err != nil {
			logger.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	logger.Printf("Server started on %s", ServiceAddr)

	// err := http.ListenAndServe(":8080", mux)
	// err := srv.ListenAndServe()
	// err := srvSSL.ListenAndServeTLS(CertFile, KeyFile)
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	// srvSSL.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
		// log.Printf("Error: %s", err)
	}
	// fmt.Println(message)
}
