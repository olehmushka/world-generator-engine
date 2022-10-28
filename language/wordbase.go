package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Wordbase struct {
	FemaleOwnNames []string `json:"female_own_names" bson:"female_own_names"`
	MaleOwnNames   []string `json:"male_own_names" bson:"male_own_names"`
	Words          []string `json:"words" bson:"words"`
	Slug           string   `json:"slug" bson:"slug"`
	Min            int      `json:"min" bson:"min"`
	Max            int      `json:"max" bson:"max"`
	Dupl           string   `json:"dupl" bson:"dupl"`
	M              float64  `json:"m" bson:"m"`
}

func loadAllWordbases(prefix string) ([]*Wordbase, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/" + prefix
	dirname := currDirname + "/data/wordbases/"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))
	}

	out := make([]*Wordbase, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := dirname + file.Name()
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		var sfs []*Wordbase
		if err := json.Unmarshal(b, &sfs); err != nil {
			return nil, err
		}
		out = append(out, sfs...)
	}

	return out, nil
}

func LoadAllWordbases() ([]*Wordbase, error) {
	return loadAllWordbases("")
}
