package unexported

//go:generate genny -in=generic_internal.go -out=mytype_internal.go gen "secret=*myType"
