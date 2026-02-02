package main

import ( 
  "fmt"
)

// fonction à 1 param
// variables locales: produit, i (espace de nom de produit_entiers)
func  produit_entiers(n int) int {
  var (
    produit = 1
    i = 1
    )
  for (i <= n) {
    produit = produit* i
    i = i + 1
  }
  return produit
}

func demander_nombre() int {
  var n int
  fmt.Println("Entrez un nombre entier strictement positif:")
  fmt.Scanln(&n)
  for n<=0 {
    fmt.Println("Entrez un nombre entier strictement positif:")
    fmt.Scanln(&n) 
  }
  return n
}


// Programme principal
func main() {
  var n = demander_nombre()
  var pdt = produit_entiers(n)
  fmt.Printf("produit des entiers de 1 à %v: %v \n", n, pdt)
}
