package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	col := strings.TrimPrefix(os.Args[1], "#")

	if len(col) != 6 {
		os.Exit(2)
	}

	var rgb []int

	for i := 0; i < 6; i += 2 {
		r, err := strconv.ParseInt(col[i:i+2], 16, 64)
		if err != nil {
			fmt.Println("error parsing hex:", err)
			return
		}
		rgb = append(rgb, int(r))
	}

	colorEsc := fmt.Sprintf("\033[48;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
	reset := "\033[0m"

	for range 5 {
		fmt.Printf("\n%s         %s", colorEsc, reset)
	}
}
