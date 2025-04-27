package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
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
	files, err := os.ReadDir(path)
	if strings.Contains(path, ".git") || strings.Contains(path, "node_modules") {
		return
	}
	if err != nil {
		fmt.Println(buildIterationPrecursor(iterationDepth), err)
	}

	for _, file := range files {
		colour := Reset
		if file.IsDir() {
			colour = Blue
		}
		info, err := file.Info()
		if err != nil {
			fmt.Println(buildIterationPrecursor(iterationDepth), err)
		}
		fmt.Println(buildIterationPrecursor(iterationDepth), info.Mode(), colour, file.Name(), Reset)
		if file.IsDir() && iterationDepth < maxDepth {
			interateFiles(path+"/"+file.Name(), iterationDepth+1, maxDepth)
		}
	}
}

func main() {
	dFlag := flag.Int("d", 3, "Max depth to recursively show files in directories")
	flag.Parse()
	interateFiles(".", 0, *dFlag)
}
