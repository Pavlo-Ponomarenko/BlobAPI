package pg

import (
	"blob-service/internal/data"
	"blob-service/internal/data/horizon"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type coreBlobsQ struct {
	adminSeed   string
	getBlobsURL string
	coreInfoURL string
}

func NewCoreBlobsQ(adminSeed string, getBlobsURL string, coreInfoURL string) data.BlobsQ {
	return &coreBlobsQ{
		adminSeed:   adminSeed,
		getBlobsURL: getBlobsURL,
		coreInfoURL: coreInfoURL,
	}
}

func (q *coreBlobsQ) New() data.BlobsQ {
	return NewCoreBlobsQ(q.adminSeed, q.getBlobsURL, q.coreInfoURL)
}

func (q *coreBlobsQ) GetBlobById(id string) (*data.BlobEntity, error) {
	return horizon.GetBlobById(id, q.getBlobsURL)
}

func (q *coreBlobsQ) GetBlobs(pageParams pgdb.OffsetPageParams) ([]data.BlobEntity, error) {
	return horizon.GetBlobs(pageParams, q.adminSeed, q.getBlobsURL)
}

func (q *coreBlobsQ) SaveBlob(blob *data.BlobEntity) (*data.BlobEntity, error) {
	newBlob, err := horizon.CreateBlob(blob, q.adminSeed, q.coreInfoURL)
	if err != nil {
		return nil, err
	}
	return newBlob, err
}

func (q *coreBlobsQ) DeleteBlob(id string) error {
	return horizon.DeleteBlob(id, q.adminSeed, q.coreInfoURL)
}

func (q *coreBlobsQ) UpdateBlob(id string, blob *data.BlobEntity) (*data.BlobEntity, error) {
	return horizon.UpdateBlob(id, blob, q.adminSeed, q.getBlobsURL)
}
