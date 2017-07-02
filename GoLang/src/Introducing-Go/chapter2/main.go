package main

import "fmt"

func main(){
  fmt.Println("1 + 1=",1.0 + 1.0)
  fmt.Println("My name is Kaido")
  fmt.Println(len("Hello, World"))
  fmt.Println("Hello World"[1])
  fmt.Println("Hello, " + "World")
  fmt.Println(true && true)
  fmt.Println(true && false) // and operator
  fmt.Println(true || true) // or operator
  fmt.Println(true || false) // or operator
  fmt.Println(!true) // not operator
}
