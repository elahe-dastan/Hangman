package main

import (
	"Hangman/bank"
	"Hangman/handler"
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