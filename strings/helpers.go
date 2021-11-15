package strings

import (
	"regexp"
	"strings"
)

var notAlphaNumRegexp, _ = regexp.Compile("[^a-zA-Z0-9]+")

func ReplaceNonAlphaNumWithSpace(line string) string {
	return notAlphaNumRegexp.ReplaceAllString(line," ")
}

func CamelCase(line string) (result string) {
	var sb strings.Builder
	sb.Grow(len(line)+1)
	for _,part := range strings.Split(line, " ") {
		part = strings.TrimSpace(part)
		if len(part)>0 {
			sb.WriteString(strings.ToUpper(part[:1]))
			if len(part)>1 {
				sb.WriteString(part[1:])
			}
		}
	}
	result = sb.String()
	return
}

func Capitalize(line string) (result string) {
	return strings.ToUpper(line[:1]) + line[1:]
}

func TrimFromKeyword(line string, keys []string) string {
	lastIndex := -1
	for _,key := range keys {
		index := strings.Index(strings.ToLower(line),key)
		if lastIndex==-1 {
			lastIndex = index
			continue
		}
		if index!=-1 && index < lastIndex {
			lastIndex = index
		}
	}
	if lastIndex==-1 {
		return line
	}
	return Capitalize(line[lastIndex:])
}
