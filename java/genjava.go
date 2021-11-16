package java

import (
	"gwtcodegen/gwtparser"
	"gwtcodegen/gwtstrings"
	"gwtcodegen/model"
	"os"
	"path/filepath"
	"strings"
)

type writer interface {
	WriteString(s string) (int, error)
}

func addHeader(className string, sb writer) {
	_,_ = sb.WriteString("package change.me;\n\n")
	_,_ = sb.WriteString("import static org.junit.jupiter.api.Assertions.*;\n\n")
	_,_ = sb.WriteString("import org.junit.jupiter.api.BeforeEach;\n")
	_,_ = sb.WriteString("import org.junit.jupiter.api.DisplayName;\n")
	_,_ = sb.WriteString("import org.junit.jupiter.api.Nested;\n")
	_,_ = sb.WriteString("import org.junit.jupiter.api.Test;\n\n")
	_,_ = sb.WriteString("@DisplayName(\"Feature ")
	_,_ = sb.WriteString(className)
	_,_ = sb.WriteString("\")\n")
	_,_ = sb.WriteString("class ")
	_,_ = sb.WriteString(className)
	_,_ = sb.WriteString(" {\n")
}

func addGiven(context model.Context, sb writer, given model.Given) {
	_,_ = sb.WriteString("\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@Nested\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@DisplayName(\"")
	_,_ = sb.WriteString(strings.TrimSpace(strings.ReplaceAll(given.Text, "\"","\\\"")))
	_,_ = sb.WriteString("\")\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("class ")
	_,_ = sb.WriteString(gwtstrings.CamelCase(gwtstrings.ReplaceNonAlphaNumWithSpace(given.Text)))
	_,_ = sb.WriteString(" {\n")
	addWhenArray(context, sb, given.When)
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("}\n")
}

func addWhenArray(context model.Context, sb writer, whenArray []model.When) {
	for _, when := range whenArray {
		addWhen(context.IncreaseNesting(), sb,when)
	}
}

func addWhen(context model.Context, sb writer, when model.When) {
	_,_ = sb.WriteString("\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@Nested\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@DisplayName(\"")
	_,_ = sb.WriteString(strings.TrimSpace(strings.ReplaceAll(when.Text, "\"","\\\"")))
	_,_ = sb.WriteString("\")\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("class ")
	_,_ = sb.WriteString(gwtstrings.CamelCase(gwtstrings.ReplaceNonAlphaNumWithSpace(when.Text)))
	_,_ = sb.WriteString(" {\n")
	addWhenFunction(context.IncreaseNesting(), sb)
	addThenArray(context.IncreaseNesting(), sb, when.Then)
	_, _ = sb.WriteString(context.RenderPrefix())
	_, _ = sb.WriteString("}\n")
}

func addWhenFunction(context model.Context, sb writer) {
	_, _ = sb.WriteString(context.RenderPrefix())
	_, _ = sb.WriteString("@BeforeEach\n")
	_, _ = sb.WriteString(context.RenderPrefix())
	_, _ = sb.WriteString("void when() {\n")
	_, _ = sb.WriteString(context.RenderPrefix())
	_, _ = sb.WriteString("}\n")
}

func addThenArray(context model.Context, sb writer, thenArray []model.Then) {
	for _, then := range thenArray {
		addThen(context,sb,then)
	}
}

func addThen(context model.Context, sb writer, then model.Then) {
	_,_ = sb.WriteString("\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@Test\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("@DisplayName(\"")
	_,_ = sb.WriteString(strings.TrimSpace(strings.ReplaceAll(then.Text, "\"","\\\"")))
	_,_ = sb.WriteString("\")\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("void test")
	_,_ = sb.WriteString(gwtstrings.CamelCase(gwtstrings.ReplaceNonAlphaNumWithSpace(then.Text)))
	_,_ = sb.WriteString("() {\n")
	_,_ = sb.WriteString(context.IncreaseNesting().RenderPrefix())
	_,_ = sb.WriteString("fail(\"Test not implemented.\");\n")
	_,_ = sb.WriteString(context.RenderPrefix())
	_,_ = sb.WriteString("}\n")
}

func addGivenArray(context model.Context, className string, sb writer, givenArray []model.Given) {

	addHeader(className, sb)
	for _, given := range givenArray {
		addGiven(context.IncreaseNesting(), sb,given)
	}
	addFooter(sb)
}

func addFooter(sb writer) {
	_,_ = sb.WriteString("}\n")
}


func classNameFromFilePath(filePath string) string {
	return strings.ReplaceAll(filepath.Base(filePath),filepath.Ext(filePath),"")
}

func Generate(context model.Context, source string, destination string) (err error) {
	var data []byte
	if data, err = os.ReadFile(source); err!=nil {
		return
	}
	givenArray := gwtparser.Parse(string(data))
	var sb = &strings.Builder{}
	addGivenArray(context,classNameFromFilePath(destination),sb,givenArray)
	if err=os.WriteFile(destination,[]byte(sb.String()),os.ModeAppend); err!=nil {
		return err
	}
	return
}
