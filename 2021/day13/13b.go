package main

import "fmt"

type coordinate struct {
	x, y int
}

func main() {
	var m = make(map[coordinate]bool)
	var in coordinate
	for {
		_, err := fmt.Scanf("%d,%d", &in.x, &in.y)
		if err != nil {
			break
		}
		m[in] = true
	}

	var c rune
	var num int
	for i := 0; i < 12; i++ {
		fmt.Scanf("fold along %c=%d", &c, &num)
		if c == 'x' {
			in.x = num
			in.y = 0
			m = rim(in, m)
		} else {
			in.x = 0
			in.y = num
			m = rim(in, m)
		}
	}

	fmt.Println(len(m))
	fmt.Println(stampa(m))
}

func rim(in coordinate, m map[coordinate]bool) map[coordinate]bool {
	var appoggio coordinate
	if in.x == 0 {
		for k := range m {
			if k.y >= in.y {
				appoggio.x = k.x
				appoggio.y = (in.y * 2) - k.y
				m[appoggio] = true
				delete(m, k)
			}
		}
	} else {
		for k := range m {
			if k.x >= in.x {
				appoggio.x = (in.x * 2) - k.x
				appoggio.y = k.y
				m[appoggio] = true
				delete(m, k)
			}
		}
	}
	return m
}

func stampa(m map[coordinate]bool) int {
	var app coordinate
	c := 0
	for i := 0; i < 6; i++ {
		app.y = i
		for j := 0; j < 39; j++ {
			app.x = j
			if m[app] {
				fmt.Print("*")
				c++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return c
}
