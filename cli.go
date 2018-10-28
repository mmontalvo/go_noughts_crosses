package main
import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "os"
)

// Missing:
// - do input validation
// - split into other packages (game and board)
// - improve checkWinner

var winner bool = false
var current_player string = "x"
var other_player   string = "o"
var board = [9]string{" "," "," "," "," "," "," "," "," "}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var text string
    for winner != true {
        fmt.Print("Enter coordinates (row,col): ")
        scanner.Scan()
        text = scanner.Text()
        row, column := getScanCoordinates(text)
        if (board[(row * 3) + column]) != " "{
          fmt.Println("Coordinates are not empty, please choose a new one:")
        }else{
          board[(row * 3) + column] = current_player
          if checkWinner() {
              winner = true
              fmt.Println("WINNER!", current_player)
          }
          current_player, other_player = other_player, current_player
        }
        printBoard(board)
    }
}

func winnerCombination(combination [3]int) bool {
  if (board[combination[0]] == current_player &&
      board[combination[1]] == current_player &&
      board[combination[2]] == current_player){
    return true
  }
  return false
}

func checkWinner() bool{
  if (winnerCombination([3]int {0, 1, 2}) ||
      winnerCombination([3]int {3, 4, 5}) ||
      winnerCombination([3]int {6, 7, 8}) ||
      winnerCombination([3]int {0, 3, 6}) ||
      winnerCombination([3]int {1, 4, 7}) ||
      winnerCombination([3]int {2, 5, 8}) ||
      winnerCombination([3]int {0, 4, 8}) ||
      winnerCombination([3]int {2, 4, 6})) {
       return true
     }
  return false
}

func getScanCoordinates(input string) (int, int){
  split := strings.Split(input, ",")
  row, _ := strconv.Atoi(split[0])
  column, _ := strconv.Atoi(split[1])
  return row, column
}

func printBoard(board [9]string) {
  for i := 0; i < 3; i++ {
		fmt.Printf("%s %s %s\n", board[(i * 3) + 0], board[(i * 3) + 1], board[(i * 3) + 2])
	}
}
