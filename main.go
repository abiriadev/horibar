package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := bufio.NewScanner(os.Stdin)

	for lines.Scan() {
		line := lines.Text()

		i, e := strconv.Atoi(line)

		if e != nil {
			panic(e)
		}

		fmt.Println(strings.Repeat("█", i))
	}
}
