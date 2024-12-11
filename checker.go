package main

import (
	"fmt"
	"push-swap/utils"
)

func checker(input []int, instructions []string) {
	list := utils.NewUntrackedList(input)
	for _, operation := range instructions {
		list.Do(operation)
	}
	fmt.Println(instructions)
	list.Print()
}
