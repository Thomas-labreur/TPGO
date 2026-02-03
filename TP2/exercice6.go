package main
import (
 "fmt"
)


func demander_prix(numero_article float64) float64 {
	// Déclaration
	var prix float64

	// Saisie du prix
	fmt.Printf("Entrer le prix de l'article %v ou 0 si vous souhaitez arrêter:", numero_article)
	fmt.Scan(&prix)

	// Vérification du prix saisi
	for prix < 0 {
		fmt.Printf("Prix invalide, entrez le prix de l'article %v :", numero_article)
		fmt.Scan(&prix)
	}

	// Renvoi du prix à l'appelant
	return prix
}

func main() {

	// Declarations
	var (
		somme float64 = 0.0
		prix_article float64 = -1
		i float64 = 1
	)
	
	// Boucle qui s'arrete si
	for prix_article !=0 {
		prix_article = demander_prix(i)
		somme += prix_article 
		i++
	}

	// Affichage du résultat
	fmt.Printf("La moyenne des prix de vos articles est %v€.", somme/(i-2))
}