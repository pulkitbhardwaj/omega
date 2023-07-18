package ats

import "go.uber.org/fx"

var Module = fx.Module("ats",
	fx.Provide(NewSQLConnection),
	fx.Invoke(NewHTTPServer),
)
