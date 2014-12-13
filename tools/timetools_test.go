package tools

import (
	"fmt"
	"testing"
)

func Test_timetools(test *testing.T) {
	t := &(TimeTools{})
	t1 := int64(10000000)
	t2 := t.ToTime(t1, 0)
	fmt.Println(t2)
	now := t.Now()
	if t.Before(now, t2) != false {
		fmt.Println(now, t2)
		test.Fatal("error")
	}
	duration := t.Sub(now, t2)
	fmt.Println("duration:", duration)
	// to timestamp
	stamp := t.ToTimeStamp(now)
	fmt.Println("stamp", stamp)

	stamp = t.ToTimeStamp(t2)
	fmt.Println("stamp of t2: ", stamp)
}
