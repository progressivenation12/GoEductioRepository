package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type matrix [][]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int

	_, err := fmt.Fscan(reader, &t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		var n, m int

		_, err := fmt.Fscan(reader, &n, &m)
		if err != nil {
			return
		}

		matrix := make(matrix, n)

		for i := range matrix {
			matrix[i] = make([]int, m)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				var p int
				_, err := fmt.Fscan(reader, &p)
				if err != nil {
					return
				}
				matrix[i][j] = p
			}
		}

		var k int

		_, err = fmt.Fscan(reader, &k)
		if err != nil {
			return
		}

		for i := 0; i < k; i++ {
			var c int

			_, err = fmt.Fscan(reader, &c)
			if err != nil {
				return
			}
			sort.SliceStable(matrix, func(i, j int) bool {
				return matrix[i][c-1] < matrix[j][c-1]
			})
		}

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				_, err = fmt.Fprint(writer, matrix[i][j], " ")
				if err != nil {
					return
				}
			}
			_, err = fmt.Fprint(writer, "\n")
			if err != nil {
				return
			}
		}
		_, err = fmt.Fprint(writer, "\n")
		if err != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
