package main

import (
	"fmt"
)

const helpMessage = `
Usage: pagemerge [OPTION] [HEADFILE] [OUTPUTFILE]

Options:
  help        Show this help message and exit

Examples:
  pagemerge help
  pagemerge index.html ../dist/index.html
`

func printHelp() {
	fmt.Println(helpMessage)
}
