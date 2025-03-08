package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
}

func BenchmarkFib20WithAuxMetric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}
	b.ReportMetric(4.0, "auxMetricUnits")
}

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

func Test_abc(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "check",
			args: args{2, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, abc(tt.args.a, tt.args.b), "abc(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
