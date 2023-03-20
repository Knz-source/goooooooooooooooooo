package main 

import "fmt"

func main(){

//using printf as examples

  const pi float64 = 3.16412732
  x := 5
  isbool := true 

//%f lets ask format print a float number
//%T lets show the type 
//%t using boolean
//%d used for showing integers 
  fmt.Printf(" %f \n", pi)

  fmt.Printf(" %T \n", isbool)

  fmt.Printf(" %t \n", pi)

  fmt.Printf(" %d \n", x)

}
