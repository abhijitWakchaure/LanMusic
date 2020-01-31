package music

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/abhijitWakchaure/lanmusic/gosrc/lmsresponse"
)

// ListMusic ...
func ListMusic(w http.ResponseWriter, r *http.Request) {
	mlist := listDir()
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Found %v\n", len(mlist.Songs))
	response := lmsresponse.IResponse{
		Status:  "success",
		Message: "Here is your music",
		Data:    mlist.Songs[:11],
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
	// fmt.Printf("res %v\n", string(res))
	w.Write(res)
	return
}

//ListMusicWithCursor ...
func ListMusicWithCursor(w http.ResponseWriter, r *http.Request) {
	var res []byte
	cursor := mux.Vars(r)["cursor"]
	i, err := strconv.Atoi(cursor)
	if err != nil {
		res, _ = json.Marshal(struct {
			status  string
			message string
		}{
			status:  "error",
			message: "Expecting integer Value",
		})
	}
	// fmt.Println(i)
	SetResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("Found %v\n", len(mlist.Songs))
	var from, to int
	from = i + 1
	to = from + 11
	if from > len(mlist.Songs) {
		from = 0
		to = from + 11
	}
	if to > len(mlist.Songs) {
		to = len(mlist.Songs)
	}
	response := lmsresponse.IResponse{
		Status:  "success",
		Message: "Here is your music",
		Data:    mlist.Songs[from:to],
	}
	res, err = json.Marshal(response)
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
}
