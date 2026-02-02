package main

import ( 
  "fmt"
)

// fonction (car elle a un return) avec 2 param de type float
// var locales: prix type float
// appelé par le programme principal
// calcule un prix selon la taxe et affiche les variables avant de le retourner
func sousP1(p, taxe float64) float64 {
  var prix = p * (1 + taxe/100)
  fmt.Printf("p : %v \n", p)
  fmt.Printf("taxe : %v \n", taxe)
  fmt.Printf("prix : %v \n", prix)
  return prix
}

// procédure (car pas de return) avec 2 param de type float
// var locales: cpt (int); prix (float)
// appelé par le programme principal
// Affiche un message pour savoir si le prix final dépasse ou non une limite
func sousP2(prix, l float64) {
  var cpt int
  cpt = 0
  for cpt < 3 {
    prix = prix + 10
    fmt.Printf("cpt : %v prix : %v \n", cpt, prix)
    cpt = cpt + 1
  }
  fmt.Printf("l : %v \n", l)
  fmt.Printf("prix final : %v \n", prix)
  if prix < l {
    fmt.Println("prix final inférieur à la limite")
  } else {
    fmt.Println("prix final supérieur ou égal à la limite")
  }
}

// Programme principal
// variables locales prix, tva, p
func main() {
  var (prix, tva, p float64)
  prix = 12.5
  tva = 5.5
  p = sousP1(prix, tva)
  fmt.Printf("p : %v \n", p)
  sousP2(prix, 10)
}