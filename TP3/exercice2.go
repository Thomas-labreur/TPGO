package main
import (
 "fmt";
 "strconv"
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

	// Demande le nom
	nom := input("Bonjour, comment t'appelles-tu?")

	// Demande l'age et converti en int
	age_str := input("Quel âge as-tu?")
	age, _ := strconv.Atoi(age_str)

	// Demande la taille et converti en float64
	taille_str := input("Quelle est ta taille (en m) ?")
	taille, err := strconv.ParseFloat(taille_str, 64)

	// Affiche le resultat
	fmt.Printf(
		"Bonjour %v, tu as %v ans et tu mesures %.2fm.\n",nom,age,taille)
}