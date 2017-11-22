package main 

import (  
    "fmt"
    "geometry/rectangle"
)

func main() {  
    var rectComprimento, rectAltura float64 = 6, 7
    fmt.Println("Propriedades de formas geométricas")
		/*
		* %.2f é duas casas decimais
        */
    fmt.Printf("área do retangulo %.2f\n", rectangle.Area(rectComprimento, rectAltura))
        
    fmt.Printf("diagonal do retangulo %.2f ",rectangle.Diagonal(rectComprimento, rectAltura))
}