// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package gen

import (
	"context"
)

type Querier interface {
	PlayerByID(ctx context.Context, id int64) (Player, error)
}

var _ Querier = (*Queries)(nil)