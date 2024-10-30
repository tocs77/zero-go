package main

import (
	"testing"
)

func TestRcvFunc(t *testing.T) {
	player := Player{
		name:      "knight",
		health:    50,
		maxHealth: 100,
		energy:    50,
		maxEnergy: 80,
	}

	player.HealthUpdate(10)
	player.EnergyUpdate(5)

	if player.health != 60 || player.energy != 55 {
		t.Errorf("Expected health %d and energy %d, got %d and %d", 60, 55, player.health, player.energy)
	}

	player.EnergyUpdate(-5)
	player.HealthUpdate(-10)

	if player.health != 50 || player.energy != 50 {
		t.Errorf("Expected health %d and energy %d, got %d and %d", 50, 50, player.health, player.energy)
	}

	player.EnergyUpdate(1000)
	player.HealthUpdate(1000)

	if player.health != 100 || player.energy != 80 {
		t.Errorf("Expected health %d and energy %d, got %d and %d", 100, 100, player.health, player.energy)
	}

	player.EnergyUpdate(-1000)
	player.HealthUpdate(-1000)

	if player.health != 0 || player.energy != 0 {
		t.Errorf("Expected health %d and energy %d, got %d and %d", 0, 0, player.health, player.energy)
	}
}
