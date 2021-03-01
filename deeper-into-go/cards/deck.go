package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck type
type deck []string

const wrongFilename int = -201

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", cardValue, cardSuit))
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	fmt.Println("@deck [")
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("]")
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(wrongFilename)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	size := len(d)
	for i := range d {
		newIndex := r.Intn(size - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
