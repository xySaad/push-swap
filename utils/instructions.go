package utils

import "fmt"

type List struct {
	A stack
	B stack
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
	return List{NewStack(slice), NewStack(nil)}
}

func (l *List) Sa() {
	fmt.Println("sa")
	l.A.swap()
}
func (l *List) Sb() {
	fmt.Println("sb")
	l.B.swap()
}
func (l *List) Pa() {
	fmt.Println("pa")
	l.A.push(l.B.pop())
}
func (l *List) Pb() {
	fmt.Println("pb")
	l.B.push(l.A.pop())
}

func (l *List) Ra() {
	fmt.Println("ra")
	l.A.rotate()
}
func (l *List) Rb() {
	fmt.Println("rb")
	l.B.rotate()
}

func (l *List) Rr() {
	fmt.Println("rr")
	l.A.rotate()
	l.B.rotate()
}

func (l *List) Rra() {
	fmt.Println("rra")
	l.A.reverseRotate()
}
func (l *List) Rrb() {
	fmt.Println("rrb")
	l.B.reverseRotate()
}

func (l *List) Rrr() {
	fmt.Println("rrr")
	l.A.reverseRotate()
	l.B.reverseRotate()
}

func (l *List) Ss() {
	fmt.Println("ss")
	l.A.swap()
	l.B.swap()
}
