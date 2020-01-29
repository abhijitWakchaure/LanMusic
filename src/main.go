package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"lanmusic/logger"
	"log"
	"os"
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

func main() {
	fmt.Println("### LanMusic ###")
	mlist := listDir()
	for i, ml := range mlist.Filenames {
		fmt.Println(i+1, ml)
	}
}
