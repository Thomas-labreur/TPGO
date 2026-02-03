package main
import (
 "fmt"
)

// Fonction à 2 paramètres
// Variables locales: 
func sousP1(p, taxe float64) float64 {
	// Calcul du prix taxé
	var prix = p * (1 + taxe/100) 

	// Affichage 
	fmt.Printf("p : %v \n", p) // Affiche le prix initial
	fmt.Printf("taxe : %v \n", taxe) // Affiche la taxe
	fmt.Printf("prix : %v \n", prix) // Affiche le prix taxé

	// Renvoie le prix taxé au programme appelant 
	return prix 
}

// Procédure à 2 paramètres
// Variables locales:
func sousP2(prix, l float64) {

	// Boucle avec variable locale l qui parcours les valeurs 0, 1, 2.
	for l := 0; l < 3; l++ { 
		prix = prix + 10 // Ajoute 10 au prix
		fmt.Printf("l : %v prix : %v \n", l, prix) // Affiche le nouveau prix
	}

	// Affichages 
	fmt.Printf("l : %v \n", l) // Affiche la valeur de la variable l (!= de celle de la boucle)
	fmt.Printf("prix final : %v \n", prix) // Affiche le prix final 

	// Message selon si le prix est < ou >= à la limite
	if prix < l { 
		fmt.Println("prix final inférieur à la limite")
	} else { 
		fmt.Println("prix final supérieur ou égal à la limite")
	}
}

// Programme principal
func main() {

	// Création des variables
	var ( 					
		prix, tva float64
	)

	// Affectation des variables
	prix = 12.5			
	tva = 5.5

	// Appel à la fonction sousP1
	p := sousP1(prix, tva)

	// Affichage de la valeur renvoyée par sousP1
	fmt.Printf("p : %v \n", p)

	// Affichage du prix initial 
	sousP2(prix, 100)
}
