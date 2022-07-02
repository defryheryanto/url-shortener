package httpserver

import (
	"net/http"

	"github.com/defryheryanto/url-shortener/internal/app"
	"github.com/defryheryanto/url-shortener/internal/httpserver/handler/link"
	"github.com/gorilla/mux"
)

func HandleRoutes(application *app.Application) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", link.Shorten(application)).Methods(http.MethodPost)
	r.HandleFunc("/{id}", link.GetURL(application)).Methods(http.MethodGet)

	return r
}
