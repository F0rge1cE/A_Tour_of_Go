package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if (t.Left != nil){
	Walk(t.Left, ch);
	}
	
	ch<-t.Value
	
	if (t.Right != nil){
	Walk(t.Right,ch);
	}
	return;
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	
	cht1 := make(chan int, 10)
	cht2 := make(chan int, 10)
	go Walk(t1, cht1)
	go Walk(t2, cht2)
	
	for i:=0; i<10; i++ {
		if (<-cht1 != <- cht2){
			return false
		}
	}
	
	/* // Another method, need to add close() in Walk()
	for v := range c1 {
		if v != <-c2 {
			return false
		}
	}
	*/
	
	return true
}

func main() {
	ch := make(chan int, 10)
	
	go Walk(tree.New(1), ch)
	
	for i:=0; i<10; i++ {
		fmt.Println(<-ch)
	}
	
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	
}

