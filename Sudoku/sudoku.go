package main

import (
	"fmt"
)

func main() {

	var sudoku = [9][9]int{
		{9, 3, 4, 5, 6, 8, 1, 2, 7},
		{8, 2, 6, 7, 1, 4, 5, 9, 3},
		{1, 5, 7, 9, 2, 3, 4, 6, 8},
		{2, 7, 8, 1, 5, 9, 3, 4, 6},
		{6, 4, 1, 3, 8, 7, 2, 5, 9},
		{3, 9, 5, 6, 4, 2, 7, 8, 1},
		{5, 6, 3, 4, 9, 1, 8, 7, 2},
		{7, 8, 9, 2, 3, 5, 6, 1, 4},
		{4, 1, 2, 8, 7, 6, 9, 3, 5}}

	var a = [3][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0}}

	ok := true

Sudoku:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			b := sudoku[i][j] - 1
			t := 1 << b

			if a[0][i]&t != 0 {
				ok = false
				fmt.Println("Probleme pe linia", i)
				break Sudoku
			} else {
				a[0][i] = a[0][i] | t
			}

			if a[1][j]&t != 0 {
				ok = false
				fmt.Println("Probleme pe coloana", j)
				break Sudoku
			} else {
				a[1][j] = a[1][j] | t
			}

			k := i/3*3 + j/3
			if a[2][k]&t != 0 {
				ok = false
				fmt.Println("Probleme in subset", k)
				break Sudoku

			} else {
				a[2][k] = a[2][k] | t
			}
		}
	}

	if ok {
		fmt.Println("Configuratia este buna!")
	} else {
		fmt.Println("Configuratia nu este buna!")
	}

	fmt.Println(" ----------------------- ")
	for i := 0; i < 9; i++ {
		fmt.Print("| ")
		for j := 0; j < 9; j++ {
			if j == 3 || j == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", sudoku[i][j])
			if j == 8 {
				fmt.Print("|")
			}
		}
		if i == 2 || i == 5 || i == 8 {
			fmt.Println("\n ------------------------")
		} else {
			fmt.Println()
		}
	}

}
