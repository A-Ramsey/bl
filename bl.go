package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func buildIterationPrecursor(depth int) string {
	iterationPrecursor := ""
	if depth > 0 {
		iterationPrecursor += "|"
	}
	for i := 0; i < depth; i++ {
		iterationPrecursor += "-"
	}
	return iterationPrecursor
}

func interateFiles(path string, iterationDepth int, maxDepth int) {
	files, err := ioutil.ReadDir(path)
	if strings.Contains(path, ".git") || strings.Contains(path, "node_modules") {
		return
	}
	if err != nil {
		fmt.Println(buildIterationPrecursor, err)
	}

	for _, file := range files {
		colour := Reset
		if file.IsDir() {
			colour = Green
		}
		fmt.Println(buildIterationPrecursor(iterationDepth), file.Mode(), colour, file.Name(), Reset)
		if file.IsDir() && iterationDepth < maxDepth {
			interateFiles(path+"/"+file.Name(), iterationDepth+1, maxDepth)
		}
	}
}

func main() {
	interateFiles(".", 0, 3)
}
