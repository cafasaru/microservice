package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

var txKey = struct{}{}

// execTx executes a funcion within a database useful when we need to role back if multiple dependent queries error out.
func (store *Store) execTx(c context.Context, fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(c, nil)
	if err != nil {
		log.Println("error from BeginTx")
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()

}
