package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //cria um slice de a[1] até a[3]
    // nums1 := a[:] //creates a slice which contains all elements of the array
    fmt.Println(b)
}

// func main() {  
//     c := []int{6, 7, 8} //creates and array and returns a slice reference
//     fmt.Println(c)
// }