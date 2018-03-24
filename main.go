package main

import (
	"fmt"
)

func intopost(infix string) string {
	
	//Creates a map with special characters and maps them to integers
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	//Arrray of runes
	pofix := []rune{} //rune-character on the screen diplayed in UTF-8

	//Stack stores operators from the infix regular expression
	s := []rune{}


	return string(pofix)
}

func main() {

	//Answer: ab.c*.
	fmt.Println("Infix:  ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix:  ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix:  ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix:  ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))

}
