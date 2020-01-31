package music

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/abhijitWakchaure/lanmusic/gosrc/logger"
	"github.com/dhowden/tag"
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
	Songs  []SongMetadata `json:"songs"`
	Cursor Cursor         `json:"cursor"`
}

// SongMetadata holds the metadata about the song
type SongMetadata struct {
	Title    string `json:"title"`
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	Composer string `json:"composer"`
	Genre    string `json:"genre"`
	Year     int    `json:"year"`
	Path     string `json:"path"`
	// AlbumArt *tag.Picture `json:"albumArt"`
}

// Cursor will keep track of pagination
type Cursor struct {
	Index       int  `json:"index"`
	HasNext     bool `json:"hasNext"`
	HasPrevious bool `json:"hasPrevious"`
	Total       int  `json:"total"`
	Length      int  `json:"length"`
}

var mlist MList

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
				// logger.Log(logger.INFO, err.Error())
			}
			m, err := tag.ReadFrom(a)
			if err != nil {
				// logger.Log(logger.INFO, err.Error())
			} else {
				md := SongMetadata{
					Title:    m.Title(),
					Album:    m.Album(),
					Artist:   m.Artist(),
					Composer: m.Composer(),
					Genre:    m.Genre(),
					Year:     m.Year(),
					Path:     path.Join(musicroot, f.Name()),
					// AlbumArt: m.Picture()
				}
				mlist.Songs = append(mlist.Songs, md)
			}
		}
	}
	return nil
}

func listDir() MList {
	if ok, err := exists(MUSICROOT); !ok {
		logger.Log(logger.CRITICAL, "Make sure your mount target is /Music. Err:"+err.Error())
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
