package main

func main() {
	a := 11
	var ponteiro *int = &a
	*ponteiro = 30
	b := &a
	*b = 20
	println(a)
}
