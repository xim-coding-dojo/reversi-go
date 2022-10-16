package main

import "fmt"

type Board struct {
	length  int
	width   int
	playerA rune
	playerB rune
	empty   rune
	move    rune
}

func newBoard() *Board {
	return &Board{8, 8, 'B', 'W', '.', '0'}
}

func markPlayersNextMoves(boardInput *[8][8]byte, player rune) [8][8]byte {
	var resultBoard = *boardInput
	var boardSettings = newBoard()
	var otherPlayer = boardSettings.playerA
	var y, x int

	if player == boardSettings.playerA {
		otherPlayer = boardSettings.playerB
	}

	for y = 0; y < 8; y++ {
		for x = 0; x < 8; x++ {
			if boardInput[y][x] == byte(player) {
				if boardInput[y+1][x] == byte(otherPlayer) {
					if boardInput[y+2][x] == byte(boardSettings.empty) {
						resultBoard[y+2][x] = byte(boardSettings.move)
					}
				}
				if boardInput[y-1][x] == byte(otherPlayer) {
					if boardInput[y-2][x] == byte(boardSettings.empty) {
						resultBoard[y-2][x] = byte(boardSettings.move)
					}
				}
				if boardInput[y][x+1] == byte(otherPlayer) {
					if boardInput[y][x+2] == byte(boardSettings.empty) {
						resultBoard[y][x+2] = byte(boardSettings.move)
					}
				}
				if boardInput[y][x-1] == byte(otherPlayer) {
					if boardInput[y][x-2] == byte(boardSettings.empty) {
						resultBoard[y][x-2] = byte(boardSettings.move)
					}
				}
			}
		}
	}

	return resultBoard
}

func main() {
	fmt.Println("Hello World!")
}
