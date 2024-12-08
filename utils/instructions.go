package utils

import "fmt"

type list struct {
	a stack
	b stack
}

func (l *list) Print() {
	fmt.Println(l.a)
	fmt.Println(l.b)
}

func NewList(slice []int) list {
	return list{NewStack(slice), NewStack(nil)}
}

func (l *list) Sa() {
	fmt.Println("sa")
	l.a.Swap()
}
func (l *list) Sb() {
	fmt.Println("sb")
	l.b.Swap()
}
func (l *list) Pa() {
	fmt.Println("pa")
	l.a.Push(l.b.Pop())
}
func (l *list) Pb() {
	fmt.Println("pb")
	l.b.Push(l.a.Pop())
}

func (l *list) Ra() {
	fmt.Println("ra")
	l.a.Rotate()
}
func (l *list) Rb() {
	fmt.Println("rb")
	l.b.Rotate()
}

func (l *list) Rr() {
	fmt.Println("rr")
	l.Ra()
	l.Rb()
}

func (l *list) Rra() {
	fmt.Println("rra")
	l.a.ReverseRotate()
}
func (l *list) Rrb() {
	fmt.Println("rrb")
	l.b.ReverseRotate()
}

func (l *list) Rrr() {
	fmt.Println("rrr")
	l.a.ReverseRotate()
	l.b.ReverseRotate()
}
