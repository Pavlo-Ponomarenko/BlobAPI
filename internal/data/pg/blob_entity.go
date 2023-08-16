package pg

import (
	res "blob-service/resources"
	"encoding/json"
)

type blobEntity struct {
	Id   string           `db:"id"`
	Blob *json.RawMessage `db:"blob"`
}

func blobToEntity(blob *res.Blob) *blobEntity {
	entity := new(blobEntity)
	entity.Id = blob.ID
	entity.Blob = blob.Attributes.Value
	return entity
}

func entityToBlob(entity *blobEntity) *res.Blob {
	blob := new(res.Blob)
	blob.ID = entity.Id
	blob.Type = "blob"
	blob.Attributes.Value = entity.Blob
	return blob
}

func entitiesToBlobs(entities []blobEntity) []res.Blob {
	blobs := make([]res.Blob, 0, 20)
	for i := range entities {
		blobs = append(blobs, *entityToBlob(&entities[i]))
	}
	return blobs
}
