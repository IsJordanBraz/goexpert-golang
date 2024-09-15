package main

func main() {
	var teste interface{} = "Jordan"

	res, err := teste.(int)

	if err {
		println("ERROR")
	}

	println(res)
}
