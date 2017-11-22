package main

import (  
    "fmt"
)

func main() {  
    b := 255
    var a *int = &b
    fmt.Printf("Tipo de a é %T\n", a)
    fmt.Println("O endereço na memória de b é ", a)
}