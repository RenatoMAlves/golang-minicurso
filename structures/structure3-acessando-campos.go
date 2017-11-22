package main

import (  
    "fmt"
)

type MembroLiga struct {  
    primeiroNome, ultimoNome string
    idade, salario         int
}

func main() {  
    emp4 := MembroLiga{"Barry", "Allen", 30, 6000}
    fmt.Println("Primeiro Nome:", emp4.primeiroNome)
    fmt.Println("Último Nome:", emp4.ultimoNome)
    fmt.Println("Idade:", emp4.idade)
	fmt.Printf("Salário: $%d", emp4.salario)
	
	var emp5 MembroLiga
    emp5.primeiroNome = "Dick"
    emp5.ultimoNome = "Grayson"
    fmt.Println("Membro 5:", emp5)
}