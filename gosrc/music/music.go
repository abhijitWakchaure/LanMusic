package music

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/abhijitWakchaure/lanmusic/gosrc/lmsresponse"
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

// SetResponseHeaders ...
func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func listDir() MusicList {
	var mlist MusicList
	if ok, err := exists(MUSICROOT); !ok {
		logger.Log(logger.CRITICAL, "Make sure your mount target is /Music")
		fmt.Println(err)
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

// ListMusic ...
func ListMusic(w http.ResponseWriter, r *http.Request) {
	mlist := listDir()
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Found %v\n", len(mlist.Filenames))
	response := lmsresponse.IResponse{
		Status:  "success",
		Message: "Here is your music",
		Data:    mlist.Filenames,
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
	fmt.Printf("res %v\n", string(res))
	w.Write(res)
	return
}
