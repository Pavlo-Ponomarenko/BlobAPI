package horizon

import (
	"blob-service/internal/data"
	"encoding/json"
)

type blobToJSON struct {
	data.BlobEntity
}

func (b blobToJSON) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.BlobEntity)
}
