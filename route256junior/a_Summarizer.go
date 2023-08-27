package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var t, a, b, out int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	_, err := fmt.Fscan(reader, &t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		_, err := fmt.Fscan(reader, &a, &b)
		if err != nil {
			return
		}
		out = a + b

		_, err = writer.WriteString(strconv.Itoa(out) + "\n")
		if err != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		return
	}
}
