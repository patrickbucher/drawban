package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/patrickbucher/drawban"
)

// usage:	go run cmd/demo.go border line1 line2 ... lineN
// example:	go run cmd/demo.go '#' Hello World
// output:
// #########
// # Hello #
// # World #
// #########
func main() {
	border := rune(os.Args[1][0])
	lines := os.Args[2:]
	banner := strings.Join(drawban.DrawBanner(lines, border), "\n")
	fmt.Println(banner)
}
