package save

import (
	resp "go-url-shortener/internal/lib/api/response"
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json: "url" validate: "required, url"`
	Alias string `json: "alias, omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json: "alias, omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
