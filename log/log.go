package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// the configuration
var LOG_CHANNEL_CNT int = 512
var PRINT_TO_CONSOLE bool = true
var PRINT_TO_FILE bool = true
var PRINT_WITH_TIME bool = true

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	FATAL = "FATAL"
)

type Logger struct {
	writer     *os.File
	logChannel chan *LogRecord
}

func NewLogger(writer *os.File) *Logger {
	logger := &(Logger{})
	logger.writer = writer
	logger.logChannel = make(chan *LogRecord, LOG_CHANNEL_CNT)
	logger.Handle()
	return logger
}

func (self *Logger) Handle() {
	go func() {
		for record := range self.logChannel {
			self.WriteLog(record)
		}
	}()
}

func (self *Logger) ToString(record *LogRecord) string {
	buf, _ := json.Marshal(record.msg)
	msg := string(buf)
	level := record.level
	title := record.title
	return self.FormatOutput(title, msg, level)
}

func (self *Logger) FormatOutput(title, msg, level string) string {
	var output string
	if PRINT_WITH_TIME == true {
		now := time.Now().Format("2006-01-02 15:04:05 ")
		output = fmt.Sprintf("%s", now)
	}
	output = fmt.Sprintf("%s[%s]: title:: %s, msg:: %s", output, level, title, msg)
	return output
}

func (self *Logger) WriteLog(record *LogRecord) {
	buf := self.ToString(record)
	if self.writer == nil || PRINT_TO_CONSOLE == true {
		fmt.Println(buf + "\n")
	}
	if self.writer != nil {
		self.writer.WriteString(buf + "\n")
	}
}

// set writer
func (self *Logger) SetWriter(writer *os.File) {
	self.CloseWriter()
	self.writer = writer
}

// get writer
func (self *Logger) GetWriter() *os.File {
	return self.writer
}

// close
func (self *Logger) Close() {
	close(self.logChannel)
	self.CloseWriter()
}
func (self *Logger) CloseWriter() {
	if self.writer != nil {
		self.writer.Close()
	}
}

func (self *Logger) Debug(title string, msg interface{}) {
	self.logChannel <- NewLogRecord(msg, title, DEBUG)
}
func (self *Logger) Info(title string, msg interface{}) {
	self.logChannel <- NewLogRecord(msg, title, INFO)
}
func (self *Logger) Warn(title string, msg interface{}) {
	self.logChannel <- NewLogRecord(msg, title, WARN)
}
func (self *Logger) Error(title string, msg interface{}) {
	self.logChannel <- NewLogRecord(msg, title, ERROR)
}
func (self *Logger) Fatal(title string, msg interface{}) {
	self.logChannel <- NewLogRecord(msg, title, FATAL)
}

// initiate one object to use as one global logger
var GLogger *Logger = nil

// two ways to get logger instance to write the log
func GetGLogger() *Logger {
	if GLogger == nil {
		err := errors.New("Logger is not initialized")
		panic(err)
	}
	return GLogger
}

//as one server, the initGlogger() should be called at the begining of program
func InitGLogger(logFilePath interface{}) {
	fmt.Println("INIT>> initiate the global logger")
	if logFilePath == nil || logFilePath.(string) == "" {
		GLogger = NewLogger(nil)
	} else {
		writer, err := os.OpenFile(logFilePath.(string), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		GLogger = NewLogger(writer)
	}
}
