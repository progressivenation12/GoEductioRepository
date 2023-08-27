package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Temperature struct {
	minTemp int
	maxTemp int
}

const (
	greaterOrEqual  = ">="
	lessOrEqual     = "<="
	nonExistingTemp = -1
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
	var parameter string
	rand.Seed(time.Now().UnixNano())

	_, err := fmt.Fscan(in, &t)
	if err != nil {
		return
	}

	for i := 0; i < t; i++ {
		temperature := Temperature{
			minTemp: 15,
			maxTemp: 30,
		}

		var n, temp int
		_, err := fmt.Fscan(in, &n)
		if err != nil {
			return
		}

		for i := 0; i < n; i++ {
			_, err := fmt.Fscan(in, &parameter, &temp)
			if err != nil {
				return
			}

			switch parameter {
			case greaterOrEqual:
				if temp > temperature.minTemp {
					temperature.minTemp = temp
				}
			case lessOrEqual:
				if temp < temperature.maxTemp {
					temperature.maxTemp = temp
				}
			}

			if temperature.minTemp == temperature.maxTemp {
				_, err = fmt.Fprintln(out, temperature.maxTemp)
				if err != nil {
				}
			} else if temperature.minTemp < temperature.maxTemp {
				_, err = fmt.Fprintln(out, temperature.minTemp)
				if err != nil {
					return
				}
			} else {
				_, err = fmt.Fprintln(out, nonExistingTemp)
				if err != nil {
				}
			}
		}
		_, err = fmt.Fprintln(out, "")
		if err != nil {
			return
		}
	}
}
