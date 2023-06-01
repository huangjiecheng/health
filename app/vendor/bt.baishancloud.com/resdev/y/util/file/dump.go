package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func DumpFile(filePath string, content interface{}) (err error) {
	dataByte, err := json.Marshal(content)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filePath, dataByte, os.ModePerm)
	if err != nil {
		return
	}
	return
}
