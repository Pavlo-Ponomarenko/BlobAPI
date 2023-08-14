package data

import (
	res "blob-service/resources"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var getBlob = "select * from blobs where id = $1"
var saveBlob = "insert into blobs (id, blob) values ($1, $2)"
var deleteBlob = "delete from blobs where id = $1"
var getBlobsPage = "select * from blobs"
var updateBlob = "update blobs set blob = $1 where id = $2"

func CreateNewBlobsQ() (IBlobsQ, error) {
	connection := "postgres://user:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("DB connection error")
	}
	newQ := BlobsQ{db: db}
	return &newQ, err
}

func (q *BlobsQ) Close() {
	q.db.Close()
}

func retrieveBlob(rows *sql.Rows) (*res.BlobModel, error) {
	blob := new(res.BlobModel)
	blob.Type = "blob"
	bytea := make([]byte, 0, 1000)
	err := rows.Scan(&blob.ID, &bytea)
	if err != nil {
		return nil, err
	}
	attributeMap := make(map[string]interface{})
	err = json.Unmarshal(bytea, &attributeMap)
	if err != nil {
		return nil, err
	}
	blob.Attributes.Value = &attributeMap
	return blob, nil
}

func (q *BlobsQ) GetBlobById(id string) (*res.BlobModel, error) {
	rows, err := q.db.Query(getBlob, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return nil, err
	}
	if !rows.Next() {
		fmt.Println("No blob with such id")
		return nil, errors.New("")
	}
	return retrieveBlob(rows)
}

func formQuery(params map[string]string) (*string, error) {
	query := getBlobsPage
	order, exists := params["order"]
	if exists && (order == "asc" || order == "desc") {
		query = query + " order by id " + order
	}
	limit, exists := params["limit"]
	if exists {
		limitValue, err := strconv.Atoi(limit)
		if err != nil || limitValue < 0 {
			return nil, err
		}
		query = query + " limit " + limit
		number, exists := params["number"]
		if exists {
			numberValue, err := strconv.Atoi(number)
			if err != nil || numberValue < 1 {
				return nil, errors.New("")
			}
			query = query + " offset " + strconv.Itoa(limitValue*(numberValue-1))
		}
	}
	return &query, nil
}

func (q *BlobsQ) GetBlobs(params map[string]string) ([]res.BlobModel, error) {
	query, err := formQuery(params)
	if err != nil {
		return nil, err
	}
	rows, err := q.db.Query(*query)
	blobs := make([]res.BlobModel, 0, 20)
	for rows.Next() {
		blob, _ := retrieveBlob(rows)
		blobs = append(blobs, *blob)
	}
	return blobs, nil
}

func (q *BlobsQ) SaveBlob(blob *res.Blob) error {
	jsonRepr, err := json.Marshal(blob.Data.Attributes.Value)
	if err != nil {
		fmt.Println("Attributes encoding error")
		return err
	}
	_, err = q.db.Exec(saveBlob, blob.Data.ID, jsonRepr)
	if err != nil {
		fmt.Println("SQL execution error:", err)
	}
	return err
}

func (q *BlobsQ) DeleteBlob(id string) error {
	_, err := q.db.Query(deleteBlob, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return err
	}
	return nil
}

func (q *BlobsQ) IdIsPresent(id string) bool {
	rows, err := q.db.Query(getBlob, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return false
	}
	if !rows.Next() {
		fmt.Println("No blob with such id")
		return false
	}
	return true
}

func (q *BlobsQ) UpdateBlob(id string, blob *res.Blob) error {
	jsonRepr, err := json.Marshal(blob.Data.Attributes.Value)
	_, err = q.db.Query(updateBlob, jsonRepr, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return err
	}
	return nil
}
