package music

import (
	"net/http"
	"strconv"

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
	logger.Log(logger.INFO, "Received cursor index:", c)
	if c > (len(mlist.Songs)-1) || c < 0 {
		response = lmsresponse.GetResponseBytes(lmsresponse.ERROR, "Invalid cursor index", nil)
		w.Write(response)
		return
	}
	var songSlice []SongMetadata
	if (c + PAGESIZE) > len(mlist.Songs)-1 {
		songSlice = mlist.Songs[c:]
	} else {
		songSlice = mlist.Songs[c:(c + PAGESIZE)]
	}
	data := MList{
		Songs: songSlice,
		Cursor: Cursor{
			HasNext:     hasNext(c + PAGESIZE),
			HasPrevious: hasPrevious(c + PAGESIZE),
			Index:       c + PAGESIZE,
			Length:      len(songSlice),
			Total:       len(mlist.Songs),
		},
	}
	response = lmsresponse.GetResponseBytes(lmsresponse.SUCCESS, "Here is your favorite music", data)
	w.Write(response)
}
