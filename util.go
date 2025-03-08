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

func Fib(u uint) uint {
	if u <= 1 {
		return 1
	}
	return Fib(u-2) + Fib(u-1)
}

func abc(a, b int) int {
	return a + b
}
