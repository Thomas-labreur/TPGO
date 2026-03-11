package main

import (
	"fmt"
	"strconv"
)

///////////////////////////////////////////////////////////
// Les deux fonctions suivantes n'étaient pas demandées. //
///////////////////////////////////////////////////////////

func input(message string) string {

	// Déclaration
	var resultat string

	// Lecture de la console
	fmt.Println(message)
	fmt.Scan(&resultat)

	// Renvoi du résultat
	return resultat
}

func saisirTVA() float64 {
	// Demander tva et convertir en float64
	tva_str := input("Entrer un taux de TVA (5.5, 10 ou 20):")
	tva, err := strconv.ParseFloat(tva_str, 64)

	// Vérification
	for err != nil || (tva!=5.5 && tva!=10 && tva!=20) {
		tva_str = input("Taux invalide ! Entrer un taux de TVA (5.5, 10 ou 20):")
		tva, err = strconv.ParseFloat(tva_str, 64)
	}

	return tva
}

//////////////////////
// Travail demandé ///
//////////////////////

// SUJET B et C
func saisirNombre() int {
	// Demander un nombre et convertir en int
	n_str := input("Entrer un nombre entier strictement positif:")
	n, err := strconv.Atoi(n_str)

	// Vérification
	for err != nil || n <= 0 {
		n_str = input("Entrer un nombre entier strictement positif:")
		n, err = strconv.Atoi(n_str)
	}

	return n
}

// SUJET A ET C
func saisirPrix() float64 {
	// Demander un prix et convertir en float64
	prix_str := input("Entrer un prix strictement positif:")
	prix, err := strconv.ParseFloat(prix_str, 64)

	// Vérification
	for err != nil || prix <= 0 {
		prix_str = input("Prix invalide ! Entrer un prix strictement positif:")
		prix, err = strconv.ParseFloat(prix_str, 64)
	}

	return prix
}

// SUJET C ET D
func stockerPrix(phts []float64) {

	// On parcours la liste et on saisit un prix dans chaque élément.
	for i, _ := range phts {
		phts[i] = saisirPrix()
	}
}

// SUJET A ET B
func calculerTTC(phts []float64, tva float64) []float64 {

	// Nouvelle liste qui contiendra les prix ttc
	pttcs := make([]float64, len(phts))

	// On parcours la liste des prix ht et on calcule le ttc
	for i, pht := range phts {
		pttcs[i] = pht * (1+tva/100)
	}

	return pttcs
}

////////////////////////////////////////////////////////////////////////
// Programme principal pour tester les sous-programmmes (non demandé) //
////////////////////////////////////////////////////////////////////////

func main() {

	// Saisir nombre d'articles et tva
	n_articles := saisirNombre()
	tva := saisirTVA()

	// Saisir les prix ht
	phts := make([]float64, n_articles)
	stockerPrix(phts)

	// Calculer les prix ttc
	pttcs := calculerTTC(phts, tva)
	fmt.Printf("Avec un taux de %v%%, les prix de vos articles sont:\n", tva)
	for _, p := range pttcs {
		fmt.Printf("%.2f ", p)
	}
}