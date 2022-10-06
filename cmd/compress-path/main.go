package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dnnrly/abbreviate/data"
	"github.com/dnnrly/abbreviate/domain"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("compress-path " + Version)
		os.Exit(0)
	}

	var (
		matcher     *domain.Matcher
		abbreviator domain.Abbreviator
	)

	currentPath, err := filepath.Abs(".")

	if err != nil {
		fmt.Printf("%s", ".")
		os.Exit(0)
	}

	homeDir, homeDirErr := os.UserHomeDir()

	if homeDirErr == nil && strings.Contains(currentPath, homeDir) {
		currentPath = strings.Replace(currentPath, homeDir, "~", 1)
	}

	newPath := currentPath

	if len(currentPath) > 15 {
		matcher = data.Abbreviations["en-us"]["common"]
		abbreviator = matcher

		parts := strings.Split(currentPath, string(os.PathSeparator))

		for i, v := range parts {
			if i < len(parts)-1 && len(v) > 5 {
				parts[i] = abbreviator.Abbreviate(parts[i])
			}
		}

		newPath = strings.Join(parts, string(os.PathSeparator))
	}

	fmt.Printf("%s", newPath)
}
