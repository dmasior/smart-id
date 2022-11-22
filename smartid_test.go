package smartid

import (
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustNew()
	}
}

func BenchmarkNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustNewString()
	}
}
