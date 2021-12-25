package main
import (
  "fmt"
  "time"
)

var xPlayer = "X"
var oPlayer = "O"
var aiTurn = true
var deep int = 0
var playAgain string
func main() {
	
  playingBoard := map[uint]string{1: "-", 2: "-", 3: "-",
                                  4: "-", 5: "-", 6: "-", 
                                  7: "-", 8: "-", 9: "-"}

  var gameNotEnded bool = true

  for gameNotEnded {
    if (aiTurn == false) {
      var spot uint
      GETSPOT: fmt.Println("Please enter a spot. Numbering starts from 1 in the top-left")
      fmt.Scanln(&spot)
      if (playingBoard[spot] == "-") {
        playingBoard[spot] = oPlayer
      } else {
        fmt.Println("Please enter a valid value")
        goto GETSPOT
      }
    } else {
      deep = 0
      fmt.Println("Bot's turn.....")
      time.Sleep(2 * time.Second)
      computer_move := computerMove(playingBoard)
      fmt.Println("Chose", computer_move)
      fmt.Println("Searched", deep, "levels deep")
      playingBoard[computer_move] = xPlayer
    }
    drawBoard(playingBoard)
    
    if (checkWin(playingBoard, oPlayer)) {
      fmt.Println("Man triumphed over Machine!!")
      gameNotEnded = false
    } else if (checkWin(playingBoard, xPlayer)) {
      fmt.Println("Machine triumphed over Man!!")
      gameNotEnded = false
    } else if (checkTie(playingBoard)) {
      fmt.Println("It's a tie :(")
      gameNotEnded = false
    }
    aiTurn = !aiTurn
  }
}


func miniMax( board map[uint]string, depth int, isMaxing bool ) int {
  deep++
  if (checkWin(board, xPlayer)) {
    return 1
  } else if (checkWin(board, oPlayer)) {
    return -1
  } else if (checkTie(board)) {
    return 0
  }
  
  if isMaxing {
    bestScore := -1000
    for key, value := range board {
      if value == "-" {
        board[key] = xPlayer
        score := miniMax(board, depth+1, false)
        board[key] = "-"
        if score > bestScore {
          bestScore = score
        }
      }
    } 
    return bestScore
  } else {
    bestScore := 1000
    for key, value := range board {
      if value == "-" {
        board[key] = oPlayer
        score := miniMax(board, depth+1, true) 
        board[key] ="-"
        if score < bestScore {
          bestScore = score
        }
      }
    }
    return bestScore
  }
}

func checkWin (board map[uint]string, char string) bool {
  if ( 
      (board[1] == board[2] && board[1] == board[3] && board[1] == char) ||
      (board[4] == board[5] && board[4] == board[6] && board[4] == char) ||
      (board[7] == board[8] && board[7] == board[9] && board[7] == char) ||
      (board[1] == board[4] && board[1] == board[7] && board[1] == char) ||
      (board[2] == board[5] && board[2] == board[8] && board[2] == char) ||
      (board[3] == board[6] && board[3] == board[9] && board[3] == char) ||
      (board[1] == board[5] && board[1] == board[9] && board[1] == char) ||
      (board[3] == board[5] && board[3] == board[7] && board[3] == char)) {
        return true
      } else {
        return false
      }
}

func checkTie(board map[uint]string) bool {
  for _, v := range board {
    if (v == "-") {
      return false
    }
    
  }
  return true
}

func computerMove(board map[uint]string) uint {
  var bestMove uint
  var bestScore = -1000
  bestMove = 0 
  
  for key, value := range board {
    if value == "-" {
      board[key] = xPlayer
      score := miniMax(board, 0, false)
      board[key] = "-"
      if score > bestScore {
        bestScore = score
        bestMove = key
      }
    }
  }
  return bestMove
}

func drawBoard (board map[uint]string) {
  fmt.Println(board[1] + "  |  " + board[2] + "  |  " + board[3])
  fmt.Println("-------------")
  fmt.Println(board[4] + "  |  " + board[5] + "  |  " + board[6])
  fmt.Println("-------------")
  fmt.Println(board[7] + "  |  " + board[8] + "  |  " + board[9])
}