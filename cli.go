package main
import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
    "os"
    "math/rand"
)

var winner bool = false
var current_player string = "x"
var other_player   string = "o"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    board := [3][3]string{}
    var text string
    for winner != true {
        printBoard(board)
        fmt.Print("Enter coordinates (row,col): ")
        scanner.Scan()
        text = scanner.Text()
        row, column := getScanCoordinates(text)
        board[row][column] = current_player
        if checkWinner() {
            winner = true
        }
    }
}

func checkWinner() bool{
  return (rand.Intn(100) % 2) == 0
}

func getScanCoordinates(input string) (int, int){
  split := strings.Split(input, ",")
  row, _ := strconv.Atoi(split[0])
  column, _ := strconv.Atoi(split[1])
  return row, column
}

func printBoard(board [3][3]string) {
  fmt.Println(board)
}
