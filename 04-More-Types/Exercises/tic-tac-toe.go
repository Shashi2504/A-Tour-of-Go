package main
import (
	"fmt"
	"strings"
)

func main() {
	board_game := [][] string {
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},

	}

	board_game[0][0] = "X"
	board_game[2][2] = "O"
	board_game[1][2] = "X"
	board_game[1][0] = "O"
	board_game[0][2] = "X"

	for i := 0; i < len(board_game); i++ {
	fmt.Printf("%s\n", strings.Join(board_game[i], " "))
	}

}
