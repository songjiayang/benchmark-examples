package benchmark

import (
	"log"
	"testing"
)

func BenchmarkStringEmptyCheckWithDefault(b *testing.B) {
	var foo string
	for i := 0; i < b.N; i++ {
		if foo != "" {
			log.Panicf("empty string should be default")
		}
	}
}

func BenchmarkStringEmptyCheckWithLen(b *testing.B) {
	var foo string
	for i := 0; i < b.N; i++ {
		if len(foo) != 0 {
			log.Panicf("empty string length should be 0")
		}
	}
}
