package middleware

import (
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
	// Logger *log.Logger
}

func NewLogHandlers(logger *log.Logger) *Handlers {
	// return &Handlers{logger: logger}
	return &Handlers{logger: logger}
}

// Logger returns a middleware that logs HTTP requests.
// func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
func (h *Handlers) LogHandler(next http.HandlerFunc, page string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		// h.logger.Println("request processed")
		// defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		defer h.logger.Printf("%s request processed in %s\n", page, time.Since(startTime))
		next.ServeHTTP(w, r)
		// next(w, r)
	}
}
