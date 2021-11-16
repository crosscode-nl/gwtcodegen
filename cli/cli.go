package cli

import (
	"gwtcodegen/java"
	"gwtcodegen/model"
)


type GenJavaCmd struct {
	Source string `arg:"" help:"Input file."`
	Destination string `arg:"" help:"Output file."`
}

func (r *GenJavaCmd) Run(ctx *model.Context) error {
	return java.Generate(*ctx, r.Source,r.Destination)
}

var CLI struct {
	GenJava GenJavaCmd `cmd:"" help:"Generate java file"`
	IndentChar string `default:"tab" enum:"tab,space" help:"Indentation character (tab or space)"`
	IndentNum int `default:"1" optional:"" help:"Number of indentations"`
}