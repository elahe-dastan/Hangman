package main

import (
	"github.com/elahe-dastan/hangman/bank"
	"github.com/elahe-dastan/hangman/handler"
)

func main() {
	b := bank.New()
	selected := b.Choose()

	p := handler.New(selected)
	p.Setup()

	for {
		if p.Show() {
			break
		}

		p.Play()
	}
}