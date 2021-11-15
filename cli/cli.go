package cli

import (
	"gwtcodegen/java"
	"strings"
)

type Context struct {
	IndentChar string
	IndentNum int
	Nested int
}

func (c Context) IncreaseNesting() Context {
	c.Nested++
	return c
}

func (c Context) GetIndentChar() string {
	switch c.IndentChar {
		case "space": return " "
		case "tab": return "\t"
		default: return "\t"
	}
}

func (c Context) RenderPrefix() string {
	return strings.Repeat(c.GetIndentChar(),c.IndentNum*c.Nested)
}

type GenJavaCmd struct {
	Source string `arg:"" help:"Input file."`
	Destination string `arg:"" help:"Output file."`
}

func (r *GenJavaCmd) Run(ctx *Context) error {
	return java.Generate(*ctx, r.Source,r.Destination)
}

var CLI struct {
	GenJava GenJavaCmd `cmd:"" help:"Generate java file"`
	IndentChar string `default:"tab" enum:"tab,space" help:"Indentation character (tab or space)"`
	IndentNum int `default:"1" optional:"" help:"Number of indentations"`
}