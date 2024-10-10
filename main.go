package main

import (
	"log/slog"
	"os"

	"github.com/titusjaka/gomicro-pinger/commands"

	"github.com/alecthomas/kong"
)

type App struct {
	Pinger commands.PingerCmd `kong:"cmd,name=pinger"`
	Ponger commands.PongerCmd `kong:"cmd,name=ponger"`
}

func main() {
	slog.SetDefault(
		slog.Default().With(
			slog.String("HOSTNAME", os.Getenv("HOSTNAME")),
		),
	)

	var app App
	kCtx := kong.Parse(
		&app,
	)
	kCtx.FatalIfErrorf(kCtx.Run())
}
