package generic

// Type is the placeholder type that indicates a generic value.
// When genny is executed, variables of this type will be replaced with
// references to the specific types.
type Type interface{}

// Number is the placehoder type that indiccates a generic numerical value.
// When genny is executed, variables of this type will be replaced with
// references to the specific types.
type Number float64
