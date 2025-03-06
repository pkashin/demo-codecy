package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Unix style absolute path with spaces",
			path:     `/usr/local/my project/tmp/main`,
			expected: `"/usr/local/my project/tmp/main"`,
		},
		{
			name:     "Unix style relative path with spaces",
			path:     "./my project/tmp/main",
			expected: "./my project/tmp/main",
		},
		{
			name:     "Unix style absolute path without spaces",
			path:     `/usr/local/project/tmp/main`,
			expected: `/usr/local/project/tmp/main`,
		},
		{
			name:     "Empty path",
			path:     "",
			expected: "",
		},
		{
			name:     "Simple path",
			path:     "main.go",
			expected: "main.go",
		},
		{
			name:     "Command with argument",
			path:     "sh main.sh",
			expected: "sh main.sh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatPath(tt.path)
			assert.Equal(t, tt.expected, result, "formatPath(%q)", tt.path)
		})
	}
}
