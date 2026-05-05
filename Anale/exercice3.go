package main
import (
    "fmt"
    "bufio"
    "os"
    )

func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

func extraire_nom_prenom(mail string) (string, string) {
	mailR := []rune(mail)
	var index_point, index_arobase int

	for i, r := range mailR {
		if r == '.' {
			index_point = i
		} else if r =='@' {
			index_arobase = i
		}
	}

	prenom := string(mailR[0:index_point])
	nom := string(mailR[index_point+1:index_arobase])
	return nom, prenom
}

func main() {
	mail := input("Entrez votre adresse e-mail universitaire: ")
	nom, prenom := extraire_nom_prenom(mail)
	fmt.Printf("Bonjour %v %v !", prenom, nom)
}

