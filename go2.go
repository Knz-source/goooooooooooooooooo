package main

import "fmt"

//variable is name given to storage area that can programs could manipulate.
//constants fixed values that cannot be changed
func main() {

  const pi float64 = 3.14159265359
  // could declare multiple variables like this

  var (
    varA = 2
    varB = 3
  )

  fmt.Println(varA, varB)
  //formating the print len of both vars and the vars are being non equals values - variables
  
  //strings are a serie of chars between " or ` like var example string = "stringhere Here" e.g
  //we could define vars like this with strings
  var Name string =  "meu nome"

  //and then we could get the length of the string 
  fmt.Println(len(Name))
}
