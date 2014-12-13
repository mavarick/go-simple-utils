package log

type LogRecord struct {
	level string
	title string
	msg   interface{}
}

func NewLogRecord(msg interface{}, title string, level string) *LogRecord {
	record := &(LogRecord{})
	record.msg = msg
	record.title = title
	record.level = level
	return record
}
