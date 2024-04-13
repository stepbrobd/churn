package lib

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

// takes in:
// - a local file path or a http{,s}:// url
// - a pointer to a slice of interface that's json marshalable
// returns
// - err if any
func Import(uri string, entities interface{}) error {
	var body []byte
	var err error

	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		// download the file
		res, err := http.Get(uri)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
	} else {
		body, err = os.ReadFile(uri)
		if err != nil {
			return err
		}
	}

	// unmarshal the file
	err = json.Unmarshal(body, &entities)
	if err != nil {
		return err
	}

	return nil
}
