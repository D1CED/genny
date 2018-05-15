package queue

//go:generate genny -in=generic_queue.go -out=int_queue.go gen "Something=int"
//go:generate genny -in=generic_queue.go -out=float32_queue.go gen "Something=float32"
//go:generate genny -in=generic_queue.go -out=changed/int_queue.go -pkg=changed gen "Something=int"
