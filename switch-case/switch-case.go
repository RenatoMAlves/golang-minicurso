package main

import (  
    "fmt"
)

func main() {  
    switch cor := 3; cor {
    case 1:
        fmt.Println("Vermelho")
    case 2:
        fmt.Println("Azul")
    case 3:
        fmt.Println("Verde")
    case 4:
        fmt.Println("Cinza")
    case 5:
        fmt.Println("Preto")
    default: 
        fmt.Println("Numero de cor invalido")
    }
}