package main
import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"time"
)

// Ce code implémente le jeu du pendu: le jeu choisit un mot secret que
// l'utilisateur doit deviner. Il propose des lettres et perd une vie si
// la lettre n'est pas dans le mot.

// On utilise input pour les saisies
func input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}

// Transforme les majuscules en minuscules
func lower(c rune) rune {
	if c>='A' && c<='Z' {
		return rune(c+32)
	}
	return rune(c)
}

// Demande a l'utilisateur de saisir 1 lettre
func demander_rune() rune {
	str := input("Proposez une lettre: ")
	r := []rune(str)
	for len(r)!=1 || lower(r[0]) < 'a' || lower(r[0]) > 'z' {
		str = input("Saisie invalide. Proposez une lettre: ")
		r = []rune(str)
	}
	return lower(r[0])
}

// Genere aléatoirement un mot dans une liste prédéfinie de ~200 mots
func generer_mot() string {
    WORDS := []string{
	"ordinateur", "programmation", "developpement", "application", "interface",
	"variable", "fonction", "structure", "algorithme", "compilation",
	"execution", "memoire", "processeur", "reseau", "internet",
	"navigateur", "serveur", "client", "donnee", "analyse",
	"modele", "apprentissage", "intelligence", "artificielle", "statistique",
	"probabilite", "distribution", "matrice", "vecteur", "scalaire",
	"iteration", "condition", "boucle", "recursion", "parametre",
	"argument", "retour", "erreur", "exception", "debug",
	"optimisation", "performance", "complexite", "architecture", "systeme",
	"logiciel", "materiel", "clavier", "souris", "ecran",
	"fenetre", "bureau", "fichier", "dossier", "extension",
	"format", "compression", "decompression", "chiffrement", "securite",
	"authentification", "autorisation", "utilisateur", "motdepasse", "session",
	"connexion", "deconnexion", "requete", "reponse", "protocole",
	"http", "https", "tcp", "udp", "adresse",
	"port", "socket", "flux", "tampon", "thread",
	"processus", "concurrence", "parallele", "synchronisation", "verrou",
	"mutex", "canal", "goroutine", "planification", "priorite",
	"temps", "horloge", "date", "duree", "timestamp",
	"conversion", "entier", "flottant", "chaine", "rune",
	"caractere", "alphabet", "mot", "phrase", "texte",
	"langue", "francais", "anglais", "espagnol", "allemand",
	"italien", "portugais", "culture", "histoire", "geographie",
	"pays", "ville", "village", "route", "chemin",
	"montagne", "riviere", "lac", "mer", "ocean",
	"plage", "foret", "arbre", "fleur", "animal",
	"chien", "chat", "oiseau", "poisson", "cheval",
	"voiture", "train", "avion", "bateau", "velo",
	"transport", "energie", "electricite", "batterie", "moteur",
	"essence", "diesel", "hybride", "solaire", "eolien",
	"temperature", "pression", "vitesse", "acceleration", "force",
	"masse", "volume", "densite", "poids", "mesure",
	"unite", "metre", "kilometre", "seconde", "minute",
	"heure", "jour", "semaine", "mois", "annee",
	"calendrier", "evenement", "festival", "concert", "theatre",
	"cinema", "film", "musique", "chanson", "instrument",
	"guitare", "piano", "violon", "batterie", "artiste",
	"peinture", "dessin", "sculpture", "photographie", "image",
	"couleur", "forme", "ligne", "texture", "design",
	"mode", "vetement", "chemise", "pantalon", "chaussure",
	"chapeau", "manteau", "echarpe", "gant", "sac",
	"bijou", "montre", "lunette", "maison", "appartement",
	"cuisine", "salon", "chambre", "bureau", "jardin",
	"porte", "fenetre", "mur", "plafond", "sol",
	"table", "chaise", "canape", "lit", "armoire",
	"etagere", "lampe", "lumiere", "ombre", "bruit",
	"silence", "voix", "parole", "discussion", "conversation",
}
	// Initialiser la graine (important !)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(WORDS))
	return WORDS[index]
}

// Créée un string de la taille du mot secret avec uniquement des '_'
func creer_indice(mot string) string {
	motR := []rune(mot)
	indiceR := make([]rune, len(motR))
	for i,_ := range(motR) {
		indiceR[i] = '_'
	}
	return string(indiceR)
}

// Met à jour l'indice et le nombre de vies selon la proposition du joueur
func ajuster_indice(mot, indice string, lettre rune, vies int) (string, bool, int) {
	motR := []rune(mot)
	indiceR := []rune(indice)
	victoire := true // dit si la partie est gagnée 
	compte := 0		// compte le nombre de lettres justes découvertes

	// On parcours le mot 
	for i, c := range(motR) {
		// Si on tombe sur la lettre proposée
		if c==lettre {
			indiceR[i] = lettre // on la révèle dans l'indice
			compte++ // on incrémente le compte
		}

		// Si on tombe sur un _ c'est qu'on a pas encore gagné
		if indiceR[i] == '_' {
			victoire = false
		}
	}

	// On affiche un message et on met a jour le nombre de vies
	if compte==0 {
		fmt.Printf("Non il n'y a pas de %c dans le mot secret.\n", lettre)
		vies--
	} else {
		fmt.Printf("Il y a %v %c dans le mot secret.\n", compte, lettre)
	}

	return string(indiceR), victoire, vies
}

func main() {
	mot := generer_mot()
	indice := creer_indice(mot)
	victoire := false
	vies := 9
	var lettre rune

	// Le jeu continue tant qu'on a ni gagné ni perdu
	for !victoire && vies >0 {
		lettre = demander_rune()
		indice, victoire, vies = ajuster_indice(mot, indice, lettre, vies)
		if victoire {
			fmt.Printf("Bravo ! Le mot était %v.\n", mot)
		} else if vies == 0 {
		    fmt.Printf("Perdu ! Le mot était %v.\n", mot)
		} else {
			fmt.Printf("Indice: %v (%v vies restantes)\n", indice, vies)
		}
	}
}