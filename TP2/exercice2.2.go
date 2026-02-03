package main
import (
 "fmt"
)

func demander_nombre_articles() int {
	// Déclaration
	var n_articles int

	// Saisie du nombre d'articles
	fmt.Printf("Combien d'articles avez vous? \n")
	fmt.Scan(&n_articles)

	// Vérification du nombre d'articles
	for (n_articles <= 0 || n_articles >10) {
		fmt.Printf("Vous ne pouvez avoir qu'entre 1 et 10 articles.\n")
		fmt.Printf("Combien d'articles avez vous? \n")
		fmt.Scan(&n_articles)
	}

	// Renvoi du nombre d'articles à l'appelant
	return n_articles
}

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

	// Declarations
	var (
		somme float64 = 0.0
		n_articles = demander_nombre_articles()
	)
	
	// Boucle de 1 à n_articles
	for i:=1; i<=n_articles; i++ {
		somme += demander_prix(i) // demande le prix du ième article et l'ajoute à la somme
	}

	// Affichage du résultat
	fmt.Printf("Le prix total est %v€.", somme)
}