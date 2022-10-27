package language

import (
	"encoding/json"
	"io/ioutil"
	"path"
	"runtime"
)

func loadAllFamilies(prefix string) ([]string, error) {
	_, filename, _, _ := runtime.Caller(1)
	b, err := ioutil.ReadFile(path.Dir(filename) + "/" + prefix + "/data/families/families.json")
	if err != nil {
		return nil, err
	}
	var out []string
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}

	return out, nil
}

func LoadAllFamilies() ([]string, error) {
	return loadAllFamilies("")
}
