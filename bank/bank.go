package bank

import (
	"math/rand"
	"time"
)

type Bank struct {
	words []string
}

func New() Bank {
	return Bank{words: []string{"parham", "raha", "amirkabir", "alibaba"}}
}

// nolint: gosec
func (b Bank) Choose() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int() % len(b.words)

	return b.words[r]
}