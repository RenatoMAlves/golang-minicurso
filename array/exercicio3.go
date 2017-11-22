package main

import ( 
	"fmt"
)
func main(){
	b := []int{10,9,7,8,7,6,2,7,4,10}
	a := []int{10,9,7,8,7}
	c := []int{}
	var aux bool

	for cont:=0; cont < len(b); cont ++{
		aux = true
		for cont2:=0; cont2 < len(a); cont2++{
			if b[cont] == a[cont2]{
				aux = false
			} 
		}

		if aux == true{
			c = append(c, b[cont])
		}
	}

	fmt.Println(c)
}