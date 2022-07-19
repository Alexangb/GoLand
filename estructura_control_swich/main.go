package main

import "fmt"

func main() {
	//Imprimir el dia de la semana que el usuario eliga
	var dia int
	fmt.Println("Digita un numero del dia de la semana")
	fmt.Scanln(&dia)

	// if dia==1 {
	// 	fmt.Println("usted digito lunes")
	// }else if dia==2{
	// 	fmt.Println("usted digito martes")
	// }else if dia==3{
	// 	fmt.Println("usted digito miercoles")
	// }else if dia==4{
	// 	fmt.Println("usted digito jueves")
	// }else if dia==5{
	// 	fmt.Println("usted digito viernes")
	// }else if dia==6{
	// 	fmt.Println("usted digito sabado")
	// }else if dia==7{
	// 	fmt.Println("usted digito domingo")
	// }else{
	// 	fmt.Println("usted digito un numero invalido")
	// }

	switch dia {
	case 1:
		fmt.Println("usted digito el numero que pertenece al dia lunes")
	case 2:
		fmt.Println("usted digito el numero que pertenece al dia martes")
	case 3:
		fmt.Println("usted digito el numero que pertenece al dia miercoles")
	case 4:
		fmt.Println("usted digito el numero que pertenece al dia Jueves")
	case 5:
		fmt.Println("usted digito el numero que pertenece al dia viernes")
	case 6:
		fmt.Println("usted digito el numero que pertenece al dia sabado")
	case 7:
		fmt.Println("usted digito el numero que pertenece al dia domingo")
	default:
		fmt.Println("usted digito un numero incorrecto")
	}

}
