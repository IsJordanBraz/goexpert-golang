package main

func main() {
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Infinito")
	}
}
