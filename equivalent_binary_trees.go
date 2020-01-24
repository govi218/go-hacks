package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// ch <- 10
	ch <- (*t).Value
	if (*t).Left != nil {
		Walk((*t).Left, ch)
	}
	if (*t).Right != nil {
		Walk((*t).Right, ch)
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	var stack []int
	
	i, j := 0,0
	
	for i < 10 {
		// read channels
		a := <-ch1 
		i++ // a should be read exactly 10 times
		
		b := 0
		if j < 10 {
			b = nonChanResult(ch2) // don't care about b's error
			j++	
		}
		
		found := false
		// first go through stack
		for k := 0; k < len(stack); k++ {
			if a == stack[k] {
				found = true
				break
			}
		}
		
		// found a in stack
		if found == true {
			continue
		}
		
		// if not found in stack, check rest of eles
		for a != b {		
			// if no more bs
			if j == 10 {
				return false
			}
			// push b to stack
			stack = append(stack, b)
			// keep reading till match
			b = nonChanResult(ch2)
			j++
		}
	}
	return true
}

func nonChanResult(ch chan int) int {
	return <- ch
}

func main() {
	k := 1
	
	for i := 1; i < 10; i++ {	
		t1 := tree.New(k)
		t2 := tree.New(k)

		fmt.Println(t1)
		fmt.Println(t2)
		fmt.Println(Same(t1,t2))
	}
}

