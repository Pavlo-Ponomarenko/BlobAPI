package service

import (
	"blob-service/internal/service/handlers"
	"blob-service/internal/service/requests"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)
	r.Route("/blob-service", func(r chi.Router) {
		// configure endpoints here
		r.Get("/blob/{id}", requests.Get_blob_by_id)
		r.Get("/blobs", requests.Get_page_of_blobs)
		r.Post("/blob/create", requests.Create_new_blob)
		r.Delete("/blob/{id}", requests.Delete_blob)
	})

	return r
}
