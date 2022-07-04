package aboutpage

import (
	"apistdlib/middleware"
	"log"
	"net/http"
)

const message = "This is About Page"

// can have other dependency injected
type Handlers struct {
	logger *log.Logger
	mwl    *middleware.Handlers
}

// initiated in main
func NewHandlers(logger *log.Logger) *Handlers {
	mwl := middleware.NewLogHandlers(logger)
	return &Handlers{
		logger: logger,
		mwl:    mwl,
	}
}

// mw := middleware.NewLogHandlers(logger)

// func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		startTime := time.Now()
// 		// h.logger.Println("request processed")
// 		// defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
// 		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
// 		next.ServeHTTP(w, r)
// 		// next(w, r)
// 	}
// }
//

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	// mwl := mw.NewLogHandlers(h.logger)
	// mux.HandleFunc("/", h.Home)
	// mux.HandleFunc("/about", h.Logger(h.About))
	// mux.HandleFunc("/about", h.Logger(h.About))
	mux.HandleFunc("/about", h.mwl.LogHandler(h.About))
}

// func Home(w http.ResponseWriter, r *http.Request) {
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	// log.Println("request processed")
	// h.logger.Println("request processed")
	// h.logger.Println("request processed")
	// fmt.Println(message)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	// w.Write([]byte(message))
	if err != nil {
		// log.Printf("Error: %s", err)
		// h.logger.Printf("Error: %s", err)
		h.logger.Printf("Error: %s", err)
		// fmt.Println("Error")
	}
}
