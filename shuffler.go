package main

import (
	"math/rand"
	"time"
)

// Shuffler shuffle cards in a random order
func Shuffler(d *Deck) {
	neworder := make([]Card, len(d.Cards))
	rand.Seed(int64(time.Now().Nanosecond()))
	m := make(map[int]bool)
	oid := 0
	for {
		id := rand.Intn(len(d.Cards))
		if _, ok := m[id]; !ok {
			m[id] = true
			neworder[oid] = d.Cards[id]
			oid++
		}
		if len(m) == len(d.Cards) {
			break
		}
	}
	d.Cards = neworder
}
