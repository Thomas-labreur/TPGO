package main
import (
 "fmt"
)

func demander_prix(numero_article int) float64 {
	// Déclaration
	var prix float64

	// Saisie du prix
	fmt.Printf("Entrer le prix de l'article %v :", numero_article)
	fmt.Scan(&prix)

	// Vérification du prix saisi
	for prix <= 0 {
		fmt.Printf("Prix invalide, entrez le prix de l'article %v :", numero_article)
		fmt.Scan(&prix)
	}

	// Renvoi du prix à l'appelant
	return prix
}

func main() {
	// Déclaration
	var somme float64 = 0.0

	// Boucle de 1 à 5
	for i:=1; i<=5; i++ {
		somme += demander_prix(i) // Demande le prix du ième article
	}

	// Affichage de la somme
	fmt.Printf("Le prix total est %v.", somme)
}