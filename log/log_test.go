package log

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	// create the logger
	filename := "run.log"
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("START >>")
	logger := NewLogger(writer)
	logger.Info("this is the tester", "test")
	time.Sleep(50 * time.Millisecond)

	newFilename := "run.log2"
	writer2, err := os.OpenFile(newFilename, os.O_CREATE|os.O_RDWR, 0644)
	logger.SetWriter(writer2)

	logger.Error("error tester", "test")
	time.Sleep(50 * time.Millisecond)
	logger.Close()
}

func TestCocurrence(t *testing.T) {
	num := 1000
	filename := "run.log"
	writer, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	logger := NewLogger(writer)
	var i int = 0
	for {
		if i >= num {
			break
		}
		go func() {
			info := fmt.Sprintf("[num]::%d::", i)
			title := fmt.Sprintf("title_%d", i)

			logger.Info(info, title)
		}()
		i++
	}
	time.Sleep(50 * time.Millisecond)
	logger.Close()
}
