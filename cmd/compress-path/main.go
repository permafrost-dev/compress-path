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
	var (
		matcher     *domain.Matcher
		abbreviator domain.Abbreviator
	)

	path, err := filepath.Abs(".")

	if err != nil {
		fmt.Printf("%s", ".")
		os.Exit(0)
	}

	homeDir, homeDirErr := os.UserHomeDir()

	if homeDirErr == nil {
		path = strings.Replace(path, homeDir, "~", 1)
	}

	newPath := path

	if len(path) > 15 {
		matcher = data.Abbreviations["en-us"]["common"]
		abbreviator = matcher

		parts := strings.Split(path, string(os.PathSeparator))

		for i, v := range parts {
			if i < len(parts)-1 && len(v) > 5 {
				parts[i] = abbreviator.Abbreviate(parts[i])
			}
		}

		newPath = strings.Join(parts, string(os.PathSeparator))
	}

	fmt.Printf("%s", newPath)
}
