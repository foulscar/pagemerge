package main

import (
	"os"
	"fmt"
	"strings"
	"regexp"
)

type outputBuilder struct {
	pathPrefix string
	regexp *regexp.Regexp
}

func (o *outputBuilder) getFileContent(fileName string, tabPrefix string) string {
	data, err := os.ReadFile(o.pathPrefix + "/" + fileName)
	if err != nil {
		fmt.Println(o.pathPrefix)
		fmt.Println("Error reading tag with file '" + fileName + "'")
		return ""
	}
	result := string(data)

	matches := o.regexp.FindAllStringSubmatch(result, -1)
	for _, match := range matches {
		if len(match) < 3 {
			continue
		}

		tagTabPrefix := match[1]
		tagFileName := match[2]
		tagContent := o.getFileContent(tagFileName, tagTabPrefix)

		result = strings.ReplaceAll(result, match[0], tagContent)
	}

	return addPrefixBeforeLines(result, tabPrefix)
}

func addPrefixBeforeLines(input string, prefix string) string {
	lines := strings.Split(input, "\n")
	if len(lines) < 2 || prefix == "" {
		return prefix + input
	}

	var result strings.Builder

	for i, line := range lines {
		result.WriteString(prefix + line)
		if i < len(lines) - 2 {
			result.WriteString("\n")
		}
	}

	return result.String()
}
