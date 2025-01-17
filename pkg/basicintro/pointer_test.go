package basicintro_test

import (
	"testing"

	"io.huangsam/trial/pkg/basicintro"
)

func TestPointerSetOne(t *testing.T) {
	var n int = 0
	basicintro.PointerSetOne(&n)
	if n != 1 {
		t.Errorf("Expected n to be 1, but got %d", n)
	}
}

func BenchmarkPointerSetOne(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		basicintro.PointerSetOne(&n)
	}
}
