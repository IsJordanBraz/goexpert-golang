package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	value1 := 10
	value2 := 20
	soma(&value1, &value2)
	println(value1)
}
