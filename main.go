package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := len(os.Args)
	var col string
	if a == 1 {

		chars := "0123456789abcdef"

		for range 6 {
			n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
			col += string(chars[n.Int64()])
		}
		fmt.Printf("GENERATED RANDOM COLOR: #%s\n", col)
	} else if a == 2 {
		col += strings.TrimPrefix(os.Args[1], "#")
	} else {
		os.Exit(1)
	}

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
