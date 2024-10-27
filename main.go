package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] == "help" {
		printHelp()
		os.Exit(0)
	}

	_, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println("File '" + os.Args[1] + "' does not exist or you do not have the permissions")
		printHelp()
		os.Exit(1)
	}

	filePathAbs, _ := filepath.Abs(os.Args[1])
	fileName := filepath.Base(os.Args[1])
	out := &outputBuilder{
		pathPrefix: filepath.Dir(filePathAbs),
		regexp: regexp.MustCompile(`([ \t]*)?{{(.*?)}}`),
	}
	result := out.getFileContent(fileName, "")

	outFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating file '" + os.Args[2] + "'")
		os.Exit(1)
	}
	defer outFile.Close()

	_, err = outFile.WriteString(result)
	if err != nil {
		fmt.Println("Error writing to file '" + os.Args[2] + "'")
		os.Exit(1)
	}

	fmt.Println("Merged all files into '" + os.Args[2] + "'")
}
