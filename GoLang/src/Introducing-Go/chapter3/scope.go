package main

import "fmt"

var x string = "Hello, World"

func main()  {
  // Go has also support for constants whose value can not be changed later
  const k string = "I am untouchable"
  // try to overwrite k = "No I am not"
  // results error try to rewrite - scope.go:10: cannot assign to k
  // fmt.Println(x)
  fmt.Println(k)

}

func f() {
  fmt.Println(x)
}
