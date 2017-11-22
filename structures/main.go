package main

import "./computador"  
import "fmt"

func main() {  
    var especif computador.Especificacoes
    especif.Fabricante = "apple"
    especif.Preco = 50000
    fmt.Println("Fabricante:", especif.Fabricante)
    fmt.Println("Preço:", especif.Preco)
}