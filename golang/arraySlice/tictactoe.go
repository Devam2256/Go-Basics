package main

import (
	"fmt"
	"strings"
)

func tictactoe() {
	board := [][]string {
		{"-", "-", "-"}, // it is okay not to mention "[]string" just because mentioned that board is slice of slices of strings  .     
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}