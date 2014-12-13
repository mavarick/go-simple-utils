package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	DEFAULT_COMMENT     = "//"
	ALTERNATIVE_COMMENT = "#"
	SEPERATOR           = "="
)

type Config struct {
	File   string                 // the file path
	Values map[string]interface{} // the values with hashmap type
}

func NewConfig(file string) *Config {
	fp, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(fp)
	return &(Config{file, Read(buf)})
}

func Read(buf *bufio.Reader) map[string]interface{} {
	var values map[string]interface{} = make(map[string]interface{})

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		switch {
		case len(line) == 0, strings.HasPrefix(line, DEFAULT_COMMENT), strings.HasPrefix(line, ALTERNATIVE_COMMENT):
			break
		case strings.HasPrefix(line, SEPERATOR):
			panic(fmt.Sprintf("line[%s] starts with seperator[%s]", line, SEPERATOR))
		default:
			i := strings.IndexAny(line, SEPERATOR)
			if i == -1 {
				panic(fmt.Sprintf("line [%s] contains no Seperator [%s]", line[0:20], SEPERATOR))
			}
			key := strings.TrimSpace(line[0:i])
			value := strings.TrimSpace(line[(i + len(SEPERATOR)):len(line)])

			values[key] = value
		}
		// end of file
		if err == io.EOF {
			break
		}
		// error occurs
		if err != nil {
			panic(err)
		}
	}
	return values
}

func (self *Config) GetInt(key string) int64 {
	value := self.Get(key)
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func (self *Config) GetList(key, sep string) []string {
	result := self.Get(key)
	items := strings.Split(result, sep)
	for i := 0; i < len(items); i++ {
		items[i] = strings.TrimSpace(items[i])
	}
	return items
}

func (self *Config) Get(key string) string {
	value, ok := self.Values[key]
	if ok == false {
		panic(fmt.Sprintf("no key like [%s]", key))
	}
	return value.(string)
}

func (self *Config) GetFloat(key string) float64 {
	value := self.Get(key)
	result, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func (self *Config) ToJson() string {
	result, err := json.Marshal(self.Values)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (self *Config) SetKey(key string, value interface{}) {
	self.Values[key] = value
}
