package multipletypesets

// got int_string, int_bool, float64_string, float64_bool want int_string, float64_bool
//go:generate genny -in=generic_simplemap.go -out=many_simplemaps.go gen "KeyType=int,float64 ValueType=string,bool"

//-go:generate genny -in=generic_simplemap.go -out=many_simplemaps.go gen "KeyType=int ValueType=string"
//-go:generate genny -in=generic_simplemap.go -out=many_simplemaps.go gen "KeyType=float64 ValueType=bool"
//-go:generate genny -in=generic_simplemap.go -out=many_simplemaps.go gen "KeyType=int ValueType=string" gen "KeyType=float64 ValueType=bool"
//-go:generate genny -in=generic_simplemap.go -out=many_simplemaps.go gen "KeyType=int ValueType=string KeyType=float64 ValueType=bool"
