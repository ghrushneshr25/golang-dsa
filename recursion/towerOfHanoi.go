/*
Problem-1 Discuss Towers of Hanoi puzzle.
  - Only one disk can be moved at a time.
  - Each move consists of taking the upper disk from one of the stacks and placing it on top of another stack i.e. a disk can only be moved if it is the uppermost disk on a stack.
  - No disk may be placed on top of a smaller disk.

It comprises three main recursive steps:
 1. Shift disks from the source rod to an auxiliary rod.
 2. Move the remaining disk to the target rod.
 3. Shift the disks from the auxiliary rod to the target rod.
*/
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

func (toh *TowerOfHanoi) Execute() {
	toh.Recursion(toh.NumberOfDisks, "A", "C", "B")
}
