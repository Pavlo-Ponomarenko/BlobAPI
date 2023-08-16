package pg

import (
	"blob-service/internal/data"
	"database/sql"
	"fmt"
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

func (q *blobsQ) GetBlobById(id string) (*data.BlobEntity, error) {
	var result data.BlobEntity
	q.sql = q.sql.Where(sq.Eq{"id": id})
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &result, nil
}

func (q *blobsQ) GetBlobs(pageParams pgdb.OffsetPageParams) ([]data.BlobEntity, error) {
	var result []data.BlobEntity
	q.sql = pageParams.ApplyTo(q.sql, "id")
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *blobsQ) SaveBlob(blob *data.BlobEntity) (*data.BlobEntity, error) {
	clauses := structs.Map(blob)
	var result data.BlobEntity
	q.sqlInsert = q.sqlInsert.SetMap(clauses).Suffix("returning id, blob")
	err := q.db.Get(&result, q.sqlInsert)
	return &result, err
}

func (q *blobsQ) DeleteBlob(id string) error {
	q.sqlDelete = q.sqlDelete.Where(sq.Eq{"id": id})
	return q.db.Exec(q.sqlDelete)
}

func (q *blobsQ) IdIsPresent(id string) bool {
	subQuery := fmt.Sprintf("exists (select * from %s where id = '%s') as result", blobsTable, id)
	query := sq.Select(subQuery)
	var idExists bool
	err := q.db.Get(&idExists, query)
	if err != nil || !idExists {
		fmt.Println(err)
		return false
	}
	return true
}

func (q *blobsQ) UpdateBlob(id string, blob *data.BlobEntity) (*data.BlobEntity, error) {
	blob.Id = id
	clauses := structs.Map(blob)
	q.sqlUpdate = q.sqlUpdate.SetMap(clauses).Where(sq.Eq{"id": id}).Suffix("returning id, blob")
	var result data.BlobEntity
	err := q.db.Get(&result, q.sqlUpdate)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
