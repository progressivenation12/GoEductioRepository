package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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
		var n int
		_, err := fmt.Fscan(in, &n)
		if err != nil {
			return
		}
		timeIntervals := make([][2]time.Time, n)
		checkValidTimeIntervals := true

		for j := 0; j < n; j++ {
			var inputTimeString string
			_, err := fmt.Fscan(in, &inputTimeString)
			if err != nil {
				return
			}
			input := strings.Split(inputTimeString, "-")
			timeFormat := "15:04:05"
			startTime, err := time.Parse(timeFormat, input[0])
			if err != nil {
				checkValidTimeIntervals = false
			}
			endTime, err := time.Parse(timeFormat, input[1])
			if err != nil {
				checkValidTimeIntervals = false
			}
			timeIntervals[j] = [2]time.Time{startTime, endTime}
		}
		for j, interval := range timeIntervals {
			if !isValidTime(interval[0]) || !isValidTime(interval[1]) || interval[0].After(interval[1]) {
				checkValidTimeIntervals = false
				break
			}
			for k := 0; k < j; k++ {
				if isOverlapTime(timeIntervals[k], interval) {
					checkValidTimeIntervals = false
					break
				}
			}
		}
		if checkValidTimeIntervals {
			_, err := fmt.Fprintln(out, "YES")
			if err != nil {
				return
			}
		} else {
			_, err := fmt.Fprintln(out, "NO")
			if err != nil {
				return
			}
		}
	}
}
func isValidTime(t time.Time) bool {
	return t.Hour() >= 0 && t.Hour() <= 23 &&
		t.Minute() >= 0 && t.Minute() <= 59 &&
		t.Second() >= 0 && t.Second() <= 59
}
func isOverlapTime(a, b [2]time.Time) bool {
	return (a[0].Before(b[1]) || a[0].Equal(b[1])) && (b[0].Before(a[1]) || b[0].Equal(a[1]))
}
