package main

import (
	"github.com/alecthomas/kong"
	"gwtcodegen/cli"
)

func main() {
	ctx := kong.Parse(&cli.CLI)
	err := ctx.Run(&cli.Context{})
	ctx.FatalIfErrorf(err)
}
