package main

import (
	"bufio"
	"os"
)

func main() {
	//in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			panic(err)
		}
	}(out)

}
