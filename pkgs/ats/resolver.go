package ats

import (
	"github.com/pulkitbhardwaj/omega/pkgs/ats/internal"
	"go.uber.org/fx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	fx.In

	Graph *internal.Client
}
