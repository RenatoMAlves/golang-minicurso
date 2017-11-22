package main

import "fmt"

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a
    b[0] = "Singapore"
    fmt.Println("a é ", a)
    fmt.Println("b é ", b) 
}