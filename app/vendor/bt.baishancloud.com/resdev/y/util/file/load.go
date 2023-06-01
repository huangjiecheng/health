package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func LoadFile(filePath string, content interface{}) (err error) {
	if reflect.ValueOf(content).Kind() != reflect.Ptr {
		err = fmt.Errorf("content is not ptr")
		return
	}
	dataByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, content)
	if err != nil {
		return
	}
	return
}
