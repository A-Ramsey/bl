package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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
		filePath := path + "/" + file.Name()
		osStat, _ := os.Stat(filePath)
		username, group := "", ""
		if stat, ok := osStat.Sys().(*syscall.Stat_t); ok {
			UID := strconv.FormatInt(int64(stat.Uid), 10)
			GID := strconv.FormatInt(int64(stat.Gid), 10)
			cmdUsername, _ := exec.Command("id", "-nu", UID).CombinedOutput()
			cmdGroup, _ := exec.Command("id", "-ng", GID).CombinedOutput()
			username, group = strings.TrimSpace(string(cmdUsername)), strings.TrimSpace(string(cmdGroup))
		}

		// Output File Info
		fmt.Print(buildIterationPrecursor(iterationDepth))
		fmt.Print(" ")
		fmt.Print(info.Mode())
		fmt.Print(" ")
		fmt.Print(username)
		fmt.Print(" ")
		fmt.Print(group)
		fmt.Print(" ")
		fmt.Print(info.ModTime().Format("Jan 02 15:04"))
		fmt.Print(" ")
		fmt.Print(colour)
		fmt.Print(file.Name())
		fmt.Print(Reset)
		fmt.Print("\n")
		if file.IsDir() && iterationDepth < maxDepth {
			interateFiles(filePath, iterationDepth+1, maxDepth)
		}
	}
}

func main() {
	dFlag := flag.Int("d", 3, "Max depth to recursively show files in directories")
	flag.Parse()
	interateFiles(".", 0, *dFlag)
}
