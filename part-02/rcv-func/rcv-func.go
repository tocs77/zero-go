//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name      string
	health    int
	maxHealth int
	energy    int
	maxEnergy int
}

func (p *Player) HealthUpdate(value int) {

	fmt.Println(p.name, "update", value, "health -> ", p.health)
	p.health += value
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
	if p.health < 0 {
		p.health = 0
	}
}

func (p *Player) EnergyUpdate(value int) {
	fmt.Println(p.name, "update", value, "energy -> ", p.energy)
	p.energy += value
	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}
	if p.energy < 0 {
		p.energy = 0
	}
}

func main() {

	player := Player{
		name:      "knight",
		health:    50,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 80,
	}

	fmt.Println(player)
	player.HealthUpdate(10)
	player.EnergyUpdate(5)
	fmt.Println(player)

}
