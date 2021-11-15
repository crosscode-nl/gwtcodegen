package main

import (
	"github.com/alecthomas/kong"
	"gwtcodegen/cli"
)

func main() {
	cliValues := &cli.CLI
	ctx := kong.Parse(cliValues)
	err := ctx.Run(&cli.Context{
		IndentChar: cliValues.IndentChar,
		IndentNum:  cliValues.IndentNum,
	})
	ctx.FatalIfErrorf(err)
}
