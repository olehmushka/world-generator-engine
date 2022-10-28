package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Subfamily struct {
	Slug              string     `json:"slug" bson:"slug"`
	FamilySlug        string     `json:"family_slug" bson:"family_slug"`
	ExtendedSubfamily *Subfamily `json:"extended_subfamily" bson:"extended_subfamily"`
}

func LoadAllSubfamilies() ([]*Subfamily, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/subfamilies/"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))
	}

	out := make([]*Subfamily, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := dirname + file.Name()
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		var sfs []*Subfamily
		if err := json.Unmarshal(b, &sfs); err != nil {
			return nil, err
		}
		out = append(out, sfs...)
	}

	return out, nil
}

func SearchSubfamily(slug string) (*Subfamily, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/subfamilies/"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := dirname + file.Name()
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		var sfs []*Subfamily
		if err := json.Unmarshal(b, &sfs); err != nil {
			return nil, err
		}
		for _, sf := range sfs {
			if sf.Slug == slug {
				return sf, nil
			}
		}
	}

	return nil, nil
}
