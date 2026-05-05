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

// Pour le résultat du inner join
type ligne_commande_produit struct {
	numCommande int
	numProd int
	quantite int
	libelleProd string
	prixProd float64
}

func select_liste_lignes( liste_produits []produit,  liste_lignes []ligne_commande, x int) []ligne_commande_produit {

	// Select * -> on renvoie toutes les  lignes de commandes, donc une liste
	resultat := []ligne_commande_produit{}

	// A chaque ligne, on parcours prod a la recherche du produit qui correspond
	for _, ligne := range  liste_lignes {

		// On ajoute la condition ligne.numCommande==x, de cette façons si le numero
		// de commande n'est pas le bon, on ne parcours meme pas liste_produits, on évite donc
		// des instructions inutiles.
		for i:=0 ; i<len(liste_produits) && ligne.numCommande==x ; i++ {

			// Si on trouve un produit correspondant, on doit créer l'objet ligne_commande_produit
			if liste_produits[i].numProd == ligne.numProd {
				res := ligne_commande_produit{
					numCommande: ligne.numCommande,
					numProd: ligne.numProd,
					quantite: ligne.quantite,
					libelleProd: liste_produits[i].libelleProd,
					prixProd: liste_produits[i].prixProd,
				}
				resultat = append(resultat, res)
			}
		}
	}
	return resultat
}

func main() {

	// Exemple de table Produits
	liste_produits := []produit{
		{0, "souris", 10.35},
		{1, "ordinateur", 1200},
		{2, "piles", 2.99},
	}

	// Exemple de table LigneCommandes
	liste_lignes := []ligne_commande{
		{0, 0, 1},
		{0, 0, 1},
		{1, 0, 8},
		{2, 1, 94},
		{2, 2, 1},
		{3, 2, 10},
		{4, 1, 1},
	}

	// Select * from  lignes_commandes inner join produits using numProd where numCommande=x
	x:=2
	lignes_select := select_liste_lignes(liste_produits, liste_lignes, x)
	fmt.Printf("Les produits commandés par la commande numero %v sont:\n", x)
	for _, res := range lignes_select {
		fmt.Printf("%v\n", res)
	}
}