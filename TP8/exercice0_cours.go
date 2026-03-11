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

// La table Produits peut être modélisée par un liste de produits []produit

func afficherProduits(prods []produit) {

	// On affiche le nom et le prix de chaque produits
	for _, prod := range prods {
		fmt.Printf("%v (%v euros)\n", prod.libelleProd, prod.prixProd)
	}
}

func main() {

	// Exemple de table Produits
	produits := []produit{
		{0, "souris", 10.35},
		{1, "ordinateur", 1200},
		{2, "piles", 2.99},
	}

	// Affichage
	afficherProduits(produits)
}