package main
import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    )

///////////////////////
// FONCTIONS ANNEXES //
///////////////////////

type adresse struct{
	rue string
	codePostal int
	ville string
}

type etudiant struct{
	numero int
	nom string
	prenom string
	adresse adresse
	notes []float64
}

func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

func saisirNote() float64 {
	// Saisie
	note_str := input("Saisir une note: ")
	note, err2 := strconv.ParseFloat(note_str, 64)
	for err2 != nil || note <0 || note > 20 {
		note_str = input("Note invalide! Saisir une note: ")
		note, err2 = strconv.ParseFloat(note_str, 64)
	}
	return note
}

func saisirNumEtu(etudiants []etudiant) int {
	// Saisie d'un numero
	num_str := input("Saisir un numero etudiant: ")
	num, err := strconv.Atoi(num_str)
	_, _, err2 := findEtu(num, etudiants)

	// On verifie la conformité du numero
	// Cas d'erreur: pas un numero, numero <0 (sauf -1), numero inexistant (sauf -1)
	for err!=nil || (num !=-1 && num <0) || (err2!=nil && num!=-1) { 
		num_str = input("Numero invalide ou inexistant. Saisir un numero: ")
		num, err = strconv.Atoi(num_str)
		_, _, err2 = findEtu(num, etudiants)
	}

	return num
}

func saisirNbNotes() int {
	// Saisie
	num_str := input("Combien de notes voulez vous entrer: ")
	num, err := strconv.Atoi(num_str)
	for err != nil || num <0 {
		num_str = input("Nombre invalide! Rééssayez: ")
		num, err = strconv.Atoi(num_str)
	}
	return num
}

/////////////////////////
// FONCTIONS DEMANDEES //
/////////////////////////

func inputEtu() etudiant {

	// 1. Saisir le numero
	num_str := input("Numero de l'etudiant: ")
	num, err := strconv.Atoi(num_str)
	for err!=nil || num<0 {
		num_str = input("Numero invalide ! Numero de l'etudiant: ")
		num, err = strconv.Atoi(num_str)
	}

	// 2. Saisir le nom et prenom
	nom := input("Nom de l'étudiant: ")
	prenom := input("Prenom de l'étudiant: ")

	// 3. Saisir une adresse
	fmt.Println("Adresse de l'étudiant:")
	rue := input("Rue: ")
	code_str := input("Code postal: ")
	code, err := strconv.Atoi(code_str)
	for err!=nil || num<0 {
		num_str = input("Code postal invalide ! Code postal: ")
		num, err = strconv.Atoi(num_str)
	}
	ville := input("Ville: ")

	// Créer les structures
	adresse := adresse{rue, code, ville}
	etudiant := etudiant{num, nom, prenom, adresse, []float64{}}

	return etudiant
}

func etuToString(etudiant etudiant) string {

	// Conversion en string des éléments qui n'en sont pas
	numero_str := strconv.Itoa(etudiant.numero)
	code_str := strconv.Itoa(etudiant.adresse.codePostal)
	adresse_str := etudiant.adresse.rue + ", " + code_str + " " + etudiant.adresse.ville

	// Construction de la chaîne 
	resultat := "Numero etudiant: " + numero_str + "\n"
	resultat += "Nom: " + etudiant.nom + "\n"
	resultat += "Prenom: " + etudiant.prenom + "\n"
	resultat += "Adresse : " + adresse_str

	// S'il y a des notes, on les ajoute a la chaîne une a une
	if len(etudiant.notes) > 0 {
		resultat += "\nNotes: "
		for _, note := range etudiant.notes {
			resultat += strconv.FormatFloat(note, 'f', -1, 64) + ", "
		}
		resultat = resultat[:len(resultat)-2] // enlever le dernier ", "
	}
	return resultat
}

func findEtu(numero int, etudiants []etudiant) (int, etudiant, error) {

	// Parcours de la liste et renvoi de l'étudiant
	for i, etu := range etudiants {
		if etu.numero == numero {
			return i, etu, nil
		}
	}
	
	// Si pas trouvé, on renvoie une erreur
	return -1, etudiant{}, fmt.Errorf("étudiant avec le numéro %d introuvable", numero)
}

func ajoutNote(etudiants []etudiant, num int, note float64) error {
	
	// On cherche l'étudiant dans la liste
	i, etu, err := findEtu(num, etudiants)

	// S'il exite on saisit une note
	if err==nil {
		etudiants[i].notes = append(etu.notes, note)
	}

	// Dans tous les cas, on renvoie l'erreur
	return err
}

func moyEtu(etu etudiant) (float64, error) {

	// Si pas de notes, la moyenne est 0 et erreur
	if len(etu.notes) == 0 {
		return 0, fmt.Errorf("L'etudiant %v n'a pas de notes.", etu.numero)

	// Sinon on calcule la moyenne et on la renvoie
	} else {
		somme := 0.0
		for _, note := range etu.notes {
			somme += note
		}
		return somme / float64(len(etu.notes)), nil
	}
}

/////////////////////////
// PROGRAMME PRINCIPAL //
/////////////////////////

func main() {
	// etudiants := []etudiant{
	// 	{1, "Saquet", "Frodon", adresse{"Cul de sac", 13790, "La Compté"}, []float64{}},
	// 	{12, "Alvarez", "Judy", adresse{"Charter Street", 18877, "Night City"}, []float64{}},
	// 	{31, "Holmes", "Sherlock", adresse{"221 Baker Street", 20077, "London"}, []float64{}},
	// 	{27, "Parker", "May", adresse{"20 Ingram Street", 76183, "New York City"}, []float64{}},
	// }

	var (
        nb_notes int
        note float64
	    num int
	)

	// 1. Saisie des étudiants
	fmt.Println("SAISIE DES ETUDIANTS\n")
	etudiants := make([]etudiant, 2)
	for i, _ := range etudiants {
		etudiants[i] = inputEtu()
		fmt.Println("\n")
	}
	
	// 2. Saisie des notes	   
	fmt.Println("NOTATION DES ETUDIANTS")
	fmt.Println("Saisir -1 comme numero etudiant pour stopper la notation.\n")

	// On saisit un numero
    num = saisirNumEtu(etudiants)

	// Tant que ce numero n'est pas -1, on rentre des notes
	for num != -1 {
		nb_notes = saisirNbNotes()
		for i:=0; i<nb_notes; i++ {
			note = saisirNote()
			ajoutNote(etudiants, num, note)
			fmt.Println("Note enregistrée.\n")
		}
		num = saisirNumEtu(etudiants)
	}
	
	// 3. Affichage de la liste des étudiants
	fmt.Println("\nLISTE DES ETUDIANTS\n")
	for _, etu := range etudiants {
		fmt.Println(etuToString(etu))

		// Gestion de la moyenne
		moyenne, err := moyEtu(etu)
		if err == nil {
			fmt.Printf("Moyenne: %.2f\n", moyenne)
		} else {
			fmt.Println(err)
		}
		fmt.Println("\n")
	}
}
