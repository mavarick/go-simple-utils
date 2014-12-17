package main

import (
	"fmt"
	"lib/flagtool"
)

func main() {
	params := flagtool.ParseFlag()

	fmt.Println(params)
}
