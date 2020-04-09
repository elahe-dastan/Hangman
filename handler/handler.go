package handler

import (
	"bufio"
	"fmt"
	"os"
)

type Problem struct {
	selected string
	// Key is ASCII of the characters of the selected word and value is an slice containing the number of places
	// of the character
	data   map[int][]int
	result []int
}

func New(word string) Problem {
	return Problem{
		selected: word,
		data:     make(map[int][]int),
		result:   make([]int, len(word)),
	}
}

func (p Problem) Setup() {
	for i, l := range p.selected { // i is index of character and l is the character itself
		var s []int

		if m, ok := p.data[int(l)]; ok { // m is the slice showing positions of character "l" if OK
			s = m
		} else {
			s = make([]int, 0)
		}

		s = append(s, i)
		p.data[int(l)] = s
	}
}

func (p Problem) Show() bool {
	completed := true

	for _, val := range p.result {
		if val == 0 {
			fmt.Print("_")

			completed = false
		} else {
			fmt.Print(string(val))
		}
	}

	fmt.Println()

	return completed
}

func (p Problem) Play() {
	fmt.Print("PLayer: ")

	reader := bufio.NewReader(os.Stdin)
	letter, _, _ := reader.ReadRune()

	if s, ok := p.data[int(letter)]; ok {
		for _, index := range s {
			p.result[index] = int(letter)
		}
	}

	fmt.Println()
}
