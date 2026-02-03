package main
import (
 "fmt";
 "strconv";
 "math"
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

func demander_limite() int {

	// Demande un nombre
	n, err := strconv.Atoi(input("Entrer une limite entière supérieure à 2: "))

	// Vérifie que c'est un entier supérieur à 2
	for (err != nil || n < 2) {
		fmt.Println("Saisie invalide !\n")
		n, err = strconv.Atoi(input("Entrer une limite entière supérieure à 2: "))
	} 

	// Renvoi à l'appelant
	return n
}

func somme_diviseurs_propres(n int) int {

	// Calcul de la racine carrée de n
	racine := int(math.Sqrt(float64(n)))  // Attention aux types !

	// Calcul de la somme des diviseurs propres
	somme := 1 // 1 est toujours diviseur propre
	for i:=2; i<=racine; i++ {
		if n%i==0 {			// i divise N si le reste de N/i est 0
			somme += i
			somme += n/i
		}
	}

	// Renvoi à l'appelant
	return somme
}

func main() {

	// Demande la limite
	limite := demander_limite()

	// Calcule la somme des diviseurs propres de tous les nombre sous la limite
	for i:=2; i<=limite; i++ {
		somme := somme_diviseurs_propres(n)
		if somme==i {
			Printf("%v ", i) // Affiche le nombre s'il est parfait
		}
	}

}