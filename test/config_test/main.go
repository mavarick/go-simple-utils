package main

import (
	"fmt"
	"lib/config"
)

func main() {
	config := config.NewConfig("t.cfg")

	fmt.Println(config)
	fmt.Println(config.ToJson())
}
