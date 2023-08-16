package data

import (
	res "blob-service/resources"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type BlobsQ interface {
	New() BlobsQ

	GetBlobById(id string) (*res.Blob, error)
	GetBlobs(pageParams pgdb.OffsetPageParams) ([]res.Blob, error)
	SaveBlob(blob *res.Blob) (*res.Blob, error)
	DeleteBlob(id string) error
	IdIsPresent(id string) bool
	UpdateBlob(id string, blob *res.Blob) (*res.Blob, error)
}
