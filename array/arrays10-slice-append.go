package main

import (  
    "fmt"
)

func main() {  
    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("carros:", cars, "tem tamanho antigo de", len(cars), "e capacidade", cap(cars))
    cars = append(cars, "Toyota")
    fmt.Println("carros:", cars, "tem novo tamanho de", len(cars), "e capacidade", cap(cars))
}