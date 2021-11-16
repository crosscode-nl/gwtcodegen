package main

import (
	"github.com/alecthomas/kong"
	"gwtcodegen/cli"
	"gwtcodegen/model"
)

func main() {
	cliValues := &cli.CLI
	ctx := kong.Parse(cliValues)
	err := ctx.Run(&model.Context{
		IndentChar: cliValues.IndentChar,
		IndentNum:  cliValues.IndentNum,
	})
	ctx.FatalIfErrorf(err)
}
