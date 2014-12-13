package main

import (
	"lib/log"
	"time"
)

func main() {
	filename := "run.log"

	log.InitGLogger(filename)
	//log.GetGLogger().Info("test", "this is the test")
	log.GLogger.Info("test", "this is the test")

	time.Sleep(time.Millisecond * 50)
	log.PRINT_WITH_TIME = false
	//log.GetGLogger().Info("test2", "this is the test2")

	time.Sleep(time.Millisecond * 50)
	//log.GetGLogger().Close()

}
