package main

import "fmt"

func main() {  
    a, b := 20, 30 // declare variables a and b
    fmt.Println("a é", a, "b é", b)
    b, c := 40, 50 // b is already declared but c is new
    fmt.Println("b é", b, "c é", c)
    b, c = 80, 90 // assign new values to already declared variables b and c
    fmt.Println("mudança de b é", b, "c é", c)
}	