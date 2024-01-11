package save

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/http-server/middleware"
	apiResp "github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/lib/api/response"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/lib/logger/sl"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/lib/random"
	"github.com/khaydarovm/otus-golang-professional/sample_projects/url-shortener/internal/storage"
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	apiResp.Response
	Alias string `json:"alias,omitempty"`
}

//go:generate mockery --name=URLSaver
type URLSaver interface {
	SaveURL(url, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetRequestID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, apiResp.Error("failed to decode request"))

			return
		}

		if err = validator.New().Struct(req); err != nil {
			log.Error("invalid request", sl.Err(err))

			validatorErrs := err.(validator.ValidationErrors)
			render.JSON(w, r, apiResp.ValidationError(validatorErrs))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(6)
		}

		id, err := urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info("saveURL error", sl.Err(err))

			render.JSON(w, r, apiResp.Error("url already exist"))

			return
		}

		if err != nil {
			log.Error("failed to add url", sl.Err(err))

			render.JSON(w, r, apiResp.Error("failed to add url"))

			return
		}

		log.Info("url added", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Response: apiResp.OK(),
			Alias:    alias,
		})
	}
}
