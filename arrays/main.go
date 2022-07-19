package main

import "fmt"

func main() {

	// var notas int

	var x [3]int
	fmt.Println(x)

	var k [3][2]float64
	fmt.Println(k)

	//asignar valo a un array
	x[1] = 25
	fmt.Println(x)

	//acceder a un indice correcto

	fmt.Println(x[1])

	//declarar e inicializar un array 1

	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)
	//declarar e inicializar arrays 2
	j := [...]int{1, 2, 3, 4, 5}
	fmt.Println(j)
	//declarar e inicializar arrays 3
	i := [...]int{
		1,
		2,
		3,
		4,
	}
	fmt.Println(i)
	//arrays de cualquier tipo
	f := [...]float64{1.2, 3, 1.5, 8.3}
	fmt.Println(f)

	//funcion len() indica el tamaño del array
	fmt.Println(len(f))
	//imprimir el ultimo elemento de un array
	fmt.Println(len(f) - 1)
	//comparar arrays
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	// d:=[4]int{1,2,3}

	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(c == b)
}
