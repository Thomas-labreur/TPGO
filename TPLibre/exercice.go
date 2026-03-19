package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

func checkImage(image string) bool {

	// Cas 1: image vide
	imageR := []rune(image)
	if len(imageR) == 0 {
		return false
	}

	// Cas 2: autre caractere que B et W
	for _, r := range imageR {
		if r!='B' && r!='W' {
			return false
		}
	}
	return true
}

func saisirImage() string {
	image := input("Veuillez saisir une image (suite de W et B): ")
	for !checkImage(image) {
		image = input("Saisie invalide! Veuillez saisir une image (suite de W et B): ")
	}
	return image
}

func compresserImage(image string) string {

	// Declarations et initialisations des variables
	imageR := []rune(image)
	resultat := ""
	i, count := 0, 0
	var premier_caractere rune

	// Parcours de la chaine
	for i<len(imageR) {

		// Sauvegarde du 1er caractère
	    premier_caractere = imageR[i]

		// Parcours des suivants en augmentant le compteur
		for i<len(imageR) && imageR[i] == premier_caractere {
			count++
			i++
		}
		resultat += strconv.Itoa(count) // Convertion entier vers string
		resultat += string(premier_caractere)   // Conversion rune vers string
		count = 0
	}
	return resultat
}

func main() {
	image := saisirImage()
	fmt.Printf("Image de base: %v\n", image)
	fmt.Printf("Compression: %v", compresserImage(image))
}