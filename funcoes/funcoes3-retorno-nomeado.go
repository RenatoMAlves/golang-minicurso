package main

import (  
    "fmt"
)

func rectProporcoes(comprimento, altura float64)(area, perimetro float64) {  
    area = comprimento * altura
    perimetro = (comprimento + altura) * 2
    return //nenhum valor explícito
}

func main() {  
	 area, perimetro := rectProporcoes(10.8, 5.6)
	//  area, _ := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f perimetro %f", area, perimetro) 
}