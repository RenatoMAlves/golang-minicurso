package main

import (  
    "fmt"
    "unsafe"
)

func main() {  
    var a int = 89
    b := 95
    fmt.Println("valor de a é ", a, " e valor de b é ", b)
    fmt.Printf("o tipo de a é %T, o tamanho de a é %d", a, unsafe.Sizeof(a))
    fmt.Printf("\no tipo de b é %T, o tamanho de b é %d", b, unsafe.Sizeof(b)) 
}