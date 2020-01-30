package music

import (
	"encoding/json"
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

// MusicList ...
type MusicList struct {
	Filenames []string
}

func listDir() MusicList {
	var mlist MusicList
	if ok, _ := exists(MUSICROOT); !ok {
		logger.Log(logger.CRITICAL, "Make sure your mount target is /Music")
	}
	files, err := ioutil.ReadDir(MUSICROOT)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		mlist.Filenames = append(mlist.Filenames, f.Name())
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

func ListMusic(w http.ResponseWriter, r *http.Request) {
	mlist := listDir()
	w.WriteHeader(http.StatusOK)
	response := struct {
		status  string
		message string
		data    []string
	}{
		status:  "success",
		message: "Here is your music",
		data:    mlist.Filenames,
	}
	res, err := json.Marshal(response)
	if err != nil {
		res, _ = json.Marshal(struct {
			status  string
			message string
		}{
			status:  "error",
			message: "Unable to list your music",
		})
	}
	w.Write(res)
	return
}
