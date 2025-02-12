package main

import (
	"strings"
	"testing"
)

func BenchmarkEscapeSequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// using octal escape seq
		_ = "\040\040\137\040\040\040\040\040\040\040\040\040\040"
	}
}

func BenchmarkPlainString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "  _          "
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.WriteString("  _          ")
		_ = sb.String()
	}
}
