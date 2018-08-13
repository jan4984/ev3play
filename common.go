package ev3play

import (
	"path/filepath"
	"io/ioutil"
)

func chomp(b []byte) []byte {
	if b[len(b)-1] == '\n' {
		return b[:len(b)-1]
	}
	return b
}

func attributeOf(path, name, attr string) (string, error) {
	path = filepath.Join(path, name, attr)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", nil
	}
	return string(chomp(b)), nil
}

func setAttributeOf(path, name, attr, data string) error {
	path = filepath.Join(path, name, attr)
	err := ioutil.WriteFile(path, []byte(data), 0)
	if err != nil {
		return err
	}
	return nil
}
