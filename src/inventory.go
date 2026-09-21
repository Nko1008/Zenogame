package main

import "fmt"

const (
	itemUpg = "Augmentation d'inventaire"
	maxUpg  = 3
)

func countItem(c *Character, name string) int {
	n := 0
	for _, it := range c.Inventory {
		if it == name {
			n++
		}
	}
	return n
}

func checkInventorySpace(c *Character) bool {
	return len(c.Inventory) < c.InventoryMax
}

func addInventory(c *Character, item string) bool {
	if !checkInventorySpace(c) {
		fmt.Println("Inventaire plein !")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	return true
}

// removeInventory retire n exemplaires de item. Renvoie false s'il n'y en a pas assez.
func removeInventory(c *Character, item string, n int) bool {
	if countItem(c, item) < n {
		return false
	}
	kept := make([]string, 0, len(c.Inventory))
	for _, it := range c.Inventory {
		if it == item && n > 0 {
			n--
			continue
		}
		kept = append(kept, it)
	}
	c.Inventory = kept
	return true
}

func upgradeInventorySlot(c *Character) bool {
	if c.UpgradeCount >= maxUpg {
		fmt.Println("Vous avez déjà utilisé 3 augmentations d'inventaire.")
		return false
	}
	c.InventoryMax += 10
	c.UpgradeCount++
	fmt.Printf("Capacité de l'inventaire : %d (%d/%d améliorations)\n", c.InventoryMax, c.UpgradeCount, maxUpg)
	return true
}

func accessInventory(c *Character) {
	for {
		fmt.Printf("\n--- Inventaire (%d/%d) ---\n", len(c.Inventory), c.InventoryMax)
		for i, it := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, it)
		}
		fmt.Println("0. Retour")
		choice := readChoice("> ")
		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(c.Inventory) {
			fmt.Println("Choix invalide.")
			continue
		}
		useItem(c, c.Inventory[choice-1])
	}
}
