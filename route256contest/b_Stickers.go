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

	var n, start, end int
	var stickerLetter, r string

	_, err := fmt.Fscan(in, &stickerLetter, &n)
	if err != nil {
		return
	}

	stickerSlice := []byte(stickerLetter)

	for i := 0; i < n; i++ {
		_, err := fmt.Fscan(in, &start, &end, &r)
		if err != nil {
			return
		}

		for i := start - 1; ; {
			for j := 0; j < len(r); j++ {
				stickerSlice[i+j] = r[j]
			}
			break
		}
	}
	_, err = fmt.Fprintln(out, string(stickerSlice))
	if err != nil {
		return
	}
}
