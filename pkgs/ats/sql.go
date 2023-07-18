package ats

import (
	"context"

	"log"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/omega/pkgs/ats/internal"
)

func NewSQLConnection(lc fx.Lifecycle) *internal.Client {
	// Open the database connection.
	db, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening database", err)
	}
	// Decorates the sql.Driver with entcache.Driver.
	drv := entcache.NewDriver(db)
	// Create an ent.Client.
	client := internal.NewClient(internal.Driver(drv))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(entcache.Skip(ctx))
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
	return client
}
