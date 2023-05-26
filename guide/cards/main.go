package main

import "fmt"

//func main() {
//	//var card string = "Ace of Spades"
//	//card := "Ace of Spades"
//	//card = "Five of Diamonds"
//	card := newCard()
//	fmt.Println(card)
//}

func main() {
	// slice
	cards := []string{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
