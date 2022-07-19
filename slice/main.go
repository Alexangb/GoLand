package main

import "fmt"

func main() {
	//declarar slice
	var j []int
	fmt.Println("slice ",j)
	//declarar slice 2
	x:=[]int {1,2,3}
	fmt.Println(x)
	//declarar slice con make, indicando la longitud
	y:=make([]int,5)
	fmt.Println("slice y", y)
	fmt.Println("longitud de y:" ,len(y))
	fmt.Println("capacidad de y: ",cap(y))

	//declarar slice con make,indicando la longitud y la capacidad
	k:=make([]int,5,10)
	fmt.Println("slice k:",k)
	fmt.Println("longitud de k: ",len(k))
	fmt.Println("capaciad de k: ",cap(k))

	//definimos un array con 10 elementos de tipo int
	var ar=[10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println("arrar ar: ",ar)
	//definimos 2 slice de tipo int
	var a,b[]int
	fmt.Println("slice a:", a)
	fmt.Println("slice b: ",b)
}