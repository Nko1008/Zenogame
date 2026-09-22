package main

import "fmt"

// Équipements du sujet
const (
	itemHat   = "Chapeau de l'aventurier"
	itemTunic = "Tunique de l'aventurier"
	itemBoots = "Bottes de l'aventurier"
)

// Set de Zeno
const (
	itemCap      = "Casquette étoilée"
	itemJacket   = "Veste des Skylands"
	itemSneakers = "Baskets à étoile"
)

// Set de Noxar
const (
	itemCrown     = "Couronne du Néant"
	itemCloak     = "Manteau d'ombre"
	itemVoidBoots = "Bottes de l'Éclipse"
)

const (
	slotHead  = "tête"
	slotTorso = "torse"
	slotFeet  = "pieds"
)

type equipInfo struct {
	Slot       string
	HP         int // bonus de PV maximum
	Initiative int // bonus d'initiative
}

var equipData = map[string]equipInfo{
	itemHat:   {slotHead, 10, 0},
	itemTunic: {slotTorso, 25, 0},
	itemBoots: {slotFeet, 15, 0},

	itemCap:      {slotHead, 8, 2},
	itemJacket:   {slotTorso, 20, 1},
	itemSneakers: {slotFeet, 10, 3},

	itemCrown:     {slotHead, 20, 0},
	itemCloak:     {slotTorso, 35, 0},
	itemVoidBoots: {slotFeet, 20, 0},
}

func slotPtr(c *Character, slot string) *string {
	switch slot {
	case slotHead:
		return &c.Equipment.Head
	case slotTorso:
		return &c.Equipment.Torso
	default:
		return &c.Equipment.Feet
	}
}

func equip(c *Character, item string) {
	info, ok := equipData[item]
	if !ok {
		return
	}
	slot := slotPtr(c, info.Slot)
	removeInventory(c, item, 1) // libère une place avant de rendre l'ancien équipement
	if *slot != "" {
		old := equipData[*slot]
		c.MaxHP -= old.HP
		c.Initiative -= old.Initiative
		c.Inventory = append(c.Inventory, *slot)
		fmt.Printf("%s retourne dans l'inventaire.\n", *slot)
	}
	*slot = item
	c.MaxHP += info.HP
	c.Initiative += info.Initiative
	if c.HP > c.MaxHP {
		c.HP = c.MaxHP
	}
	fmt.Printf("Équipé : %s (PV max : %d, initiative : %d)\n", item, c.MaxHP, c.Initiative)
}

// useItem : utilisé depuis l'inventaire et depuis le combat.
func useItem(c *Character, item string) {
	if _, ok := equipData[item]; ok {
		equip(c, item)
		return
	}
	switch item {
	case itemUpg:
		if upgradeInventorySlot(c) {
			removeInventory(c, item, 1)
		}
	case "Livre de Sort : Star Shot":
		before := len(c.Skills)
		spellBook(&c.Skills)        
		if len(c.Skills) > before { // le livre n'est consommé que si le sort est appris
			removeInventory(c, item, 1)
		}
	default:
		
		fmt.Println("Cet objet n'est pas encore utilisable :", item)
	}
}
