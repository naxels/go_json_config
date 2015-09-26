package config

import (
	"encoding/json"
	"io/ioutil"
)

//Parse modifies the given struct containing Configuration specification after opening the file
func Parse(fileLocation string, c interface{}) error {
	file, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &c)
	return err
}
