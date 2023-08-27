package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Interval struct {
	Start time.Time
	End   time.Time
}

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

		timeIntervals := make([]Interval, n)
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
				break
			}

			endTime, err := time.Parse(timeFormat, input[1])
			if err != nil {
				checkValidTimeIntervals = false
				break
			}

			timeIntervals[j] = Interval{
				Start: startTime,
				End:   endTime,
			}
		}

		if intervalsOverlap(timeIntervals) {
			checkValidTimeIntervals = false
		} else {
			checkValidTimeIntervals = true
		}

		if checkValidTimeIntervals {
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

func intervalsOverlap(timeIntervals []Interval) bool {
	sort.Slice(timeIntervals, func(i, j int) bool {
		return timeIntervals[i].Start.Before(timeIntervals[j].Start)
	})

	for i := 1; i < len(timeIntervals); i++ {
		prev := timeIntervals[i-1]
		cur := timeIntervals[i]
		if cur.Start.After(cur.End) || cur.Start.Before(prev.End) || prev.End.Equal(cur.Start) {
			return true
		}
		//if prev.Start.Before(cur.End) || prev.Start.Equal(cur.End) && cur.Start.Before(prev.End) || cur.Start.Equal(prev.End) {
		//	return true
		//}
	}
	return false
}
