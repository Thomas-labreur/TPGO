package main

import ( 
  "fmt"
)

func max_entiers(a, b int) int {
  var max int
  if (a > b) {
    max = a
  } else {
    max=b
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
  fmt.Printf("max entre %v et %v: %v \n", a, b, max_entiers(a, b))
}