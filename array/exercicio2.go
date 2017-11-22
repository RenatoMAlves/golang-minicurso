package main

import ( 
	"fmt"
)
func main(){
	numeros := []int{10,9,7,8,7,6,2,7,4,10}
	menorValor := 10
	maiorValor := 0
	for cont:=0; cont < 10; cont ++{
		if numeros[cont] > maiorValor{
			maiorValor = numeros[cont]
		}

		if numeros[cont] < menorValor{
			menorValor = numeros[cont]
		}
	}

	fmt.Println(menorValor)
	fmt.Println(maiorValor)
}