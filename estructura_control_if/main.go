package main

import "fmt"

func main() {
	// if 5 < 6 {
	// 	fmt.Println("5 es menor que 6")
	// }
	a := 5
	if a < 6 {
		fmt.Println("5 es menor que 6")
	} else if a>6{
		fmt.Println("a es igual que 6")
	}else{
		fmt.Println("a es mayor que 6")
	}
	fmt.Println("fin del programa")
}
