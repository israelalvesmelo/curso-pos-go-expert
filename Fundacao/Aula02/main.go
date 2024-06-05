package main

const wording = "Hello, World"

var (
	b  bool = true
	i  int
	s  string
	f1 float32
	f2 float64
)

func main() {
	str := "test"
	println(wording)
	println(str)
	println(b)
	println(i)
	println(s)
	println(f1)
	println(f2)
	b = false
	println(b)

}
