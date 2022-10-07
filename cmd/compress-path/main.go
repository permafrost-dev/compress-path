package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dnnrly/abbreviate/data"
	"github.com/dnnrly/abbreviate/domain"
)

func main() {
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

	if homeDirErr == nil && strings.HasPrefix(currentPath, homeDir) {
		currentPath = strings.Replace(currentPath, homeDir, "~", 1)
	}

	if len(currentPath) <= 15 {
		fmt.Printf("%s", currentPath)
		os.Exit(0)
	}

	matcher = data.Abbreviations["en-us"]["common"]
	abbreviator = matcher

	parts := strings.Split(currentPath, string(os.PathSeparator))
	newParts := make([]string, len(parts))
	var wg sync.WaitGroup
	wg.Add(len(parts))

	for i, v := range parts {
		go func(i int, v string) {
			defer wg.Done()
			if i < len(newParts)-1 && len(v) > 5 {
				newParts[i] = abbreviator.Abbreviate(v)
			} else {
				newParts[i] = v
			}

			if i == len(newParts)-1 {
				newParts[i] = "\033[38;5;159m\033[1m" + newParts[i] + "\033[22m\033[38;5;0m"
			} else {
				newParts[i] = "\033[38;5;123m" + newParts[i] + "\033[38;5;0m"
			}
		}(i, v)
	}

	wg.Wait()
	sep := string(os.PathSeparator)
	//sep := "\033[38;5;255m\ue216\033[38;5;0m"
	fmt.Printf("%s", strings.Join(newParts, sep))
}
