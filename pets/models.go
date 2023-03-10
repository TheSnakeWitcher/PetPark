// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package pets

import (
	"database/sql"

	"github.com/tabbed/pqtype"
)

type Pet struct {
	ID       int32                 `db:"id" json:"id"`
	Name     sql.NullString        `db:"name" json:"name"`
	Picked   sql.NullTime          `db:"picked" json:"picked"`
	Location sql.NullString        `db:"location" json:"location"`
	Data     pqtype.NullRawMessage `db:"data" json:"data"`
}
