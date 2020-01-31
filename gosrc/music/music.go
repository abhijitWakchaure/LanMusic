package music

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/abhijitWakchaure/lanmusic/gosrc/logger"
	"github.com/dhowden/tag"
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
	Title    string       `json:"title"`
	Album    string       `json:"album"`
	Artist   string       `json:"artist"`
	Composer string       `json:"composer"`
	Genre    string       `json:"genre"`
	Year     int          `json:"year"`
	Path     string       `json:"path"`
	AlbumArt *tag.Picture `bson:"albumArt" json:"albumArt"`
}

func songPath(files []os.FileInfo, mlist *MList, musicroot string) error {
	for _, f := range files {
		if f.IsDir() {
			curRoot := path.Join(musicroot, f.Name())
			folder, err := os.Open(curRoot)
			if err != nil {
				log.Fatal(err)
			}
			files, err := folder.Readdir(-1)
			if err != nil {
				log.Fatal(err)
			}
			err = songPath(files, mlist, curRoot)
			if err != nil {
				log.Fatal(err)
			}
		}
		if !f.IsDir() {
			a, err := os.Open(path.Join(musicroot, f.Name()))
			defer a.Close()
			if err != nil {
				logger.Log(logger.INFO, err.Error())
			}
			m, err := tag.ReadFrom(a)
			if err != nil {
				logger.Log(logger.INFO, err.Error())
			} else {
				md := MObject{
					Title:    m.Title(),
					Album:    m.Album(),
					Artist:   m.Artist(),
					Composer: m.Composer(),
					Genre:    m.Genre(),
					Year:     m.Year(),
					Path:     path.Join(musicroot, f.Name()),
					AlbumArt: m.Picture()}
				mlist.Filenames = append(mlist.Filenames, md)
			}
		}
	}
	return nil
}

func listDir() MList {
	var mlist MList
	if ok, err := exists(MUSICROOT); !ok {
		logger.Log(logger.CRITICAL, "Make sure your mount target is "+MUSICROOT+" Err:"+err.Error())
	}
	// files, err := ioutil.ReadDir(MUSICROOT)
	folder, err := os.Open(MUSICROOT)
	if err != nil {
		log.Fatal(err)
	}
	files, err := folder.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	// os.Open(MUSICROOT).ReadDir()
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
