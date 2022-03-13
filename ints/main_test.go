package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() { //if your code changes so that the example
	sum := Add(1, 5) //is no longer valid, your build will fail
	fmt.Println(sum)
	//Output: 6
}
