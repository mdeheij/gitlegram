package gitlab

import (
	"encoding/json"
	"errors"
)

func Parse(jsonBody string) (r Request, err error) {
	err = json.Unmarshal([]byte(jsonBody), &r)

	if !r.IsValid() {
		return r, errors.New("Request invalid or unsupported")
	}

	return r, err
}
