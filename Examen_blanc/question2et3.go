package main
package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

func code_correct(code string) bool {
	codeR := []rune(code)  // Pas obligatoire car les caractères '1' à '9' sont codés sur 1 octet
	somme := 0
	n:= len(codeR)

	for i:=0; i<n; i++ {
		if i%2==0 {
			// ATTENTION A LA CONVERSION CARACTERE VERS ENTIER
			// Par ex '1' est codé en ASCII par 49 donc int('1') donne 49.
			// Pour le convertir, on va donc faire int('1'-'0') ce qui donne 49-48 donc 1.
			// On peut aussi utiliser strconv.Atoi('1') (sans gestion de l'erreur, on le fera dans le main)
			somme += int(codeR[n-1-i] - '0') 
		} else {
			somme += (int(codeR[n-1-i] - '0')*2)%9
		}
	}

	return somme%10 == 0
}

func main () {

	// Lecture de la saisie
	scanner := bufio.NewScanner(os.Stdin)  
	fmt.Println("Entrez un code numérique sans espaces: ")
	scanner.Scan()  // fmt.Scan possible, mais pb si l'utilisateur met des espaces (seul la partie avant le 1er espace sera prise en compte)
	code := scanner.Text()

	// Vérification du code
	code_int, err := strconv.Atoi(code) // On tente de convertir pour vérifier si le code ne contient que des chiffres. Si c'est le cas, la conversion devrait marcher.
	for err!=nil || code_int<0 { 
		fmt.Println("Code invalide! Entrez un code numérique sans espaces: ")
		scanner.Scan() 
	    code = scanner.Text()
		code_int, err = strconv.Atoi(code)
	}

	// Affichage du résultat
	if code_correct(code) { // On applique bien à code (string) et pas a code_int ! code_int sert uniquement à vérifier la saisie
		fmt.Printf("Le code '%v' est correct!", code)
	} else {
		fmt.Printf("Le code '%v' n'est pas correct!", code)
	}
}