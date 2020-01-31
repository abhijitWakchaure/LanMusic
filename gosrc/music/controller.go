package music

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhijitWakchaure/lanmusic/gosrc/logger"

	"github.com/gorilla/mux"

	"github.com/abhijitWakchaure/lanmusic/gosrc/lmsresponse"
)

// ListMusic ...
func ListMusic(w http.ResponseWriter, r *http.Request) {
	mlist := listDir()
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Found %v\n", len(mlist.Songs))
	data := MList{
		Songs: mlist.Songs[:10],
		Cursor: Cursor{
			HasNext:     true,
			HasPrevious: false,
			Index:       10,
			Length:      10,
			Total:       len(mlist.Songs),
		},
	}
	response := lmsresponse.GetResponseBytes(lmsresponse.SUCCESS, "Here is your favorite music", data)
	// fmt.Printf("res %v\n", string(res))
	w.Write(response)
	return
}

//ListMusicWithCursor ...
func ListMusicWithCursor(w http.ResponseWriter, r *http.Request) {
	var response []byte
	cursor := mux.Vars(r)["cursor"]
	i, err := strconv.Atoi(cursor)
	// fmt.Println(i)
	if err != nil {
		response = lmsresponse.GetResponseBytes(lmsresponse.ERROR, "Expecting integer Value", nil)
	}
	logger.Log(logger.INFO, "Cursor:", i)
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Found %v\n", len(mlist.Songs))
	var from, to int
	from = i + 1
	to = from + 10
	if from > len(mlist.Songs) {
		from = 0
		to = from + 10
	}
	if to > len(mlist.Songs) {
		to = len(mlist.Songs)
	}
	data := MList{
		Songs: mlist.Songs[from:to],
		Cursor: Cursor{
			HasNext:     true,
			HasPrevious: true,
			Index:       i + 10,
			Length:      10,
			Total:       len(mlist.Songs),
		},
	}
	response = lmsresponse.GetResponseBytes(lmsresponse.SUCCESS, "Here is your favorite music", data)
	w.Write(response)
}
