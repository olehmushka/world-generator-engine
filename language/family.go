package language

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

func LoadAllFamilies() ([]string, error) {
	_, fn, _, _ := runtime.Caller(1)
	currDirname := path.Dir(fn) + "/"
	filename := currDirname + "/data/families/families.json"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var out []string
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}

	return out, nil
}
