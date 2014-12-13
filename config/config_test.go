package config

import (
	"fmt"
	"testing"
)

func TestConfigJson(t *testing.T) {
	//config := NewConfigJson("t.log")
	//if config.Content != "12345" {
	//	t.Fatal("error")
	//}
}

func TestConfig(t *testing.T) {
	config := NewConfig("t.cfg")
	if config.Values["ip"] != "127.0.0.1" {
		t.Fatal(fmt.Sprintf("value error: %s", config.Values["ip"]))
	}
	if config.GetInt("age") != 125 {
		t.Fatal("err age")
	}
	if len(config.GetList("fruits", ",")) != 3 {
		t.Fatal("err fruit length")
	}
	if config.GetList("fruits", ",")[2] != "banana" {
		t.Fatal("err fruit list value")
	}
}
