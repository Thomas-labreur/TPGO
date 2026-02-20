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
    
    // Notes et coefficients
	notes := []float64{8, 9.4, 10, 14, 2, 7.5, 18, 6}
	coeff := []float64{2, 2, 1, 2, 2, 1, 2, 1}
	
	// Notes et coeff d'informatique
	notes_info := notes[2:6] // uniquement les notes 2, 3, 4 et 5
	coeff_info := coeff[2:6]
	notes_info[len(notes_info)-1] =10
	
	fmt.Printf("La moyenne est %.2f.\n", moyenne_ponderee(notes, coeff))
	fmt.Printf("La moyenne en informatique est %.2f.\n", moyenne_ponderee(notes_info, coeff_info))

    // Les DEUX moyennes ont changé, car notes_info est une vue de notes, pas une nouvelle variable indépendante, si je modifie l'un, je modifie l'autre aussi !!
}