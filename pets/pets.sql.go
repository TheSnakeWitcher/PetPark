// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: pets.sql

package pets

import (
	"context"
	"database/sql"
)

const addPet = `-- name: AddPet :execresult
insert into pets ( id , name, picked ,location,data) 
values( $1,$2,now(),$3,'{}')
`

type AddPetParams struct {
	ID       int32          `db:"id" json:"id"`
	Name     sql.NullString `db:"name" json:"name"`
	Location sql.NullString `db:"location" json:"location"`
}

func (q *Queries) AddPet(ctx context.Context, arg AddPetParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addPet, arg.ID, arg.Name, arg.Location)
}

const delPet = `-- name: DelPet :exec
delete from pets where id = $1
`

func (q *Queries) DelPet(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, delPet, id)
	return err
}

const getPet = `-- name: GetPet :one
select id, name, picked, location, data from pets where id = $1
`

func (q *Queries) GetPet(ctx context.Context, id int32) (Pet, error) {
	row := q.db.QueryRowContext(ctx, getPet, id)
	var i Pet
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Picked,
		&i.Location,
		&i.Data,
	)
	return i, err
}

const listPets = `-- name: ListPets :many
select id, name, picked, location, data from pets order by name
`

func (q *Queries) ListPets(ctx context.Context) ([]Pet, error) {
	rows, err := q.db.QueryContext(ctx, listPets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Pet{}
	for rows.Next() {
		var i Pet
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Picked,
			&i.Location,
			&i.Data,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateState = `-- name: UpdateState :exec
update pets 
set status = 'adopted'
where id = ?
`

func (q *Queries) UpdateState(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, updateState, id)
	return err
}
