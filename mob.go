package main

import "math"

type Mob struct {
	Name   string `json:"name"`
	Hp     int    `json:"hp"`
	Atk    int    `json:"atk"`
	Def    int    `json:"def"`
	Symbol string `json:"symbol"`
}

func (m *Mob) getDmg(player *Player) int {
	return int(math.Max(0, (math.Ceil(float64(m.Hp)/float64(player.atk-m.Def)) - 1) * float64(m.Atk - player.def)))
}
