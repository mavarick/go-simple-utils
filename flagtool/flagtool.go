package flagtool

import (
	"flag"
	"fmt"
	"syscall"
)

const (
	usage = "Print the Usage"
)

const USAGE = `Usage:
	this program is mainly for ......

	for getting usage, pls type:
		-help/usage, will print the usage of the program.
`

func ParseFlag() map[string]interface{} {
	isUsage := flag.Bool("usage", false, usage)
	isHelp := flag.Bool("help", false, usage)
	// add flag option here to get new parameter, TODO

	flag.Parse()
	if *isHelp == true || *isHelp == true {
		fmt.Println(USAGE)
		syscall.Exit(0)
	}

	params := make(map[string]interface{})
	params["usage"] = *isUsage
	// add code here to add new values , TODO
	return params
}
