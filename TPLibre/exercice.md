# Exercice compression d'image

## Reflexion

Le but de l'exercice est de créer un algo qui transforme une image représentée par chaîne de 'W' et de 'B' comme "WWWWBBBBWWWWWWWW" en une version compressée "4W4B8W".

Comme l'énoncé demande que l'utilisateur saisisse lui même la chaîne, il va falloir faire un **sous-programme de saisie**. Ce sous programme devra vérifier que tous les caractères de la chaîne sont soit 'B', soit 'W'.  

Cas extrême: que faire si la chaîne saisie est vide? Redemander la saisie, donc à traiter dans le sous-programme de saisie.

Dans tous les autres cas, une fois la saisie valide, il faut un **sous-programme pour compresser**. 

Réfléchissons sur un exemple simple: WWBB devrait donc donner 2W2B. On rappelle que les chaines sont en fait des tranches de caractères (enfin d'octet, mais on peut les convertir en runes donc pas de pb). Comme toujours avec les tranches, il va falloir parcourir la tranche. Le premier caractère est 'W', puis 'W' puis 'B'. Il faudrait donc garder en mémoire le 1er caractère ('W') et parcourir la tranche **jusqu'a** tomber sur un autre caractère ('B'), en faisant augmenter un compteur. Quand on tombe sur un autre caractère, le compteur contient exactement le nombre 'W', ici 2. On peut concaténer le compteur (2) et le 1er caractère qu'on avait gardé ('W') au résultat (qu'on aurait donc prélablement intialisé comme une chaîne vide). Puis on ramène le compteur à 0 et on garde en mémoire le premier caractère suivant (ici 'B'), et on recommence **jusqu'a** être au bout de la chaîne.

## Algorithme

Ok c'est l'idée mais comment on le code. Ici je vois que je vais avoir besoin de **2 boucles** car j'ai utilisé 2 fois "jusqu'a". La premiere boucle ça sera un "jusqu'a être au bout de la chaine" et à l'intérieur je vais faire une autre boucle "jusqu'a tomber sur un caractère différent du 1er". 

Dans la première je vais utiliser un indice i qui vaudra 0 au début et la condition de boucle sera i\<len(image). Avant de lancer la seconde, je dois garder dans une variable le caratère image[i]Dans la seconde, j'aurais un autre indice, compteur, qui va de 0 jusqu'a la recontre d'un caractere différent. Les deux indices augmentent a chaque étape de la seconde boucle.

Etapes:
- CONVERTIR image EN TYPE []rune
- DECLARER resultat (string), i, compte (int), premier_caractere (rune)
- AFFECTER "" A resultat, 0 A i, 0 a compte
- TANT QUE i \< len(image)
    - AFFECTER image[i] A premier_caractere
    - TANT QUE image[i]==premier_caractere 
        - INCREMENTER i 
        - INCREMENTER compte
    - FIN TANT QUE
    - CONVERTIR compte ET premier_caracter EN TYPE string
    - AFFECTER resultat + compte + premier_caracter
    - AFFECTER 0 A compte
- FIN TANT QUE
- RENVOYER resultat

Il y a un dernier problème, dans le second tant que, rien n'empeche que i dépasse la longueur de la chaîne, il faut donc rajouter la condition i\< len(image) dans le second aussi.