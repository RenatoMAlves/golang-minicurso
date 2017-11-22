package main

import (  
    "fmt"
)

func main() {  
    num := 99
    if num <= 50 {
        fmt.Println("o numero é menor ou igual 50")
    } else if num >= 51 && num <= 100 {
        fmt.Println("o numero está entre 51 e 100")
    } else {
        fmt.Println("o numero é maior que 100")
    }

}