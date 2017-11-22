package main

import ( 
	"fmt"
)
func main(){
	numeros := []int{10,9,7,8,7,6,2,7,4,10}
	var resultado int
	for cont:=0; cont < 10; cont ++{
		resultado = resultado + numeros[cont]
	}

	fmt.Println(int(resultado)/10)
}