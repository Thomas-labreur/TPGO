package main

import (
    "fmt"
    )
    
func reussite(notes []float64, coeff []float64) {
    
    // Déclarations des variables
    var (
	    somme_ponderee float64 = 0
	    somme_coeff float64 = 0
	    )
	    
	// Parcours des listes
	i:=0
	for i<len(notes) &&  notes[i]>=7 {
		somme_ponderee += notes[i]*coeff[i]
		somme_coeff += coeff[i]
		i++
	}
	
	if i<len(notes) {
	    fmt.Printf("Tu as eu un %v, dommage.\n", notes[i])
	} else {
	    fmt.Printf("Toutes tes notes sont supérieures à 7, bravo! \n")
	}
	
	// Calcul de moyenne
	moyenne := somme_ponderee / somme_coeff
	fmt.Printf("Ta moyenne est %.2f.\n", moyenne)
	
	if i == len(notes) && moyenne >= 10 {
	    fmt.Println("Bravo, tu as réussi ton année !\n")
	} else {
	    fmt.Println("Oh non :'( essaie encore, tu peux y arriver !\n")
	}
}

func main() {
	notes := []float64{8, 9.4, 10, 14, 7, 7.5, 18, 8}
	coeff := []float64{2, 2, 1, 2, 2, 1, 2, 1}
	reussite(notes, coeff)
	
}