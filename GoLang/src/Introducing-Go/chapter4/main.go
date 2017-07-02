// The for statement
package main

import "fmt"

func main() {
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i += 1 // i = i + 1
  }

  // Second oprion also write same thing
  for i:= 1; i <=10; i++ {
    fmt.Println("Järgmine nr on:", i)
  }

  // if statement odd even number
  for i:=1; i<=10; i++ {
    if i%2 == 0 {
      //even
      fmt.Println(i, "even")
    } else {
      // odd
      fmt.Println(i, "odd")
    }
  }

  // else if syntaks demo
  for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
      // divisible by 2
      fmt.Println(i, "divisible by 2")
    } else if i % 3 == 0 {
      // divisible by 3
      fmt.Println(i, "divisible by 3")
    } else if i % 4 == 0 {
      // divisible by 4
      fmt.Println(i, "divisible by 4")
    }
  }
}
