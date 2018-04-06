// Adapted from https://swtch.com/~rsc/regexp/regexp1.html
// Adapted from https://github.com/ianmcloughlin

package main

import (
	"fmt"
	"bufio"
	"os"
)

// Setting up a struct that hold type rune & pointers to two edges
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}


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

// Helper function, searches through all states
func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)
	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
}

//match regular expresion with a string
func poMatch(postfix string, s string) bool {
	//isMatch is false(default)
	isMatch := false
	//postFixtoNfa func to get the nfa of input
	poNfa := postfixToNfa(postfix)
	//Set current and next nfa pointer arrays
	current := []*state{}
	next := []*state{}
	
	// adds current to function addState
	current = addState(current[:], poNfa.initial, poNfa.out)
	
	//loops through each rune 
	for _, r := range s {
		//loop through each nfa pointer in current
		for _, c := range current {
			//if symbol of the nfa point equals the rune
			// then next will equal to the state with input
			if c.symbol == r {
				next = addState(next[:], c.edge1, poNfa.out)
			}
		}
		// set current to next and when next equals to pointer array
		current, next = next, []*state{}
	}
	// Loop through nfa pointer in current
	for _, c := range current {
		// if the current state is equal to the accept state of the nfa is match is set to true
		if c == poNfa.out {
			isMatch = true
			break
		}
	}

	return isMatch
}

// Used to create a nfa
func postfixToNfa(postfix string) *nfa {
	//stack of nfa pointers
	nfaStack := []*nfa{}
	// Loop through each rune
	for _, r := range postfix {
		//switch statement for each rune that handles
		switch r {

		case '.':
			//pops two fragments off of the stack concatentates the two and addes it back onto the stack
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.out.edge1 = frag2.start
			nfaStack = append(nfaStack, &nfa{start: frag1.start, out: frag2.out})

		case '|':
			//pops two fragments off the stack creates a new start state from frag1 & start states
			// create a blank state and adds a point to new values to the stack
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			start := state{edge1: frag1.start, edge2: frag2.start}
			out := state{}
			frag1.out.edge1 = &out
			frag2.out.edge1 = &out

			nfaStack = append(nfaStack, &nfa{start: &start, out: &out})

		case '*':
			// takes one fragment off the stack then creates a empty out state,
			// create a new start state, then state and edge been included to the new out state
			// appends the pointer to nfa with new values
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			out := state{}
			start := state{edge1: frag.start, edge2: &out}
			frag.out.edge1 = frag.start
			frag.out.edge2 = &out

			nfaStack = append(nfaStack, &nfa{start: &start, out: &out})
		default:
			//create empty out state
			//creates a new start state wiht symbol = the rune
			//and edge 1 = the new out state then appends a pointer to
			//a new nfa with new values
			out := state{}
			start := state{symbol: r, edge1: &out}
			nfaStack = append(nfaStack, &nfa{start: &start, out: &out})
		}
	}



	return nfaStack[0]
}

func main() {
	
	//scanner for user input for regular expresion
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter you regular expresion")// Infix notation
	scanner.Scan()
	
	//user input to compare the string and regular expresion
	infix := scanner.Text()
	fmt.Println("Enter the string you want to check against your first Input")
	scanner.Scan()
	
	// Takes in input to check the string
	string := scanner.Text()
	
	//converts the infix notation to postfix notation
	postfix := intopost(infix)
	
	fmt.Println(poMatch(strings.ToLower(postfix), strings.ToLower(checkString)))

	
	
	

	
	

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
