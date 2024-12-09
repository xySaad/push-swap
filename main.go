package main

import (
	"push-swap/utils"
)

func main() {
	input := []int{2, 1, 3, 6, 8, 5}
	list := utils.NewList(input)
	list.Print()
	list.Rrr()
	list.Print()
}
