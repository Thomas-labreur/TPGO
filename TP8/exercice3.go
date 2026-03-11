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


func selectLibelleProd2(prods []produit, x int) (string, error) {
	
	// On parcours la liste, et on renvoie le libellé si le numProd correspond
	for _, prod := range prods {
		if prod.numProd == x {
			return prod.libelleProd, nil
		}
	}

	// Si on a pas trouvé le numProd, on renvoie une erreur
	return "", fmt.Errorf("Pas de produit correspondant.")
}

func main() {

	// Exemple de table Produits
	produits := []produit{
		{0, "souris", 10.35},
		{1, "ordinateur", 1200},
		{2, "piles", 2.99},
	}

    // Select libelleProd from produits where numProd=x
	x := 2
	libelle, err := selectLibelleProd2(produits, x)
	if err==nil {
		fmt.Printf("Le libellé du produit numero %v est '%v'.\n", x, libelle)
	} else {
	    fmt.Println(err)
	}
}