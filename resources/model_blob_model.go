/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type BlobModel struct {
	Key
	Attributes BlobModelAttributes `json:"attributes"`
}
type BlobModelResponse struct {
	Data     BlobModel `json:"data"`
	Included Included  `json:"included"`
}

type BlobModelListResponse struct {
	Data     []BlobModel `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustBlobModel - returns BlobModel from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustBlobModel(key Key) *BlobModel {
	var blobModel BlobModel
	if c.tryFindEntry(key, &blobModel) {
		return &blobModel
	}
	return nil
}
