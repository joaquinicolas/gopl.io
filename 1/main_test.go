package main

import "testing"

func BenchmarkPrintArgs(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		printArgs([]string {"a", "b", "c", "d", "e", "j"})
	}
}
func BenchmarkPrintArgsRange(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		printArgsRange([]string {"a", "b", "c", "d", "e", "j"})
	}
}