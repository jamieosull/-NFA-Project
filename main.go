// Adapted from https://swtch.com/~rsc/regexp/regexp1.html
// Adapted from https://github.com/ianmcloughlin

package main

import (
	"fmt"
	"bufio"
)

// Setting up a struct that hold type rune & pointers to two edges
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}


//As the compiler scans the postfix expression, it maintains a stack of computed NFA fragments.
// Literals push new NFA fragments onto the stack, while operators pop fragments off the stack and then push a new fragment.
// For example, after compiling the abb in abb.+.a., the stack contains NFA fragments for a, b, and b. The compilation of the . 
// that follows pops the two b NFA fragment from the stack and pushes an NFA fragment for the concatenation bb..
// Each NFA fragment is defined by its start state and its outgoing arrows:
type nfa struct { 
	start *state 
	out  *state 
}





func intopost(infix string) string {
	
	//Creates a map with special characters and maps them to integers
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	//Arrray of runes
	pofix := []rune{} //rune-character on the screen diplayed in UTF-8

	//Stack stores operators from the infix regular expression
	s := []rune{}
	
	//Loop over the input
	//range(convert into array of runes using UTF)
	for _, r := range infix {
		switch {

		case r == '(':
			//puts open bracket at the end of the stack
			s = append(s, r)
		case r == ')':
			//Pop of the stack until an opening bracket is found
			//len(s)-1 = the last element on the stack
			for s[len(s)-1] != '(' {
				pofix = append(pofix, s[len(s)-1]) //last element of 's'
				s = s[:len(s)-1]                   //everything in 's' except the last character
			}

			s = s[:len(s)-1]
		//if a special character is found
		case specials[r] > 0:
			//while the stack still has an element on it and the precedence of the current character that reads is <= the precedence of the top element of the stack
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				//Takes character of the stack and sticks into pofix
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			
			s = append(s, r)
		
		default:
			//Takes the default characters and sticks it at the end of pofix
			pofix = append(pofix, r)
		}
	}

	//If there is anything left on the stack, append to pofix
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}


	return string(pofix)
}

func main() {
	
	//scanner for user input for regular expresion
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter you regular expresion")
	scanner.Scan()
	
	//user input to compare the string and regular expresion
	//infix := scanner.Text()
	fmt.Println("Enter the string you want to check against your first Input")
	scanner.Scan()
	
	// Takes in input to check the string
	//string := scanner.Text()
	

	
	

	//Answer: ab.c*.
	//fmt.Println("Infix:  ", "a.b.c*")
	//fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer: abd|.*
	//fmt.Println("Infix:  ", "(a.(b|d))*")
	//fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//Answer: abd|.c*.
	//fmt.Println("Infix:  ", "a.(b|d).c*")
	//fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	//fmt.Println("Infix:  ", "a.(b.b)+.c")
	//fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))

}
