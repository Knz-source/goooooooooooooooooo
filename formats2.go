package main 

import "fmt" 

func main(){

  var name string = "Study Guide" //strings
  const pi float64 = 3.14123264 //floats
  win := true //booleans
  x := 80

  //using println to lead with this example. first we get the lenght of Study Guide = 11 and then add the text with the followed sentence. 
  fmt.Println(len(name))
  fmt.Println(name + " - Golang Lessons Understanding format")

  fmt.Printf("%f \n", pi)
  fmt.Printf("%.4f \n", pi)
  fmt.Printf("%T \n", name)
  fmt.Printf("%T \n", pi)
  fmt.Printf("%t \n", win)
  fmt.Printf("%d \n", x)
  fmt.Printf("%b \n", 25)
  fmt.Printf("%c \n", 33)
  fmt.Printf("%x \n", 15)
  fmt.Printf("%e \n", pi)

}

//using printf - format printing defining %f to lead with float numbers - decimal point but no exponent e.g. 123.456 and could use %F as the same.  
//%.xf <- if the const pi atributed value = 3.14123264 the x will be the location - 3f = 142 
//%T ask the type of the attributed value as example = name == string and pi == floats numbers.
//%t uses to format true or false values - booleans operations
//%e scientific notation 
//%x hexadecimal notation
//%b decimalles scientific notation with exponent a power of two.
//%c the char represented by the corresponding unicode code point 
//%d base 10
