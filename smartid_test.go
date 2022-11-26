package smartid

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	id, err := New()

	if err != nil {
		t.Errorf("error creating new smartid")
	}
	if got, want := len(id), 27; got != want {
		t.Errorf("id len should be 27")
	}
	if got, want := len(id.String()), 27; got != want {
		t.Errorf("id string len should be 27, got %d", len(id.String()))
	}
}

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

func Example() {
	id := MustNew()

	fmt.Println(id.String())
}
