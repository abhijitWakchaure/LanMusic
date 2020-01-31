package music

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/abhijitWakchaure/lanmusic/gosrc/logger"
)

// MUSICROOT is the parent directory for your music
var MUSICROOT = "/Music"

func init() {
	if ok, _ := exists(MUSICROOT); !ok {
		MUSICROOT = path.Join(os.Getenv("HOME"), "Music")
		logger.Log(logger.INFO, fmt.Sprintf("Using new music root %s instead of default root /Music", MUSICROOT))
	}
}

// MList ...
type MList struct {
	Filenames []MObject
}

// MObject ...
type MObject struct {
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
}

func songPath(files []os.FileInfo, mlist *MList, musicroot string) error {
	for _, f := range files {
		if f.IsDir() {
			curRoot := path.Join(musicroot, f.Name())
			files, err := ioutil.ReadDir(curRoot)
			if err != nil {
				log.Fatal(err)
			}
			err = songPath(files, mlist, curRoot)
			if err != nil {
				log.Fatal(err)
			}
		}
		if !f.IsDir() {
			md := MObject{Name: f.Name(), Size: f.Size(), IsDir: f.IsDir(), Path: path.Join(musicroot, f.Name())}
			mlist.Filenames = append(mlist.Filenames, md)
		}
	}
	return nil
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
	err = songPath(files, &mlist, MUSICROOT)
	if err != nil {
		log.Fatal(err)
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
