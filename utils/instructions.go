package utils

import "fmt"

type List struct {
	a stack
	b stack
}

func (l *List) Print() {
	fmt.Println(l.a)
	fmt.Println(l.b)
}

func NewList(slice []int) List {
	return List{NewStack(slice), NewStack(nil)}
}

func (l *List) Sa() {
	fmt.Println("sa")
	l.a.Swap()
}
func (l *List) Sb() {
	fmt.Println("sb")
	l.b.Swap()
}
func (l *List) Pa() {
	fmt.Println("pa")
	l.a.Push(l.b.Pop())
}
func (l *List) Pb() {
	fmt.Println("pb")
	l.b.Push(l.a.Pop())
}

func (l *List) Ra() {
	fmt.Println("ra")
	l.a.Rotate()
}
func (l *List) Rb() {
	fmt.Println("rb")
	l.b.Rotate()
}

func (l *List) Rr() {
	fmt.Println("rr")
	l.Ra()
	l.Rb()
}

func (l *List) Rra() {
	fmt.Println("rra")
	l.a.ReverseRotate()
}
func (l *List) Rrb() {
	fmt.Println("rrb")
	l.b.ReverseRotate()
}

func (l *List) Rrr() {
	fmt.Println("rrr")
	l.a.ReverseRotate()
	l.b.ReverseRotate()
}
