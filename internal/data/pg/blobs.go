package pg

import (
	"blob-service/internal/data"
	"blob-service/internal/data/horizon"
	"fmt"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type blobsQ struct {
	db *pgdb.DB
}

func NewBlobsQ(db *pgdb.DB) data.BlobsQ {
	return &blobsQ{}
}

func (q *blobsQ) New() data.BlobsQ {
	return NewBlobsQ(q.db)
}

func (q *blobsQ) GetBlobById(id string) (*data.BlobEntity, error) {
	return horizon.GetBlobById(id)
}

func (q *blobsQ) GetBlobs(pageParams pgdb.OffsetPageParams) ([]data.BlobEntity, error) {
	return horizon.GetBlobs(pageParams)
}

func (q *blobsQ) SaveBlob(blob *data.BlobEntity) (*data.BlobEntity, error) {
	newBlob, err := horizon.CreateBlob(blob)
	if err != nil {
		fmt.Println("Saving to blockchain failed: ", err)
		return nil, err
	}
	return newBlob, err
}

func (q *blobsQ) DeleteBlob(id string) error {
	return horizon.DeleteBlob(id)
}

func (q *blobsQ) UpdateBlob(id string, blob *data.BlobEntity) (*data.BlobEntity, error) {
	return horizon.UpdateBlob(id, blob)
}
