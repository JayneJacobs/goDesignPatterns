package prototype

import (
	"fmt"
)

// Musician struct {
type Musician struct {
	Name, Instrument string
	Position         int
}

const (

	// LeadSinger = iota
	LeadSinger = iota
	// GuitarPlayer IOTA
	GuitarPlayer
	// BassPlayer IOTA
	BassPlayer
	// Drummer IOTA
	Drummer
	// KeyboardPlayer IOTA
	KeyboardPlayer
)

// NewMusician (role int) *Musician
func NewMusician(role int) *Musician {
	switch role {
	case GuitarPlayer:
		return &Musician{"", "Guitar Player", 2}
	case LeadSinger:
		return &Musician{"", "Lead Singer", 1}
	case Drummer:
		return &Musician{"", "Drummer", 4}
	case KeyboardPlayer:
		return &Musician{"", "Keyboard Playter", 3}
	default:
		panic("Unsupported role")
	}
}

// Mainlike () {
func Mainlike() {
	m := NewMusician(GuitarPlayer)
	m.Name = "Jim"
	fmt.Println(m)
}
