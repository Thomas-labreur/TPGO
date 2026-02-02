package main

import ( 
  "fmt"
)

func max_entiers(a, b, c int) int {
  var max int
  if (a >= b && a >= c)  {
    max = a
  } else if (b >= a && b >= c) {
    max=b
  } else {
    max = c
  }
  return max
}

func demander_nombre() int {
  var n int
  fmt.Println("Entrez un nombre entier:")
  fmt.Scanln(&n)
  return n
}

// Programme principal
func main() {
  var a = demander_nombre()
  var b = demander_nombre()
  var c = demander_nombre()
  fmt.Printf("max entre %v, %v et %v: %v \n", a, b, c, max_entiers(a, b, c))
}