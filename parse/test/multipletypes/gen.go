package multipletypes

//go:generate genny -in=generic_simplemap.go -out=interface_int_simplemap.go gen "KeyType=interface{} ValueType=int"
//go:generate genny -in=generic_simplemap.go -out=string_int_simplemap.go gen "KeyType=string ValueType=int"
//go:generate genny -in=generic_simplemap.go -out=custom_types_simplemap.go gen "KeyType=*MyType1 ValueType=*MyOtherType"
