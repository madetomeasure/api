
package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20150312181816(txn *sql.Tx) {
  raw_sql := `
  CREATE TABLE subscribers(
    id serial primary key,
    address varchar(255) not null,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
  )
  `
  txn.Exec(raw_sql)
}

// Down is executed when this migration is rolled back
func Down_20150312181816(txn *sql.Tx) {
  txn.Exec(`DROP TABLE subscribers`)
}
