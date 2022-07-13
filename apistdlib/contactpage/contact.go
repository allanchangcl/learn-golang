package contactpage

import (
	"log"
	"net/http"

	"apistdlib/middleware"
)

const (
	message  = "This is Contact Page"
	pagename = "contactpage"
)

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

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/contact", h.mwl.LogHandler(h.About, pagename))
}

func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(message))
	if err != nil {
		h.logger.Printf("Error: %s", err)
	}
}
