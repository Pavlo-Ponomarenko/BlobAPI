package horizon

import (
	"blob-service/internal/data"
	"encoding/json"
)

type apiData struct {
	Key
	Attributes    DataAttributes    `json:"attributes"`
	Relationships DataRelationships `json:"relationships"`
}

type apiDataList struct {
	Data     []apiData         `json:"data"`
	Included []json.RawMessage `json:"included"`
	Links    *Links            `json:"links"`
	Meta     json.RawMessage   `json:"meta,omitempty"`
}

type apiDataById struct {
	Data     apiData           `json:"data"`
	Included []json.RawMessage `json:"included"`
}

type Key struct {
	ID   string       `json:"id"`
	Type ResourceType `json:"type"`
}

type ResourceType string

type DataAttributes struct {
	// is used to restrict using of data through rules
	Type  uint64          `json:"type"`
	Value json.RawMessage `json:"value"`
}

type DataRelationships struct {
	Owner *Relation `json:"owner,omitempty"`
}

type Relation struct {
	Data  *Key   `json:"data,omitempty"`
	Links *Links `json:"links,omitempty"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	Self  string `json:"self"`
}

func apiDataToEntity(blobData apiData) *data.BlobEntity {
	blobEntity := new(data.BlobEntity)
	rawBlob := blobData.Attributes.Value
	_ = json.Unmarshal(rawBlob, blobEntity)
	return blobEntity
}

func apiDataListToEntities(blobDataList apiDataList) []data.BlobEntity {
	entities := make([]data.BlobEntity, 0, 20)
	for i := range blobDataList.Data {
		entity := apiDataToEntity(blobDataList.Data[i])
		entities = append(entities, *entity)
	}
	return entities
}
