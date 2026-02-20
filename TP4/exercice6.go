package main

import (
    "fmt"
	"strconv"
    )

// Je reprend ma fonction input du TP3
func input(message string) string {
	var resultat string
	fmt.Println(message)
	fmt.Scan(&resultat)
	return resultat
}

func demander_nb_notes() int {

	// Demander le nombre
	n_str := input("Combien de notes avez-vous?")
	n, err := strconv.Atoi(n_str)

	// Verifier le nombre
	for err != nil || n<1 {
		n_str = input("Nombre invalide! Combien de notes avez-vous?")
		n, err = strconv.Atoi(n_str)
	}
	return n
}

func demander_notes_coeff(n int) ([]float64, []float64) {
    
    // Déclarer variables
	var (
	    note_str string
		note float64
		coeff_str string
		coeff float64
		err error
	)
	
	// Initialiser les listes
	notes  := make([]float64, n)
	coeffs := make([]float64, n)
	
	// Parcours des listes
	for i, _ := range notes {
	    
	    // Demander note
		note_str = input("Entrer une note:")
		note, err = strconv.ParseFloat(note_str, 64)
		
		// Vérifier note
		for err != nil || note<0 || note>20 {
		    note_str = input("Note invalide! Entrer une note:")
		    note, err = strconv.ParseFloat(note_str, 64)
		}
		
		// Demander coefficient
		coeff_str = input("Entrer son coefficient:")
		coeff, err = strconv.ParseFloat(coeff_str, 64)
		
		// Verifier coefficient
		for err != nil || coeff<0 {
		    coeff_str = input("Coefficient invalide! Entrer son coefficient:")
		    coeff, err = strconv.ParseFloat(coeff_str, 64)
		    
		}
		    
		// Ajouter note et coefficient dans les listes
		notes[i] = note
		coeffs[i] = coeff
	}
	
	// Renvoi
	return notes, coeffs
}
    
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
	n := demander_nb_notes()
	notes, coeffs := demander_notes_coeff(n)
	fmt.Printf("Vos notes sont %v.\n", notes)
	fmt.Printf("Leurs coefficients sont %v.\n", coeffs)
	fmt.Printf("La moyenne est %v.\n", moyenne_ponderee(notes, coeffs))
}