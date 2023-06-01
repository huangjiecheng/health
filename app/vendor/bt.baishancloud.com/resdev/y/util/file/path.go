package file

import "os"

func PathExist(path string) (exist bool, err error) {
	_, err = os.Stat(path)
	if err == nil {
		exist = true
		return
	}
	if os.IsNotExist(err) {
		err = nil
		return
	}
	return
}
