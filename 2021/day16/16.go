package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var dec string
	fmt.Scan(&dec)

	operazioni := make(map[int]string)
	scanner := bufio.NewScanner(os.Stdin)
	var uno int
	var due string
	for scanner.Scan() {
		in := scanner.Text()
		fmt.Sscanf(in, "%d = %s", &uno, &due)
		operazioni[uno] = due
	}

	fmt.Println(decToBin(dec, operazioni))
}

func decToBin(s string, operazioni map[int]string) int {
	var out string
	for i := 0; i < len(s); i++ {
		out += operazioni[int(s[i])]
	}
	bin, _ := strconv.Atoi(out)
	return bin
}
