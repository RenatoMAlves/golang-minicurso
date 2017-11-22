package main

import (  
    "fmt"
)

func main() {  
    a, b := 5.67, 8.97
    fmt.Printf("o tipo de a é %T e b %T\n", a, b)
    soma := a + b
    diferenca := a - b
    fmt.Println("soma", soma, "diferença", diferenca)

    no1, no2 := 56, 89
    fmt.Println("soma", no1+no2, "diferenca", no1-no2)
}