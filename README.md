# Zenogame
# Skylandia

Mini-jeu de rôle en ligne de commande (CLI) développé en Go dans le cadre du Projet RED.

Incarnez Zeno, un elfe de niveau 1, dans un système d'inventaire et d'économie, équipez-vous auprès du forgeron, apprenez des sorts, et affrontez des monstres au tour par tour : un combat d'entraînement contre un Gobelin, puis le combat final contre **Noxar**.

## Fonctionnalités

- Personnage (Zeno) avec classe, niveau, points de vie, inventaire
- Système de sorts et de mana
- Marchand et forgeron pour acheter/fabriquer de l'équipement
- Système d'expérience et de montée de niveau
- Combat au tour par tour (initiative, attaque, sorts, objets)
- Combat d'entraînement contre un Gobelin, et combat final contre Noxar

## Prérequis

- [Go](https://go.dev/dl/) installé (version 1.20 ou supérieure recommandée)

Vérifiez votre installation avec :
```
go version
```

## Installation

Clonez le dépôt puis placez-vous dans le dossier du code source :
```
git clone <lien-du-depot>
cd projet-red_Skylandia/src
```

## Lancement

Depuis le dossier `src` (celui qui contient les fichiers `.go`) :
```
go run .
```

Un menu s'affiche dans le terminal ; suivez les instructions à l'écran pour jouer.

## Structure du dépôt

```
.
├── src/     # Code source du jeu
├── docs/    # Document de gestion de projet
└── README.md
```

## Auteurs

- Magri Amy
- Salas Valérie
- Koraichi Norhann
