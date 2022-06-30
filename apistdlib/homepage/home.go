package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello"

type Handlers struct {
	logger *log.Logger
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{logger: logger}
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		// h.logger.Println("request processed")
		// defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
		next.ServeHTTP(w, r)
		// next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	// mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/", h.Logger(h.Home))
}

// func Home(w http.ResponseWriter, r *http.Request) {
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	// log.Println("request processed")
	// h.logger.Println("request processed")
	// fmt.Println(message)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	// w.Write([]byte(message))
	if err != nil {
		// log.Printf("Error: %s", err)
		h.logger.Printf("Error: %s", err)
		// fmt.Println("Error")
	}
}
