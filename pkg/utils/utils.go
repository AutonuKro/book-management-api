package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, error := ioutil.ReadAll(r.Body); error == nil {
		if error := json.Unmarshal([]byte(body), x); error != nil {
			return
		}
	}
}
