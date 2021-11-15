package java

import (
	"gwtcodegen/cli"
	"gwtcodegen/gwtparser"
	"gwtcodegen/gwtstrings"
	"gwtcodegen/model"
	"os"
	"strings"
)

type writer interface {
	WriteString(s string) (int, error)
}

func addHeader(sb writer) {
}

func addGiven(context cli.Context, sb writer, given model.Given) {
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@Nested(\"")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@DisplayName(\"")
	_,_ = sb.WriteString(strings.ReplaceAll(given.Text, "\"","\\\""))
	_,_ = sb.WriteString("\")\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("class")
	_,_ = sb.WriteString(gwtstrings.CamelCase(gwtstrings.ReplaceNonAlphaNumWithSpace(given.Text)))
	_,_ = sb.WriteString(" {\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("}\n")
}

func addGivenArray(context cli.Context, sb writer, givenArray []model.Given) {

	addHeader(sb)
	for _, given := range givenArray {
		addGiven(context.IncreaseNesting(), sb,given)
	}
	addFooter(sb)
}

func addFooter(sb writer) {
}


func Generate(context cli.Context, source string, destination string) (err error) {
	var data []byte
	if data, err = os.ReadFile(source); err!=nil {
		return
	}
	givenArray := gwtparser.Parse(string(data))
	var sb = &strings.Builder{}
	addGivenArray(
		context.IncreaseNesting(),
		sb,
		givenArray)
	return
}
