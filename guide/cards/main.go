package main

//func main() {
//	//var card string = "Ace of Spades"
//	//card := "Ace of Spades"
//	//card = "Five of Diamonds"
//	card := newCard()
//	fmt.Println(card)
//}

func main() {
	// slice
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
