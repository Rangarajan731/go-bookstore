package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, w interface{}) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal([]byte(body), w)
		return
	}
}
