//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	gqlCfg := "./config.yml"

	ex, err := entgql.NewExtension(
		entgql.WithRelaySpec(true),
		entgql.WithNodeDescriptor(true),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaGenerator(),
		entgql.WithConfigPath(gqlCfg),
		entgql.WithSchemaPath("./entity.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.Annotations(entgql.RelayConnection()),
		// entc.TemplateDir("./internal/entity/template"),
	}

	if err := entc.Generate("./schema", &gen.Config{
		Templates: entgql.AllTemplates,
		Features:  []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	cfg, err := config.LoadConfig(gqlCfg)
	if err != nil {
		log.Fatalf("loading gqlgen config: %v", err)
	}

	if err = api.Generate(cfg); err != nil {
		log.Fatalf("running gqlgen codegen: %v", err)
	}
}
