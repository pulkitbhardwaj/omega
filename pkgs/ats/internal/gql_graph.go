package internal

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/google/uuid"
)

type Graph interface {
	entgql.TxOpener

	Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (Noder, error)

	Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error)
}
