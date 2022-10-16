package main

import "fmt"

type MoveDirection struct {
	direction   string
	multiplierX int
	multiplierY int
}

type Board struct {
	height int
	width  int

	playerLocationOffset int
	moveOffset           int

	identifier map[string]rune

	moveDirections map[string]MoveDirection
}

type BoardParameters struct {
	boardSettings Board
	boardInput    [8][8]byte
	boardResult   *[8][8]byte
	player        rune
}

func newBoard() *Board {
	var up = MoveDirection{"up", 0, -1}
	var down = MoveDirection{"down", 0, 1}
	var left = MoveDirection{"left", -1, 0}
	var right = MoveDirection{"right", 1, 0}

	moveDirectionMap := make(map[string]MoveDirection)
	moveDirectionMap["up"] = up
	moveDirectionMap["down"] = down
	moveDirectionMap["left"] = left
	moveDirectionMap["right"] = right

	identifierMap := make(map[string]rune)
	identifierMap["playerA"] = 'B'
	identifierMap["playerB"] = 'W'
	identifierMap["empty"] = '.'
	identifierMap["move"] = '0'

	return &Board{8, 8, 1, 2, identifierMap, moveDirectionMap}
}

func markPlayersNextMoves(boardInput *[8][8]byte, player rune) [8][8]byte {
	var resultBoard = *boardInput
	var boardSettings = newBoard()
	var otherPlayer = boardSettings.identifier["playerA"]
	var boardPositionY, boardPositionX int

	if player == boardSettings.identifier["playerA"] {
		otherPlayer = boardSettings.identifier["playerB"]
	}

	var boardParameters = BoardParameters{boardSettings: *boardSettings, boardInput: *boardInput, boardResult: &resultBoard}

	for boardPositionY = 0; boardPositionY < boardSettings.height; boardPositionY++ {
		for boardPositionX = 0; boardPositionX < boardSettings.width; boardPositionX++ {
			if boardInput[boardPositionY][boardPositionX] == byte(player) {
				for _, direction := range boardSettings.moveDirections {
					checkOpponentAndMove(boardParameters, boardPositionX, boardPositionY, direction, otherPlayer)
				}
			}
		}
	}

	return resultBoard
}

func checkOpponentAndMove(boardParameters BoardParameters, playerX int, playerY int, direction MoveDirection, otherPlayer rune) {
	var boardInput = boardParameters.boardInput
	var boardSettings = boardParameters.boardSettings

	if boardInput[playerY+(boardSettings.playerLocationOffset*direction.multiplierY)][playerX+(boardSettings.playerLocationOffset*direction.multiplierX)] == byte(otherPlayer) {
		var offsetY = boardSettings.moveOffset * direction.multiplierY
		var offsetX = boardSettings.moveOffset * direction.multiplierX

		checkMove(boardParameters, playerX, playerY, offsetY, offsetX)
	}
}

func checkMove(boardParameters BoardParameters, playerX int, playerY int, moveOffsetY int, moveOffsetX int) {
	var boardInput = boardParameters.boardInput
	var boardSettings = boardParameters.boardSettings
	var resultBoard = boardParameters.boardResult

	if boardInput[playerY+moveOffsetY][playerX+moveOffsetX] == byte(boardSettings.identifier["empty"]) {
		resultBoard[playerY+moveOffsetY][playerX+moveOffsetX] = byte(boardSettings.identifier["move"])
	}
}

func main() {
	fmt.Println("Hello World!")
}
