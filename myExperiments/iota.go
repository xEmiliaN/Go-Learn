package main

const (
	_       = iota // 0
	Ok      = iota // 1
	Warning = iota // 2
	Error   = iota // 3
)

func main() {

	println(Ok)      // Output: 1
	println(Warning) // Output: 2
	println(Error)   // Output: 3
}
