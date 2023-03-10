// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package pets

import (
	"context"
	"database/sql"
)

type Querier interface {
	AddPet(ctx context.Context, arg AddPetParams) (sql.Result, error)
	DelPet(ctx context.Context, id int32) error
	GetPet(ctx context.Context, id int32) (Pet, error)
	ListPets(ctx context.Context) ([]Pet, error)
	UpdateState(ctx context.Context, id int32) error
}

var _ Querier = (*Queries)(nil)
