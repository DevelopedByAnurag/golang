package main

import "fmt"

//create a new type 'deck'
//which is slice of string

type deck []string

func (d deck) print() {
	for i , card := range d{
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int){

}