/* find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

For example:

) causes him to enter the basement at character position 1.
()()) causes him to enter the basement at character position 5.
What is the position of the character that causes Santa to first enter the basement? */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var c int
	var i int
	var mem string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		mem = l
	}
	r := []rune(mem)
	for i = 0; i < len(r); i++ {
		if r[i] == '(' {
			c++
		}
		if r[i] == ')' {
			c--
		}
		if c == -1 {
			fmt.Println(i+1)
			break
		}
	}
}
