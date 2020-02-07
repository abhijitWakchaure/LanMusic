package music

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/abhijitWakchaure/lanmusic/gosrc/logger"

	"github.com/gorilla/mux"

	"github.com/abhijitWakchaure/lanmusic/gosrc/lmsresponse"
)

// ListMusic ...
func ListMusic(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	songSlice := mlist.Songs[:PAGESIZE]
	data := MList{
		Songs: songSlice,
		Cursor: Cursor{
			HasNext:     hasNext(0),
			HasPrevious: hasPrevious(0),
			Index:       PAGESIZE,
			Length:      len(songSlice),
			Total:       len(mlist.Songs),
		},
	}
	response := lmsresponse.GetResponseBytes(lmsresponse.SUCCESS, "Here is your favorite music", data)
	// fmt.Printf("res %v\n", string(res))
	w.Write(response)
	return
}

//SongSearcher returns true if search terms present in song
func SongSearcher(sq string, song *SongMetadata) bool {
	songString, _ := json.Marshal(song)
	if strings.Contains(string(songString), sq) {
		return true
	}
	return false
}

//SearchMusic searches a track in local songs root directory
func SearchMusic(w http.ResponseWriter, r *http.Request) {
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)

	searchQuery := mux.Vars(r)["searchQuery"]
	var searchResultList []SongMetadata

	for _, song := range mlist.Songs {
		isPresent := SongSearcher(searchQuery, &song)
		if isPresent {
			searchResultList = append(searchResultList, song)
		}
	}
	response := lmsresponse.GetResponseBytes(lmsresponse.SUCCESS, "Here is your favorite music", searchResultList)
	w.Write(response)
}

//StreamMusic starts streaming music
func StreamMusic(w http.ResponseWriter, r *http.Request) {
	var found bool
	songID := mux.Vars(r)["songID"]
	for _, song := range mlist.Songs {
		if song.ID == songID {
			fpath := song.Path
			found = true
			http.ServeFile(w, r, fpath)
			break
		}
		fmt.Println(song.ID)
	}
	if !found {
		response := lmsresponse.GetResponseBytes(lmsresponse.ERROR, "No such song", songID)
		w.Write(response)
	}
}

//ListMusicWithCursor ...
func ListMusicWithCursor(w http.ResponseWriter, r *http.Request) {
	var response []byte
	cString := mux.Vars(r)["cursor"]
	c, err := strconv.Atoi(cString)
	if err != nil {
		response = lmsresponse.GetResponseBytes(lmsresponse.ERROR, "Expecting integer Value", nil)
	}
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	logger.Log(logger.DEBUG, "Received cursor index:", c)
	if c > (len(mlist.Songs)-1) || c < 0 {
		response = lmsresponse.GetResponseBytes(lmsresponse.ERROR, "Invalid cursor index", nil)
		w.Write(response)
		return
	}
	var songSlice []SongMetadata
	var index int
	if (c + PAGESIZE) > len(mlist.Songs)-1 {
		songSlice = mlist.Songs[c:]
		index = 0
	} else {
		songSlice = mlist.Songs[c:(c + PAGESIZE)]
		index = c + PAGESIZE
	}
	data := MList{
		Songs: songSlice,
		Cursor: Cursor{
			HasNext:     hasNext(c + PAGESIZE),
			HasPrevious: hasPrevious(c + PAGESIZE),
			Index:       index,
			Length:      len(songSlice),
			Total:       len(mlist.Songs),
		},
	}
	response = lmsresponse.GetResponseBytes(lmsresponse.SUCCESS, "Here is your favorite music", data)
	w.Write(response)
}
