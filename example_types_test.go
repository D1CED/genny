package main_test

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int ValueType=string,int"

func Example_types() {
	type (
		KeyType   generic.Type
		ValueType generic.Type
	)

	// KeyTypeValueTypeMap comments get processed too. Awesome!
	type KeyTypeValueTypeMap map[KeyType]ValueType

	type KeyTypeValueTypeFunc func() KeyTypeValueTypeMap
}

/*
To see a real example of how to use 'genny' with 'go generate', look in the
example/go-generate directory
(https://github.com/cheekybits/genny/tree/master/examples/go-generate).

To get a _something_ for every built-in Go type plus one of your own types,
you could run:

	cat source.go | genny gen "Something=BUILTINS,*MyType"

More examples

Check out the test code files
(https://github.com/cheekybits/genny/tree/master/parse/test)
for more real examples.
*/
