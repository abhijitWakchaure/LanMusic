package music

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/abhijitWakchaure/lanmusic/gosrc/logger"
)

// const for default config
const (
	MUSICROOT = "/Music"
)

// MList ...
type MList struct {
	Filenames []MObject
}

// MObject ...
type MObject struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	IsDir bool   `json:"isDir"`
}

func listDir() MList {
	var mlist MList
	if ok, err := exists(MUSICROOT); !ok {
		logger.Log(logger.CRITICAL, "Make sure your mount target is "+MUSICROOT+" Err:"+err.Error())
	}
	files, err := ioutil.ReadDir(MUSICROOT)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		md := MObject{Name: f.Name(), Size: f.Size(), IsDir: f.IsDir()}
		mlist.Filenames = append(mlist.Filenames, md)

	}
	return mlist
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, errors.New("Path does not exists")
		} else {
			return false, err
		}
	}
	return true, nil
}

// SetResponseHeaders ...
func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
