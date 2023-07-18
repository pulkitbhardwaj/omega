package ats

import (
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/pulkitbhardwaj/omega/pkgs/ats/internal"
)

func NewGraphQLHandler(r Resolver) http.HandlerFunc {
	srv := handler.New(
		internal.NewExecutableSchema(internal.Config{
			Resolvers:  &r,
			Directives: internal.DirectiveRoot{},
		}),
	)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	srv.Use(entgql.Transactioner{TxOpener: r.Graph})

	return srv.ServeHTTP
}

func NewGraphQLPlayground() http.HandlerFunc {
	return playground.Handler("GraphQL Playground", "/")
}
