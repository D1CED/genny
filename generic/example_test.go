package generic_test

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type (
	T generic.Type
	N generic.Number

	StructT struct {
		t *T
	}
)

func MyFunctionN(n N) N {
	return n + 5
}

func Example() {
	var num N = 3
	for i := N(0); i < num; i++ {
		fmt.Println("Hello genny!")
	}
	// Output:
	// Hello genny!
	// Hello genny!
	// Hello genny!
}
