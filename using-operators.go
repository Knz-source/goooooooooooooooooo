package main

import "fmt"

func main() {
//define variable as 
    var a int = 5
    var b float32 = 4.32
    const pi float64 = 3.1415139475
//declare two variables together
    var (
    e = 8
    g = 7
    )
//declaring multiple variables again 
    x,y := 14,15

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(pi)
    fmt.Println(e, "," ,g)
    fmt.Println(x, "," ,y)
}


