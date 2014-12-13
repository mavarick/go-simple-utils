package main

import (
	"fmt"
	"lib/log"
	"lib/tools"
	"time"
)

func main() {
	log.InitGLogger(nil)
	t := &(tools.TimeTools{})
	t1 := int64(10000000)
	t2 := t.ToTime(t1, 0)
	fmt.Println(t2)
	now := t.Now()
	if t.Before(now, t2) != false {
		fmt.Println(now, t2)
		log.GLogger.Fatal("err", "err body")
	}
	duration := t.Sub(now, t2)
	fmt.Println("duration:", duration)
	// to timestamp
	stamp := t.ToTimeStamp(now)
	fmt.Println("stamp", stamp)

	stamp = t.ToTimeStamp(t2)
	fmt.Println("stamp of t2: ", stamp)

	time.Sleep(50 * time.Millisecond)
}
