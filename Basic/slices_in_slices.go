package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	var x int
	var y int
	// var sign *string
	sign:= "X"

	for i:=0; i <= len(board)*len(board[0]); i++{
		_sign:=&sign
		if i % 2 == 0 {
			*_sign = "X"
		} else {
			*_sign = "O"
		}

		fmt.Println("Coordinate X:")
		fmt.Scanln(&x)

		fmt.Println("Coordinate Y:")
		fmt.Scanln(&y)

		board[x][y] = *_sign

		for j := 0; j < len(board); j++ {
			fmt.Printf("%s\n", strings.Join(board[j], " "))
		}
	}
}