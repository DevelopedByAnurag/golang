package main

func main()  {
	cards := newDeck()
	//hand, remaining := deal(cards, 5)
	//
	//hand.print()
	//remaining.print()
	//cards.saveToFile("listcards")
	cards.shuffle()
	cards.print()
}

