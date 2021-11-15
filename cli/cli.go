package cli

import "gwtcodegen/java"

type Context struct {
}

type GenJavaCmd struct {
	Source string `arg:"" help:"Input file."`
	Destination string `arg:"" help:"Output file."`
}

func (r *GenJavaCmd) Run(ctx *Context) error {
	return java.Generate(r.Source,r.Destination)
}

var CLI struct {
	GenJava GenJavaCmd `cmd:"" help:"Generate java file"`
}