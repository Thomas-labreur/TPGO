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

func plus_cher_article(n_articles int) float64 {
	// Déclarations
	var (
		prix_article float64 = 0.0
		plus_cher float64 = 0.0
	)

	// Boucle de 1 à n_articles
	for i:=1; i<=n_articles; i++ {

		// Demander le prix
		prix_article = demander_prix(i)

		// Si le prix est plus gros que le précédent, mettre a jour le prix le plus cher
		if prix_article > plus_cher {
			plus_cher = prix_article
		}
	}

	return plus_cher
}

func main() {

	// Déclarations
	var (
		plus_cher float64 = 0.0
		n_articles = demander_nombre_articles()
	)

	// Calcul du prix du plus cher article
	plus_cher = plus_cher_article(n_articles)

	// Affichage du prix du plus cher article
	fmt.Printf("Le plus cher article vaut %v€.", plus_cher)
}