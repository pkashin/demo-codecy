package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func formatPath(path string) string {
	if !filepath.IsAbs(path) || !strings.Contains(path, " ") {
		return path
	}

	return fmt.Sprintf(`"%s"`, path)
}
