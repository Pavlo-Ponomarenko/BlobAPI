package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

type BlobsQ interface {
	New() BlobsQ

	GetBlobById(id string) (*BlobEntity, error)
	GetBlobs(pageParams pgdb.OffsetPageParams) ([]BlobEntity, error)
	SaveBlob(blob *BlobEntity) (*BlobEntity, error)
	DeleteBlob(id string) error
	UpdateBlob(id string, blob *BlobEntity) (*BlobEntity, error)
}
