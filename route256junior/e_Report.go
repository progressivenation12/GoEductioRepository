package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			panic(err)
		}
	}(out)

	var t int
	_, err := fmt.Fscan(in, &t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		var n, a int
		aTask := make([]int, 0)

		_, err := fmt.Fscan(in, &n)
		if err != nil {
			return
		}

		for i := 0; i < n; i++ {
			_, err := fmt.Fscan(in, &a)
			if err != nil {
				return
			}
			aTask = append(aTask, a)
		}

		result := counter(aTask, len(aTask))

		if result {
			_, err := fmt.Fprintln(out, "Yes")
			if err != nil {
				return
			}
		} else {
			_, err := fmt.Fprintln(out, "No")
			if err != nil {
				return
			}
		}
	}

}

func counter(slice []int, value int) bool {
	info := make(map[int]int)

	for i := 0; i < value; i++ {
		job := slice[i]

		if prevIndex, ok := info[job]; ok {
			if prevIndex != i-1 {
				return false
			}
		}
		info[job] = i
	}
	return true
}
