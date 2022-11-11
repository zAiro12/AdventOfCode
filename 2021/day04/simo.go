package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrix struct {
	value [5][5]string
}

type matrix2 struct {
	value [5][5]bool
}

func main() {
	var s, call string
	var calls []string
	var n, ans2 int
	var attempts []matrix
	var booleans []matrix2
	// Memorise the calls
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	calls = strings.Split(s, ",")
	fmt.Println(calls)
	// Memorise the tables
	attempts = MemoriseTable()
	// Create a table of booleans
	booleans = make([]matrix2, len(attempts))
	for {
		// Get a number from calls
		call = calls[n]
		n++
		// Check the matrixs and put true in the same position for booleans
		booleans = checkMatrix(attempts, booleans, call)
		// check if someone has a row or a column of true values
		ans := checkBoolean(booleans)
		if ans != -1 {
			// if yes: sum the position of numbers of false values and multply by the number who finished the game (calls[n])
			ans2 = sumValues(attempts, booleans, ans)
			break
		}

		// if no: repeat since getting a call
	}
	// That's your answer
	ans3, _ := strconv.Atoi(calls[n-1])
	fmt.Println(ans2 * ans3)
}

func sumValues(mat []matrix, booleans []matrix2, n int) int {
	var sum, num int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if booleans[n].value[i][j] == false {
				num, _ = strconv.Atoi(mat[n].value[i][j])
				sum += num
			}
		}
	}
	return sum
}

func checkBoolean(booleans []matrix2) int {
	// This function return the number of table that wins, if any. If there isn't, returns -1
	for i := 0; i < len(booleans); i++ {
		for j := 0; j < 5; j++ {
			// check for rows
			if booleans[i].value[j][0] == true && booleans[i].value[j][1] == true && booleans[i].value[j][2] == true && booleans[i].value[j][3] == true && booleans[i].value[j][4] == true {
				return i
			}
		}
		for j := 0; j < 5; j++ {
			// check for columns
			if booleans[i].value[0][j] == true && booleans[i].value[1][j] == true && booleans[i].value[2][j] == true && booleans[i].value[3][j] == true && booleans[i].value[4][j] == true {
				return i
			}
		}
	}
	return -1
}

func checkMatrix(attempts []matrix, booleans []matrix2, call string) []matrix2 {
	for k := 0; k < len(attempts); k++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if attempts[k].value[i][j] == call {
					booleans[k].value[i][j] = true
				}
			}
		}
	}
	return booleans
}

func MemoriseTable() []matrix {
	var matr matrix
	var tables []matrix
	var n string
	for {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				fmt.Scan(&n)
				if n == "end" {
					return tables
				}
				matr.value[i][j] = n
			}
		}
		tables = append(tables, matr)
	}
}
