// variables
package main

import "fmt"

func main()  {
  var x string = "Hello, World"
  var myName string
  myName = "My name is Kaido"

  fmt.Println(x)
  fmt.Println(myName)

  // Variables can overrun during program execution
  var h string
  h = "first"
  fmt.Println(h)
  h = h + "second" // can also be writen x+= "second"
  fmt.Println(h)

  // == is equality and turns boolean
  var p string = "hello"
  var l string = "world"
  fmt.Println(p == l)

  // also we can assign Variables
  k := 5
  j := "Make my day!"
  fmt.Println(k)
  fmt.Println(j)
}
