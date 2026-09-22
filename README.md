# Zenogame

Mini-jeu de rôle en ligne de commande (CLI) développé en Go dans le cadre du Projet RED.

Créez votre personnage, explorez un système d'inventaire et d'économie, équipez-vous auprès du forgeron, apprenez des sorts, et affrontez des monstres au tour par tour jusqu'au combat final contre **Noxar**.

## Fonctionnalités

* Création et gestion d'un personnage (classe, niveau, points de vie, inventaire)
* Système de sorts et de mana
* Marchand et forgeron pour acheter et fabriquer de l'équipement
* Système d'expérience et de montée de niveau
* Combat au tour par tour (initiative, attaque, sorts, objets)
* Combat d'entraînement contre un Gobelin
* Combat final contre Noxar

## Prérequis

* [Go](https://go.dev/dl/) version **1.27.1**

Vérifiez votre installation avec :

```bash
go version
```

La version attendue est :

```text
go version go1.27.1
```

## Installation

Clonez le dépôt puis placez-vous dans le dossier du code source :

```bash
git clone https://github.com/Nko1008/Zenogame.git

cd Zenogame/src
```

## Lancement

Depuis le dossier `src` (celui qui contient les fichiers `.go`) :

```bash
go run .
```

Un menu s'affiche dans le terminal ; suivez les instructions à l'écran pour jouer.

## Structure du dépôt

```text
.
├── src/     # Code source du jeu
├── docs/    # Document de gestion de projet
└── README.md
```

## Auteurs

* Magri Amy
* Salas Valérie
* Koraichi Norhann
