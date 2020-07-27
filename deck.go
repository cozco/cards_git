package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck
//which is slice of strings

type deck []string
type csuits struct {
	gameSuit  []string
	gameValue []string
}

//Creates the deck
func newDeck() deck {

	workkingDeck := deck{}

	cards := csuits{
		gameSuit:  []string{"Spades", "Hearts", "Clubs", "Diamonds"},
		gameValue: []string{"Ace", "Two", "Three", "Four"},
	}

	for suit := range cards.gameSuit {
		for value := range cards.gameValue {
			workkingDeck = append(workkingDeck, cards.gameValue[value]+" of "+cards.gameSuit[suit])
		}
	}
	/*cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}*/

	/*for suit := range cardSuits {
		for value := range cardValues {
			cards = append(cards, cardValues[value]+" of "+cardSuits[suit])
		}
	}*/

	return workkingDeck
}

//Prints out the Deck
func (d deck) PrintDeck() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Deals out a deck hand
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

//To String
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Save a list of decks to a file

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Read a file from hard drive

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		//Option #1  - log error and return call to newDeck()
		//Option #2  -  log error and quit
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

//Shuffle a deck

func (d deck) shuffle() {
	//for i := range d {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	//newPosition := rand.Intn(len(d) - 1)
	//d[i], d[newPosition] = d[newPosition], d[i]
	//}
}
