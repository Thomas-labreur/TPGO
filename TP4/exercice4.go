package main

import (
    "fmt"
    )

func indices_notes_superieures(notes []float64, ref float64) []int {
    // Ici on ne sait pas à l'avance la taille de la liste, donc on l'initialise vide et on la remplit avec append.
    resultat := []int{}
    for i, note := range notes {
        if note >= ref {
            resultat = append(resultat, i)
        }
    }
    return resultat
}

func main() {
    // Notes et coefficients
	notes := []float64{8, 9.4, 10, 14, 2, 7.5, 18, 6}
	var ref float64 = 1
	fmt.Printf("Les indices des notes supérieures à %v sont %v.\n",ref, indices_notes_superieures(notes, ref))
}