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

	// Demande l'age et vérifie le type int
    age_str := input("Quel âge as-tu?")
    age, err := strconv.Atoi(age_str)						// Rque: ici on déclare avec :=
	for err != nil {
	    fmt.Println("Erreur : l'âge doit être un entier.\n")
		age_str = input("Quel âge as-tu?")
		age, err = strconv.Atoi(age_str)					// Rque: ici on affecte avec =
	}
	
	// Demande la taille et verifie le type float64
    taille_str := input("Quelle est ta taille (en m) ?")
    taille, err := strconv.ParseFloat(taille_str, 64)
	for err != nil {
	    fmt.Println("Erreur : la taille doit être un nombre réel.\n")
		taille_str = input("Quelle est ta taille (en m) ?")
		taille, err = strconv.ParseFloat(taille_str, 64)
	}

	fmt.Printf(
		"Bonjour %v, tu as %v ans et tu mesures %.2fm.\n",
		nom,
		age,
		taille,
	)
}