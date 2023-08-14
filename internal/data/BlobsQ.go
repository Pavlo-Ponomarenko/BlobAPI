package data

import (
	res "blob-service/resources"
	"database/sql"
)

type IBlobsQ interface {
	Close()

	GetBlobById(id string) (*res.BlobModel, error)
	GetBlobs(params map[string]string) ([]res.BlobModel, error)
	SaveBlob(blob *res.Blob) error
	DeleteBlob(id string) error
	IdIsPresent(id string) bool
	UpdateBlob(id string, blob *res.Blob) error
}

type BlobsQ struct {
	db *sql.DB
}
