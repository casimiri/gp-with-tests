package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("expected sum", func(t *testing.T) {
		sum := Add(2, 2)

		expected := 4

		if sum != expected {
			t.Errorf("Expected %q but got %q", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
