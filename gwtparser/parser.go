package gwtparser

import (
	"gwtcodegen/model"
	str "gwtcodegen/gwtstrings"
	"strings"
)

func Parse(input string) (results []model.Given) {
	var normalizedLines []string
	for _, line := range strings.Split(input,"\n")  {
		normalizedLines = append(normalizedLines, str.TrimFromKeyword(line,[]string{"given","when","then"}))
	}
	for _, line := range normalizedLines {
		if line[:5]=="Given" {
			results = append(results,
				model.Given{
					Text: line,
					When: nil,
				})
		}
		if line[:4]=="When" && results != nil {
			when := results[len(results)-1].When
			results[len(results)-1].When = append(when,
				model.When{
					Text: line,
					Then: nil,
				})
		}
		if line[:4]=="Then" && results != nil && results[len(results)-1].When != nil {
			then := results[len(results)-1].When[len(results[len(results)-1].When)-1].Then
			results[len(results)-1].When[len(results[len(results)-1].When)-1].Then = append(then,
				model.Then{
					Text: line,
				})
		}
	}
	return
}