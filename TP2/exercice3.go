package main
import (
 "fmt"
)

func calculer_prix_réduit(total float64) float64 {
	// Déclarations
	var (
		prix_final float64
		reduction float64
	)

	// Calcul du taux de réduction
	switch {
		case total > 50:
			reduction = 20
		case total >= 30:
			reduction = 15
		case total >= 20:
			reduction = 10
		default:
			reduction = 0
		}

	// Affichage du message et calcul du prix réduit
	fmt.Printf("Vous bénéficiez de %v %% de remise!\n", reduction)
	prix_final = total * (1-reduction/100)

	// Renvoi à l'appelant
	return prix_final
}

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
	// Déclarations
	var (
		prix float64 = 0.0
		n_articles = demander_nombre_articles()
	)

	// Boucle de 1 à n_articles
	for i:=1; i<=n_articles; i++ {
		prix += demander_prix(i) // Demande le prix du ième article
	}

	// Affichage du total
	fmt.Printf("Le total est %.2f€.\n", prix)

	// Calcule et affichage du prix réduit
	prix = calculer_prix_réduit(prix)
	fmt.Printf("Le prix final après réduction est %.2f€.\n", prix)
}