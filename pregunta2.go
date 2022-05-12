package main

import (
	"fmt"
)

var available = 0
var noAvailable = 1

type Solution struct {
	n            int   //número de reinas
	placedQueens []int //posiciones de las reinas
	step         int   //paso actual
}

func calculateavailablePosition(placedQueens []int, rowNum int, n int) []int {
	var row = make([]int, n)
	if rowNum == 0 {
		return row
	}

	for i := 0; i < rowNum; i++ {
		var prevQueen = placedQueens[i]
		row[prevQueen] = noAvailable
		var diff = rowNum - i
		if prevQueen-diff >= 0 {
			row[prevQueen-diff] = noAvailable
		}
		if prevQueen+diff <= n-1 {
			row[prevQueen+diff] = noAvailable
		}
	}
	return row
}

func recSolution(solution Solution) [][]int { //soluciones - matriz
	var step = solution.step + 1
	var rowNum = solution.step
	var availablePos = calculateavailablePosition(solution.placedQueens, rowNum, solution.n)
	var finalAns = [][]int{}

	if step == solution.n { //solución final
		for i := 0; i < solution.n; i++ {
			if availablePos[i] == available {
				solution.placedQueens[rowNum] = i
				var successAns = make([]int, solution.n)
				for j := 0; j < solution.n; j++ {
					successAns[j] = solution.placedQueens[j]
				}
				return append(finalAns, successAns)
			}
		}
		return finalAns
	} else { //paso siguiente
		for i := 0; i < solution.n; i++ {
			if availablePos[i] == available {
				solution.placedQueens[rowNum] = i
				solution.step = step
				var ans = recSolution(solution)
				for j := 0; j < len(ans); j++ {
					finalAns = append(finalAns, ans[j])
				}
			}
		}
		return finalAns
	}
}

func totalNQueens(n int) int { //número de soluciones
	if n == 1 {
		return 1
	}

	var finalAns = 0
	for i := 0; i < n; i++ { //reinas
		var placedQueens = make([]int, n) //posiciones de las reinas
		placedQueens[0] = i
		var solution = Solution{n, placedQueens, 1}
		var ans = recSolution(solution) //soluciones
		finalAns = finalAns + len(ans)
	}
	return finalAns
}

func main() {
	fmt.Printf("Número de soluciones: %d", totalNQueens(10))
}
