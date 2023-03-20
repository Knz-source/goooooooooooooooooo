
package main

import "fmt"
//leading with pointers use & <-
//define pointer to x. 
// *x = new pointer value
func main() {
  
  x := 10
  
  changeValue(&x)
  fmt.Println(x) 

}

func changeValue(x *int) {


  *x = 7;

}
