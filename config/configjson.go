package config

import "io/ioutil"

/* NOTICE:
IF I WANT TO UNMARSHAL ONE JSON FILE,
I MUST KNOW THE JSON STRUCTURE OF IT
*/

type ConfigJson struct {
	File    string // comfiguration file
	Content string // the raw content
}

func NewConfigJson(path string) *ConfigJson {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return &(ConfigJson{path, string(content)})
}

// TODO
func ParseJson(c string) *interface{} {
	return nil
}
