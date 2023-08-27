package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	_, err := fmt.Fscan(reader, &t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		counter := make(map[int]int)
		var n, p, sum int
		_, err := fmt.Fscan(reader, &n)
		if err != nil {
			return
		}
		for i := 0; i < n; i++ {
			_, err := fmt.Fscan(reader, &p)
			if err != nil {
				return
			}
			key, _ := strconv.Atoi(strconv.Itoa(p))
			counter[key]++
		}
		for p, v := range counter {
			sum += ((v/3)*2 + (v % 3)) * p
		}

		_, err = fmt.Fprintln(writer, sum)
		if err != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		return
	}
}
