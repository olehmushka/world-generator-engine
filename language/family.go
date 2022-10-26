package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadAllFamilies(prefix string) ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	b, err := ioutil.ReadFile(pwd + "/" + prefix + "/data/families/families.json")
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
	return loadAllFamilies("language")
}
