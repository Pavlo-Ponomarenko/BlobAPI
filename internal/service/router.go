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
		r.Get("/blobs/{id}", requests.GetBlobById)
		r.Get("/blobs/", requests.GetPageOfBlobs)
		r.Post("/blobs/", requests.CreateNewBlob)
		r.Delete("/blobs/{id}", requests.DeleteBlob)
		r.Put("/blobs/{id}", requests.UpdateBlob)
		r.Patch("/blobs/{id}", requests.UpdateBlob)
	})

	return r
}
