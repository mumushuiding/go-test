package test

import (
	"database/sql"
	"testing"
)

func TestTransactiontest(t *testing.T) {
	db, err := sql.Open("mysql", "root@/test")
	if err != nil {
		panic(err)
	}
	db.Begin()
}
