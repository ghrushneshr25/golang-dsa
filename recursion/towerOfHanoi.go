package recursion

import (
	"fmt"
)

type TowerOfHanoi struct {
	MoveCount     int
	NumberOfDisks int
	Move          []string
}

func (toh *TowerOfHanoi) Recursion(current int, from, to, aux string) {
	if current == 1 {
		toh.MoveCount++
		toh.Move = append(toh.Move, fmt.Sprintf("%v -> %v", from, to)) // Only 1 disk move directly
		return
	}

	toh.Recursion(current-1, from, aux, to)
	toh.MoveCount++
	toh.Move = append(toh.Move, fmt.Sprintf("%v -> %v", from, to))
	toh.Recursion(current-1, aux, to, from)
}

func NewTowerOfHanoi(numberOfDisks int) *TowerOfHanoi {
	return &TowerOfHanoi{
		NumberOfDisks: numberOfDisks,
		Move:          []string{},
	}
}

func ExecuteTowerOfHanoi(toh *TowerOfHanoi) {
	toh.Recursion(toh.NumberOfDisks, "A", "C", "B")
}
