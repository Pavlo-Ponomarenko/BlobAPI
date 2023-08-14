package handlers

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

func getDb() (*sql.DB, error) {
	connection := "postgres://user:postgres@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("DB connection error")
	}
	return db, err
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

func GetBlobById(id string) (*res.BlobModel, error) {
	db, err := getDb()
	if err != nil {

		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(getBlob, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return nil, err
	}
	if !rows.Next() {
		fmt.Println("No blob with such id")
		return nil, errors.New("")
	}
	defer rows.Close()
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

func GetBlobs(params map[string]string) (*res.BlobModelListResponse, error) {
	query, err := formQuery(params)
	if err != nil {
		return nil, err
	}
	db, err := getDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(*query)
	blobs := make([]res.BlobModel, 0, 20)
	for rows.Next() {
		blob, _ := retrieveBlob(rows)
		blobs = append(blobs, *blob)
	}
	blobResponse := new(res.BlobModelListResponse)
	blobResponse.Data = blobs
	return blobResponse, nil
}

func SaveBlob(blob *res.Blob) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()
	jsonRepr, err := json.Marshal(blob.Data.Attributes.Value)
	if err != nil {
		fmt.Println("Attributes encoding error")
		return err
	}
	_, err = db.Exec(saveBlob, blob.Data.ID, jsonRepr)
	if err != nil {
		fmt.Println("SQL execution error:", err)
	}
	return err
}

func DeleteBlob(id string) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Query(deleteBlob, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return err
	}
	return nil
}

func IdIsPresent(id string) bool {
	db, err := getDb()
	if err != nil {
		return false
	}
	defer db.Close()
	rows, err := db.Query(getBlob, id)
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

func UpdateBlob(id string, blob *res.Blob) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()
	jsonRepr, err := json.Marshal(blob.Data.Attributes.Value)
	_, err = db.Query(updateBlob, jsonRepr, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return err
	}
	return nil
}
