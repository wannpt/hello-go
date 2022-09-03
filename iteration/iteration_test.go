package iteration

import (
	"fmt"
	"testing"
)

//caller can specify letter and how many times the character is repeated
func TestRepeat(t *testing.T) {
	get := Repeat("a", 9);
	expected := "aaaaaaaaa";

	if get != expected {
		t.Errorf("expected %q but got %q", expected, get)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 9)
	fmt.Println(result)
	// Output: aaaaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 9)
	}
}