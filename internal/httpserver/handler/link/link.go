package link

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/defryheryanto/url-shortener/internal/app"
	"github.com/defryheryanto/url-shortener/internal/errors"
	"github.com/defryheryanto/url-shortener/internal/httpserver/handler"
	"github.com/defryheryanto/url-shortener/internal/httpserver/response"
	"github.com/gorilla/mux"
)

func Shorten(application *app.Application) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		type shortenPayload struct {
			Url string `json:"url"`
		}

		var p *shortenPayload
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			if err == io.EOF {
				return errors.NewBadRequestError("Please fill data")
			}
			return err
		}

		if p.Url == "" {
			return errors.NewBadRequestError("Please fill URL")
		}

		newLink := application.LinkService.CreateLink(p.Url)
		response.WithData(w, http.StatusOK, newLink)
		return nil
	})
}

func GetURL(application *app.Application) http.HandlerFunc {
	return handler.Handle(func(w http.ResponseWriter, r *http.Request) error {
		id := mux.Vars(r)["id"]

		existingLink := application.LinkService.GetLink(id)
		if existingLink == nil || existingLink.Url == "" {
			return errors.NewNotFoundError("Link not found")
		}

		http.Redirect(w, r, existingLink.Url, http.StatusTemporaryRedirect)

		return nil
	})
}
