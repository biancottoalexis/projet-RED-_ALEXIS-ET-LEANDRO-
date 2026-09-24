# Gestion de projet - Shadow of Destiny

## Équipe
- Alexis Biancotto
- Leandro

## Répartition des tâches
Le travail a été réparti entre les deux membres de l'équipe sur les différentes
tâches du sujet (création du personnage, inventaire, marchand, forgeron,
combat, monstres), avec des points de synchronisation réguliers pour
s'assurer que les fonctions s'articulaient bien ensemble (ex : le combat
dépend à la fois du personnage et du monstre).

## Planning suivi
- Jour 1 : Setup + création du personnage (Tâches 1 à 6)
- Jour 2 : Marchand + combat de base (Tâches 7 à 12)
- Jour 3 : Économie (Tâches 13 à 15)
- Jour 4 : Équipement (Tâches 16 à 18)
- Jour 5 : Monstre + IA (Tâches 19 à 20)
- Jour 6 : Combat complet (Tâches 21 à 22)
- Jour 7 : Missions bonus (Initiative, Expérience, Combat magique, Mana,
  Amélioration du jeu, Qui sont-ils) + thème + histoire à choix
- Jour 8 : Tests, documentation, préparation de la démo, rendu final

## Difficultés rencontrées
- Comprendre l'architecture Go à deux fichiers (main.go à la racine, package
  séparé dans src/) et les règles d'export (majuscule/minuscule) entre les
  deux packages.
- Gérer l'ordre d'affichage d'un inventaire stocké en map[string]int
  (l'ordre d'une map n'est pas garanti en Go), résolu avec une slice
  auxiliaire pour fixer l'ordre d'affichage.
- Coordonner les dépendances entre missions bonus (ex : la mission Mana
  dépend de la mission Combat magique).

## Choix techniques
- Utilisation de pointeurs (`*Character`, `*Monster`) pour que les méthodes
  modifient directement l'état du personnage/monstre plutôt qu'une copie.
- Séparation d'un fichier par fonctionnalité dans `src/` (character.go,
  combat.go, marchand.go, forgeron.go, histoire.go...) pour garder le code
  lisible.