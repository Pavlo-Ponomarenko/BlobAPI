package data

import (
	res "blob-service/resources"
	"encoding/json"
)

type BlobEntity struct {
	Id   string           `db:"id"`
	Blob *json.RawMessage `db:"blob"`
}

func BlobToEntity(blob *res.Blob) *BlobEntity {
	entity := new(BlobEntity)
	entity.Id = blob.ID
	entity.Blob = blob.Attributes.Value
	return entity
}

func EntityToBlob(entity *BlobEntity) *res.Blob {
	blob := new(res.Blob)
	blob.ID = entity.Id
	blob.Type = res.BLOB
	blob.Attributes.Value = entity.Blob
	return blob
}

func EntitiesToBlobs(entities []BlobEntity) []res.Blob {
	blobs := make([]res.Blob, 0, 20)
	for i := range entities {
		blobs = append(blobs, *EntityToBlob(&entities[i]))
	}
	return blobs
}
