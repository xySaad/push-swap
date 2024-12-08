package utils

import "fmt"

type List struct {
	A stack
	B stack
}

func (l *List) Print() {
	fmt.Println(l.A)
	fmt.Println(l.B)
}

func NewList(slice []int) List {
	return List{NewStack(slice), NewStack(nil)}
}

func (l *List) Sa() {
	fmt.Println("sa")
	l.A.Swap()
}
func (l *List) Sb() {
	fmt.Println("sb")
	l.B.Swap()
}
func (l *List) Pa() {
	fmt.Println("pa")
	l.A.Push(l.B.Pop())
}
func (l *List) Pb() {
	fmt.Println("pb")
	l.B.Push(l.A.Pop())
}

func (l *List) Ra() {
	fmt.Println("ra")
	l.A.Rotate()
}
func (l *List) Rb() {
	fmt.Println("rb")
	l.B.Rotate()
}

func (l *List) Rr() {
	fmt.Println("rr")
	l.Ra()
	l.Rb()
}

func (l *List) Rra() {
	fmt.Println("rra")
	l.A.ReverseRotate()
}
func (l *List) Rrb() {
	fmt.Println("rrb")
	l.B.ReverseRotate()
}

func (l *List) Rrr() {
	fmt.Println("rrr")
	l.A.ReverseRotate()
	l.B.ReverseRotate()
}
