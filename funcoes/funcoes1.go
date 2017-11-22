package main

import (  
    "fmt"
)

func calcularConta(preco int, qtd int) int {  
    var totalpreco = preco * qtd
    return totalpreco
}

func main() {  
    preco, qtd := 90, 6
    totalpreco := calcularConta(preco, qtd)
    fmt.Println("Total preco é:", totalpreco)
}

