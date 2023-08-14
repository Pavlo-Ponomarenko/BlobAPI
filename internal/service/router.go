package service

import (
	"blob-service/internal/service/handlers"
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
		r.Get("/blobs/{id}", handlers.GetBlobById)
		r.Get("/blobs/", handlers.GetPageOfBlobs)
		r.Post("/blobs/", handlers.CreateNewBlob)
		r.Delete("/blobs/{id}", handlers.DeleteBlob)
		r.Put("/blobs/{id}", handlers.UpdateBlob)
		r.Patch("/blobs/{id}", handlers.UpdateBlob)
	})

	return r
}
