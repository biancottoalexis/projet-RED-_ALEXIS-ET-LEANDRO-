




<img width="594" height="400" alt="image" src="https://github.com/user-attachments/assets/a80673ec-90ad-4fe1-88a1-8b10dd991a30" />



# Shadow of Destiny

Mini-jeu de rôle en ligne de commande (CLI) développé en Go, dans le cadre du Projet RED (Ymmersion — Aix Ynov Campus).

*Une flamme pour te guider, une ombre pour te perdre...*

Incarnez un aventurier dans les Terres Ombreuses : choisissez parmi 6 classes, gérez votre inventaire et votre or, achetez chez le marchand, fabriquez votre équipement chez le forgeron, entraînez-vous contre un Gobelin, et vivez une histoire à choix qui vous mène à un combat final contre votre propre Ombre.

## Fonctionnalités

- Création de personnage (nom validé, choix de classe avec PV et Initiative différents)
- Inventaire interactif (potions de vie/poison/mana, sorts, équipement)
- Marchand et Forgeron avec système d'économie (or, matériaux, fabrication)
- Équipement portable avec bonus de PV max
- Amélioration de la capacité d'inventaire
- Combat tour par tour au choix de compétence (Coup de poing / Boule de Feu), avec mana et coup critique aléatoire ("Toucher du Destin")
- Initiative déterminant qui attaque en premier selon la classe
- Système d'expérience et de niveaux
- Histoire à choix en 6 chapitres, menant à un combat final contre "Ton Ombre" avec 2 fins possibles
- Musique de fond activable/désactivable
- Interface stylisée (couleurs violet/blanc, écran titre en ASCII art centré, effet machine à écrire)
- Easter egg "Qui sont-ils ?" dans les crédits

## Équipe

- Alexis
- Leandro

## Prérequis

- [Go](https://go.dev/dl/) version 1.21 ou supérieure

## Installation

    git clone https://github.com/biancottoalexis/projet-red-_alexis-et-leandro-.git
    cd projet-red-_alexis-et-leandro-

## Lancement

    go run .
