package main

import (
    "fmt"
    )
    
func moyenne_ponderee(notes []float64, coeff []float64) float64 {
    
    // Déclarations des variables
    var (
	    somme_ponderee float64 = 0
	    somme_coeff float64 = 0
	    )
	    
	// Parcours des listes
	for i, note := range notes {
		somme_ponderee += note*coeff[i]
		somme_coeff += coeff[i]
	}
	
	// Calcul de moyenne
	moyenne := somme_ponderee / somme_coeff
	return moyenne
}

func main() {
	notes := []float64{8, 9.4, 10, 14, 2, 7.5, 18, 6}
	coeff := []float64{2, 2, 1, 2, 2, 1, 2, 1}
	fmt.Printf("La moyenne est %v.", moyenne_ponderee(notes, coeff))
}