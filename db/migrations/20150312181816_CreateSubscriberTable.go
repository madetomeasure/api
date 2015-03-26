// +build ignore

package main

import (
	"database/sql"
)

// Up20150312181816 will create a simple subscribers table
func Up_20150312181816(txn *sql.Tx) {
	rawSql := `
  CREATE TABLE subscribers(
    id serial primary key,
    address varchar(255) not null,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
  )
  `
	txn.Exec(rawSql)
}

// Down20150312181816 is executed when this migration is rolled back
func Down_20150312181816(txn *sql.Tx) {
	txn.Exec(`DROP TABLE subscribers`)
}
