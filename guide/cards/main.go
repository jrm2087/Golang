package main

//func main() {
//	//var card string = "Ace of Spades"
//	//card := "Ace of Spades"
//	//card = "Five of Diamonds"
//	card := newCard()
//	fmt.Println(card)
//}

func main() {
	cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//remainingCards.print()

	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards2")
	//cards.print()

	cards.shuffle()
	cards.print()

}
