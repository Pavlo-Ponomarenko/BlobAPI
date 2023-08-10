package handlers

import (
	res "blob-service/resources"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var get_blob = "select * from blobs where id = $1"
var save_blob = "insert into blobs (blob) values ($1)"
var delete_blob = "delete from blobs where id = $1"
var get_blobs_page = "select * from blobs"

func get_DB() (*sql.DB, error) {
	connection := "postgres://user:postgres@localhost/postgres?sslmode=disable"
	return sql.Open("postgres", connection)
}

func retrieve_blob(rows *sql.Rows) (*res.Blob, error) {
	blob := new(res.Blob)
	blob.Type = "blob"
	bytea := make([]byte, 0, 1000)
	err := rows.Scan(&blob.ID, &bytea)
	if err != nil {
		return nil, err
	}
	attributes := make(map[string]interface{})
	err = json.Unmarshal(bytea, &attributes)
	if err != nil {
		return nil, err
	}
	blob.Attributes = attributes
	return blob, nil
}

func Get_blob_by_id(id int) (*res.Blob, error) {
	db, err := get_DB()
	if err != nil {
		fmt.Println("DB connection error")
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(get_blob, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return nil, err
	}
	if !rows.Next() {
		fmt.Println("No blob with such id")
		return nil, errors.New("")
	}
	defer rows.Close()
	return retrieve_blob(rows)
}

func form_query(params map[string]string) (*string, error) {
	query := get_blobs_page
	order, exists := params["order"]
	if exists && (order == "asc" || order == "desc") {
		query = query + " order by id " + order
	}
	limit, exists := params["limit"]
	if exists {
		limit_value, err := strconv.Atoi(limit)
		if err != nil || limit_value < 0 {
			return nil, err
		}
		query = query + " limit " + limit
		number, exists := params["number"]
		if exists {
			number_value, err := strconv.Atoi(number)
			if err != nil || number_value < 0 {
				return nil, err
			}
			query = query + " offset " + strconv.Itoa(limit_value*(number_value-1))
		}
	}
	return &query, nil
}

func Get_blobs(params map[string]string) (*res.BlobListResponse, error) {
	query, err := form_query(params)
	if err != nil {
		return nil, err
	}
	db, err := get_DB()
	if err != nil {
		fmt.Println("DB connection error")
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(*query)
	blobs := make([]res.Blob, 0, 20)
	for rows.Next() {
		blob, _ := retrieve_blob(rows)
		blobs = append(blobs, *blob)
	}
	blob_response := new(res.BlobListResponse)
	blob_response.Data = blobs
	return blob_response, nil
}

func Save_blob(blob *res.Blob) error {
	db, err := get_DB()
	if err != nil {
		fmt.Println("DB connection error")
		return err
	}
	defer db.Close()
	json_repr, err := json.Marshal(blob.Attributes)
	if err != nil {
		fmt.Println("Attributes encoding error")
		return err
	}
	_, err = db.Exec(save_blob, json_repr)
	if err != nil {
		fmt.Println("SQL execution error:", err)
	}
	return err
}

func Delete_Blob(id int) error {
	db, err := get_DB()
	if err != nil {
		fmt.Println("DB connection error")
		return err
	}
	defer db.Close()
	_, err = db.Query(delete_blob, id)
	if err != nil {
		fmt.Println("SQL execution error:", err)
		return err
	}
	return nil
}
