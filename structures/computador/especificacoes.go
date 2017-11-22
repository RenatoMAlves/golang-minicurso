package computador

type Especificacoes struct { //exported struct  
    Fabricante string //exported field
    modelo string //unexported field
    Preco int //exported field
}