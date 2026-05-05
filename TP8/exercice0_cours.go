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

func afficherProduits(liste_produits []produit) {

	// On affiche le nom et le prix de chaque produits
	for _, p := range  liste_produits {
		fmt.Printf("%v (%v euros)\n", p.libelleProd, p.prixProd)
	}

	// ATTENTION de bien faire la différence entre ces 3 éléments:
	// produit est un type 
	// liste_produits est une variable du type liste de produit
	// p est un élément de cette liste, c'est donc une variable du type produit
}

func main() {


	// Exemple de table Produits
	liste_produits := []produit{
		{0, "souris", 10.35},
		{1, "ordinateur", 1200},
		{2, "piles", 2.99},
	}

	// Affichage
	afficherProduits(liste_produits)
}

