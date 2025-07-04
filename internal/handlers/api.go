package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/", func(r chi.Router) {
		r.Get("/", Health)

		r.Get("/send_hello", GetHello)

		r.Get("/webhook", GetWebhook)

		r.Post("/webhook", PostWebhook)
	})

}
