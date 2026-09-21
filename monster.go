package main

import "fmt"

type Monster struct {
	Name       string
	MaxHP      int
	HP         int
	Attack     int
	Initiative int
}

func initGoblin() Monster {
	return Monster{
		Name:       "Gobelin d'entrainement",
		MaxHP:      40,
		HP:         40,
		Attack:     5,
		Initiative: 5,
	}
}

// goblinPattern : 100% de l'attaque, 200% tous les 3 tours.
func goblinPattern(m *Monster, c *Character, turn int) {
	dmg := m.Attack
	if turn%3 == 0 {
		dmg *= 2
	}
	c.HP -= dmg
	if c.HP < 0 {
		c.HP = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts\n", m.Name, c.Name, dmg)
	fmt.Printf("%s PV : %d / %d\n", c.Name, c.HP, c.MaxHP)
}

// initNoxar : boss du jeu (antagoniste principal).
// L'initiative et l'attaque sont à ajuster selon l'équilibrage voulu.
func initNoxar() Monster {
	return Monster{
		Name:       "Noxar",
		MaxHP:      120,
		HP:         120,
		Attack:     12,
		Initiative: 8,
	}
}
