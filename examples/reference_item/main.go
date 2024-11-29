package main

import (
	"fmt"
	grug "github.com/lavafroth/grugiter/grug"
)

// This iterator will stop after 5 iterations
// but yields an *int.
type Limited struct {
	Index int
}

// Always return a pointer to your object.
// In this case, your object is a *int.
// Therefore, return
// a pointer to *int = **int
func (l *Limited) Next() **int {
	index := &l.Index
	l.Index += 1
	if l.Index > 5 {
		return nil
	}
	return &index
}

func main() {
	lim := Limited{0}
	limIterator := grug.NewIterator(&lim)
	for someOrNil := limIterator.Next(); someOrNil != nil; someOrNil = limIterator.Next() {
		elementPtr := *someOrNil
		element := *elementPtr
		fmt.Printf("%d\n", element)
	}
}
