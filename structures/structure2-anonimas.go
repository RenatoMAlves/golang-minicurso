package main

import (  
    "fmt"
)

func main() {  
    emp3 := struct {
        primeiroNome, ultimoNome string
		idade, salario         int
    }{
        primeiroNome: "Diana",
        ultimoNome:  "Prince",
        idade:       31,
        salario:    50000,
    }

    fmt.Println("Membro 3", emp3)
}