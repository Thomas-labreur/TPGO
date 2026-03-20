**Données**: chaine (string)

**Sortie**: la même chaîne avec des '?' à la place des ' ' (string)

**Etapes**:
- DECLARER chaineR ET resultat DE TYPE []rune
- AFFECTER []rune(chaine) A chaineR
- POUR i DE 0 A len(chaineR)-1
    - SI chaineR[i] == ' ' ALORS AFFECTER '?' A resultat[i]
    - SINON AFFECTER chaineR[i] A resultat[i]
    - FIN SI
- FIN POUR
- CONVERTIR resultat EN TYPE string
- RENVOYER resultat