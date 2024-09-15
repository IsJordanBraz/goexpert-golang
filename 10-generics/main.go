package main

type MyNumber int

type MySyum interface {
	~int | ~float64
}

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma2[T MySyum](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T MySyum](a T, b T) bool {
	if a > b {
		return true
	}
	return false
}

func main() {
	m1 := map[string]int{"Jordan": 100, "Ana": 250}
	m2 := map[string]float64{"Jordan": 100.00, "Ana": 250.50}

	println(SomaInteiro(m1))
	println(SomaFloat(m2))
	println(Soma(m1))
	println(Soma(m2))

	println(Compara(10, 10.0))
}
