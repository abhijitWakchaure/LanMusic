package music

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abhijitWakchaure/lanmusic/gosrc/lmsresponse"
)

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
