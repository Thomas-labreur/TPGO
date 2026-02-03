package main
import (
 "fmt"
)

func input(message string) string {

	// Déclaration
	var resultat string

	// Lecture de la console
	fmt.Println(message)
	fmt.Scan(&resultat)

	// Renvoi du résultat
	return resultat
}

func main() {

	// Demande du nom 
	var nom = input("Bonjour, comment t'appelles-tu?")

	// Affichage
	fmt.Printf("Enchanté %v.\n", nom)
}