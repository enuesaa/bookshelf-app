// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dbq

import (
	"database/sql"
)

type Tea struct {
	ID         int64
	Teapod     string
	Collection string
	Teaid        string
	Value      interface{}
	Created    sql.NullTime
	Updated    sql.NullTime
}
