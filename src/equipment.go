package main

import "fmt"

const (
	itemHat   = "Chapeau de l'aventurier"
	itemTunic = "Tunique de l'aventurier"
	itemBoots = "Bottes de l'aventurier"
)


const (
	itemCap      = "Casquette étoilée"
	itemJacket   = "Veste des Skylands"
	itemSneakers = "Baskets à étoile"
)


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
	old := *slot

	if !removeInventory(c, item, 1) {
		fmt.Println("Vous ne possédez pas cet objet.")
		return
	}

	if old != "" {
		if !addInventory(c, old) {
			// Pas de place pour reposer l'ancien équipement : on annule l'échange.
			c.Inventory = append(c.Inventory, item)
			fmt.Println("Inventaire plein, impossible d'échanger l'équipement.")
			return
		}
		oldInfo := equipData[old]
		c.MaxHealth -= oldInfo.HP
		c.Initiative -= oldInfo.Initiative
		fmt.Printf("%s retourne dans l'inventaire.\n", old)
	}

	*slot = item
	c.MaxHealth += info.HP
	c.Initiative += info.Initiative
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
	fmt.Printf("Équipé : %s (PV max : %d, initiative : %d)\n", item, c.MaxHealth, c.Initiative)
}


func useItem(c *Character, item string) {
	if _, ok := equipData[item]; ok {
		equip(c, item)
		return
	}
	switch item {
	case "Potion de vie":
		takePot(c)
	case "Potion de poison":
		poisonPot(c)
	case itemUpg:
		if upgradeInventorySlot(c) {
			removeInventory(c, item, 1)
		}
	case "Livre de Sort : Star Shot":
		before := len(c.Skills)
		spellBook(&c.Skills)
		if len(c.Skills) > before { 
			removeInventory(c, item, 1)
		}
	default:
		fmt.Println("Cet objet n'est pas encore utilisable :", item)
	}
}
