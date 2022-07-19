package main

import "fmt"

func main() {
	//contador de numeros impares
	encabezado := `
	***************************
	Contador de numeros impares
	***************************
	`
	//Imprimimos el encabezado
	fmt.Println(encabezado)

	fmt.Println("Digita el primer numero")
	//decalramos la variable para almacenar el numero digitado
	var num1 int
	//leemos el numero digitado y lo guardamos en la variable num1
	fmt.Scanln(&num1)
	fmt.Println("Digita el segundo numero")
	//declaramos la variable para almacenar el segundo numero
	var num2 int
	//leemos el numero digitado y lo guardamos en la variable num2
	fmt.Scanln(&num2)
	//Declaramos la variable contador para guardar la cantidad de numeros impares
	var contador int
	//declaramos un bucle tomando como inicio y fin tomando los numeros digitados
	for i := num1; i <= num2; i++ {
		//evaluamos si el numero es impar
		if i%2 != 0 {
			//si el numero es impar
			//incrementamos el valor de la variable contador en 1
			contador++
			//imprimimos el numero impar
			fmt.Printf("%d es un numero impar\n", i)

		}
	}
	encabezado = `
	********************
	Resultado del conteo
	********************
	`
	//imprimimos el encabezado
	fmt.Println(encabezado)

	//Imprimir los resultados
	fmt.Printf("entre el %d y el %d hay %d numeros impares.", num1, num2, contador)

}
