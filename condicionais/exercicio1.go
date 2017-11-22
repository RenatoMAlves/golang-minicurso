package main

import 
	"fmt"

func main(){
	fmt.Println("Digite um numero qualquer:")
	var num int
	fmt.Scanf("%d", &num)

	if num % 2 == 0{
		fmt.Println("Este numero é par")
	} else{
		fmt.Println("Este numero é impar")
	}

}