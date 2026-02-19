package main

import (
	"strings"
	"testing"
)

// Data to join
var segments = []string{"hello", "world", "from", "go", "benchmarking"}

// 1. Using + Operator (Naive)
func BenchmarkPlusOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for _, v := range segments {
			s += v
		}
		_ = s
	}
}

// 2. Using strings.Join
func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(segments, "")
	}
}

// 3. Using strings.Builder
func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for _, v := range segments {
			sb.WriteString(v)
		}
		_ = sb.String()
	}
}
