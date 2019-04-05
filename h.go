package h

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Respond to an HTTP request by marshaling a value to JSON
func JsonResponse(w http.ResponseWriter, val interface{}, code ...int) {
	b, err := json.Marshal(val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(200)
	}
	w.Write(b)
}

// Unmarshal http.Request.Body into a target struct, closes Body.
func JsonBody(r *http.Request, target interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}
