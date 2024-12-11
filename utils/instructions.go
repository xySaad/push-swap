package utils

import (
	"fmt"
)

type List struct {
	A            stack
	B            stack
	Instructions []string
}

func (l *List) Print() {
	fmt.Print("[")
	for i := len(l.A.slice) - 1; i >= 0; i-- {
		fmt.Print(l.A.slice[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
	fmt.Print("[")
	for i := len(l.B.slice) - 1; i >= 0; i-- {
		fmt.Print(l.B.slice[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println("]")
}

func (l *List) PrintR() {
	fmt.Println(l.A.slice)
	fmt.Println(l.B.slice)
}

func NewList(slice []int) List {
	return List{NewStack(slice, true), NewStack(nil, true), []string{}}
}
func NewUntrackedList(slice []int) List {
	return List{NewStack(slice, false), NewStack(nil, false), []string{}}
}

func (l *List) Sa() {
	l.Instructions = append(l.Instructions, "sa")
	l.A.swap()
}
func (l *List) Sb() {
	l.Instructions = append(l.Instructions, "sb")
	l.B.swap()
}
func (l *List) Pa() {
	l.Instructions = append(l.Instructions, "pa")
	l.A.push(l.B.pop())
}
func (l *List) Pb() {
	l.Instructions = append(l.Instructions, "pb")
	l.B.push(l.A.pop())
}

func (l *List) Ra() {
	l.Instructions = append(l.Instructions, "ra")
	l.A.rotate()
}
func (l *List) Rb() {
	l.Instructions = append(l.Instructions, "rb")
	l.B.rotate()
}

func (l *List) Rr() {
	l.Instructions = append(l.Instructions, "rr")
	l.A.rotate()
	l.B.rotate()
}

func (l *List) Rra() {
	l.Instructions = append(l.Instructions, "rra")
	l.A.reverseRotate()
}
func (l *List) Rrb() {
	l.Instructions = append(l.Instructions, "rrb")
	l.B.reverseRotate()
}

func (l *List) Rrr() {
	l.Instructions = append(l.Instructions, "rrr")
	l.A.reverseRotate()
	l.B.reverseRotate()
}

func (l *List) Ss() {
	l.Instructions = append(l.Instructions, "ss")
	l.A.swap()
	l.B.swap()
}

func (l *List) Do(operation string) {
	switch operation {
	case "ra":
		l.Ra()
	case "pb":
		l.Pb()
	case "pa":
		l.Pa()
	}
}
