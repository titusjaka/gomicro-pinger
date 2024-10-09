package main

import (
	"github.com/titusjaka/gomicro-pinger/commands"

	"github.com/alecthomas/kong"
)

type App struct {
	Pinger commands.PingerCmd `kong:"cmd,name=pinger"`
	Ponger commands.PongerCmd `kong:"cmd,name=ponger"`
}

func main() {
	var app App
	kCtx := kong.Parse(
		&app,
	)
	kCtx.FatalIfErrorf(kCtx.Run())
}
