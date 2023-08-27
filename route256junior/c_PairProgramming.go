package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t int

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			panic(err)
		}
	}(out)

	_, err := fmt.Fscan(in, &t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		a := make([]float64, 0)
		var n int
		var p float64

		_, err := fmt.Fscan(in, &n)
		if err != nil {
			return
		}

		for i := 0; i < n; i++ {
			_, err := fmt.Fscan(in, &p)
			if err != nil {
				return
			}
			a = append(a, p)
		}

		for i := 0; i < len(a); i++ {
			var index int
			diff := 100

			if a[i] != 0 {
				for j := i + 1; j < len(a); j++ {
					if a[j] != 0 {
						b := a[i] - a[j]
						if math.Abs(b) < float64(diff) {
							diff = int(math.Abs(b))
							index = j
						}
					}
				}
				_, err = fmt.Fprintln(out, i+1, index+1)
				if err != nil {
					return
				}
				a[i], a[index] = 0, 0
			}
		}
		_, err = fmt.Fprintln(out)
		if err != nil {
			return
		}
	}
}
