package service

import (
	"blob-service/internal/config"
	"blob-service/internal/data/pg"
	"blob-service/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/kit/kv"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	conf := config.New(kv.MustFromEnv())

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxBlobsQ(pg.NewBlobsQ(conf.DB())),
		),
	)
	r.Route("/blob-service", func(r chi.Router) {
		// configure endpoints here
		r.Get("/blobs/{id}", handlers.GetBlobById)
		r.Delete("/blobs/{id}", handlers.DeleteBlob)
		r.Put("/blobs/{id}", handlers.UpdateBlob)

		r.Get("/blobs/", handlers.GetPageOfBlobs)
		r.Post("/blobs/", handlers.CreateNewBlob)
	})

	return r
}
