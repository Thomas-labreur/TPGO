package main

import "fmt"

// Création du type produit
type produit struct{
	numProd int
	libelleProd string
	prixProd float64
}

// Création du type ligne_commande
type ligne_commande struct{
	numCommande int
	numProd int
	quantite int
}

// Création du type commande
type commande struct{
	numCommande int
	dateCommande string
}

func selectLibelleProd(prods []produit) []string {

	// Initialiser une liste vide de string
	libelle_prods := make([]string, len(prods))

	// On la remplit avec les libelles des produits
	for i, prod := range prods {
		libelle_prods[i] = prod.libelleProd
	}

	return libelle_prods
}

func main() {

	// Exemple de table Produits
	produits := []produit{
		{0, "souris", 10.35},
		{1, "ordinateur", 1200},
		{2, "piles", 2.99},
	}

	// Select libelleProd from produits
	fmt.Printf("\nLes libellés des produits sont %v", selectLibelleProd(produits))
}