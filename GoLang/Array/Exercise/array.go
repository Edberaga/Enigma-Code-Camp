package main

import (
    "fmt"
)
 
func main() {
    var elems int
    
    //fmt.Print("Number of elements? ")
    fmt.Scan(&elems)
    
    var array = make([]int,elems)
    
    for i := 0; i < elems; i++ {
        //fmt.Printf("%d . Number? ", i+1)
        fmt.Scan(&array[i])
    }

    for i := range array {
        if(array[i] % 2 == 0) {
            fmt.Println(array[i])
        }
    }
}