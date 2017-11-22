package main

import (  
    "fmt"
)

func main() {  
    vegetais := []string{"batatas","tomates","berinjelas"}
    frutas := []string{"laranjas","maçãs"}
    comida := append(vegetais, frutas...)
    fmt.Println("comidas:",comida)
}