package main

import "fmt"

func answer(board *[8][8]byte, player rune) [8][8]byte {
	var otherPlayer = 'B'
	var empty = '.'
	var move = '0'
	var y, x int

	if player == 'B' {
		otherPlayer = 'W'
	}

	var resultBoard = *board

	for y = 0; y < 8; y++ {
		for x = 0; x < 8; x++ {
			if board[y][x] == byte(player) {
				fmt.Printf("[%x]", board[y][x])
				if board[y+1][x] == byte(otherPlayer) {
					if board[y+2][x] == byte(empty) {
						resultBoard[y+2][x] = byte(move)
					}
				}
				if board[y-1][x] == byte(otherPlayer) {
					if board[y-2][x] == byte(empty) {
						resultBoard[y-2][x] = byte(move)
					}
				}
				if board[y][x+1] == byte(otherPlayer) {
					if board[y][x+2] == byte(empty) {
						resultBoard[y][x+2] = byte(move)
					}
				}
				if board[y][x-1] == byte(otherPlayer) {
					if board[y][x-2] == byte(empty) {
						resultBoard[y][x-2] = byte(move)
					}
				}
			}
		}
		fmt.Printf("\n")
	}

	return resultBoard
}

func main() {
	fmt.Println("Hello World!")
}
