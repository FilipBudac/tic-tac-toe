package main

import (
	"fmt"
)

const BoardSize int = 3 * 3

const (
	Empty  int = 0
	Cross  int = 1
	Nought int = 2
)

const (
	EmptyMark  string = "_"
	CrossMark  string = "X"
	NoughtMark string = "0"
)

const (
	FirstPlayer  int = 0
	SecondPlayer int = 1
)

func translatePosition(pos int) string {
	switch pos {
	case Cross:
		return CrossMark
	case Nought:
		return NoughtMark
	}
	return EmptyMark
}

func inputPosition() int {
	pos := -1
	if _, err := fmt.Scanf("%d\n", &pos); err != nil {
		return pos
	}
	return pos
}

type Board struct {
	positions [BoardSize]int
}

func (board Board) Print() {
	for i, pos := range board.positions {
		if i%3 == 2 {
			fmt.Printf("%s \n", translatePosition(pos))
		} else {
			fmt.Printf("%s ", translatePosition(pos))
		}
	}
}

func (board *Board) Update(input int, mark int) {
	board.positions[input] = mark
}

func (board Board) isWinner(mark int) bool {
	// check rows
	for pos := 0; pos < len(board.positions); pos += 3 {
		if board.positions[pos] == mark && board.positions[pos+1] == mark && board.positions[pos+2] == mark {
			return true
		}
	}
	// check cols
	for pos := 0; pos < 3; pos += 1 {
		if board.positions[pos] == mark && board.positions[pos+3] == mark && board.positions[pos+6] == mark {
			return true
		}
	}

	// check diagonals
	if board.positions[0] == mark && board.positions[4] == mark && board.positions[8] == mark ||
		board.positions[2] == mark && board.positions[4] == mark && board.positions[6] == mark {
		return true
	}

	return false
}

func (board Board) isFull() bool {
	emptyPos := 0
	for pos := range board.positions {
		if board.positions[pos] == Empty {
			emptyPos += 1
		}
	}
	return emptyPos == 0
}

func (board Board) isPositionEmpty(input int) bool {
	return board.positions[input] == Empty
}

func (board Board) isPositionValid(input int) bool {
	return 0 <= input && input < BoardSize && board.isPositionEmpty(input)
}

type Game struct {
	currentMove   int
	currentPlayer int
	board         Board
}

func (game *Game) play() {
	for {
		game.board.Print()

		fmt.Printf("Player %d is on a move \n", game.currentPlayer+1)
		fmt.Printf("Enter position in range [0-8]: ")
		pos := inputPosition()

		if !game.board.isPositionValid(pos) {
			fmt.Printf("Invalid input, please try again! \n")
			continue
		}

		mark := game.progress()

		game.board.Update(pos, mark)
		if game.board.isWinner(mark) || game.board.isFull() {
			break
		}
	}
}

func (game *Game) progress() int {
	game.currentPlayer = (game.currentPlayer + 1) % 2

	mark := Nought
	if game.currentPlayer == SecondPlayer {
		mark = Cross
	}

	return mark
}

func (game Game) results() {
	if game.board.isFull() {
		fmt.Printf("Nobody won. It's a draw! \n")
	} else {
		fmt.Printf("Player %d won! \n", game.currentPlayer)
	}
	game.board.Print()
}

// main program

func main() {
	fmt.Printf("Play TIC-TAC-TOE \n")

	board := Board{}
	game := Game{board: board}

	game.play()
	game.results()
}
