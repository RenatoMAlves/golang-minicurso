package main

import (  
    "fmt"
)

type MembroLiga struct {  
    primeiroNome, ultimoNome string
    idade, salario         int
}

func main() {

    emp1 := MembroLiga{
        primeiroNome: "Bruce",
        idade:       37,
        salario:    50000000,
        ultimoNome:  "Wayne",
	}
	
    emp2 := MembroLiga{"Clark", "Kent", 31, 800}

    fmt.Println("Membro 1", emp1)
    fmt.Println("Membro 2", emp2)
}