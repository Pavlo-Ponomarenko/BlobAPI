package pg

import (
	"blob-service/internal/data"
	res "blob-service/resources"
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type blobsQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlInsert sq.InsertBuilder
	sqlUpdate sq.UpdateBuilder
	sqlDelete sq.DeleteBuilder
}

const blobsTable = "blobs"

func NewBlobsQ(db *pgdb.DB) data.BlobsQ {
	return &blobsQ{
		db:        db.Clone(),
		sql:       sq.Select("*").From(blobsTable),
		sqlInsert: sq.Insert(blobsTable),
		sqlUpdate: sq.Update(blobsTable),
		sqlDelete: sq.Delete(blobsTable),
	}
}

func (q *blobsQ) New() data.BlobsQ {
	return NewBlobsQ(q.db)
}

func (q *blobsQ) GetBlobById(id string) (*res.Blob, error) {
	var result blobEntity
	q.sql = q.sql.Where(sq.Eq{"id": id})
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return entityToBlob(&result), nil
}

func (q *blobsQ) GetBlobs(pageParams pgdb.OffsetPageParams) ([]res.Blob, error) {
	var result []blobEntity
	q.sql = pageParams.ApplyTo(q.sql, "id")
	err := q.db.Select(&result, q.sql)
	return entitiesToBlobs(result), err
}

func (q *blobsQ) SaveBlob(blob *res.Blob) (*res.Blob, error) {
	clauses := structs.Map(blobToEntity(blob))
	var result blobEntity
	q.sqlInsert = q.sqlInsert.SetMap(clauses).Suffix("returning id, blob")
	err := q.db.Get(&result, q.sqlInsert)
	return entityToBlob(&result), err
}

func (q *blobsQ) DeleteBlob(id string) error {
	q.sqlDelete = q.sqlDelete.Where(sq.Eq{"id": id})
	return q.db.Exec(q.sqlDelete)
}

func (q *blobsQ) IdIsPresent(id string) bool {
	blob, _ := q.GetBlobById(id)
	if blob == nil {
		return false
	}
	return true
}

func (q *blobsQ) UpdateBlob(id string, blob *res.Blob) (*res.Blob, error) {
	if !q.IdIsPresent(id) {
		return nil, errors.New("")
	}
	entity := new(blobEntity)
	entity.Id = id
	entity.Blob = blob.Attributes.Value
	clauses := structs.Map(entity)
	q.sqlUpdate = q.sqlUpdate.SetMap(clauses).Where(sq.Eq{"id": id}).Suffix("returning id, blob")
	var result blobEntity
	err := q.db.Get(&result, q.sqlUpdate)
	if err != nil {
		return nil, err
	}
	return entityToBlob(&result), nil
}
