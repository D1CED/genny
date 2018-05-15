// Writing test code
// Once you have defined a generic type with some code worth testing:
package main_test

import (
	"log"
	"testing"

	"github.com/cheekybits/genny/generic"
	"github.com/stretchr/testify/assert"
)

type MyType generic.Type

func EnsureMyTypeSlice(objectOrSlice interface{}) []MyType {
	log.Printf("%T", objectOrSlice)
	switch obj := objectOrSlice.(type) {
	case []MyType:
		log.Println("returning it untouched")
		return obj
	case MyType:
		log.Println("wrapping in slice")
		return []MyType{obj}
	default:
		panic("ensure slice needs MyType or []MyType")
	}
}

// You can treat it like any normal Go type in your test code:
func TestEnsureMyTypeSlice(t *testing.T) {

	myType := new(MyType)
	slice := EnsureMyTypeSlice(myType)
	if assert.NotNil(t, slice) {
		assert.Equal(t, slice[0], myType)
	}

	slice = EnsureMyTypeSlice(slice)
	log.Printf("%#v", slice[0])
	if assert.NotNil(t, slice) {
		assert.Equal(t, slice[0], myType)
	}

}
