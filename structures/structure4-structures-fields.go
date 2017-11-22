package main

import (  
    "fmt"
)

type Endereco struct {  
    cidade, estado string
}
type Pessoa struct {  
    nome string
    idade int
    endereco Endereco
}

func main() {  
    var p Pessoa
    p.nome = "Tony Stark"
    p.idade = 50
    p.endereco = Endereco {
        cidade: "Nova York",
        estado: "Nova York",
    }
    fmt.Println("nome:", p.nome)
    fmt.Println("idade:",p.idade)
    fmt.Println("cidade:",p.endereco.cidade)
    fmt.Println("estado:",p.endereco.estado)
}