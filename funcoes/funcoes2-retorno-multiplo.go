package main

import (  
    "fmt"
)

func rectProporcoes(comprimento, altura float64)(float64, float64) {  
    var area = comprimento * altura
    var perimetro = (comprimento + altura) * 2
    return area, perimetro
}

func main() {  
     area, perimetro := rectProporcoes(10.8, 5.6)
    fmt.Printf("Area %f perimetro %f", area, perimetro) 
}