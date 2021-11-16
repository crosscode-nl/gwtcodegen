package model

import "strings"

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
