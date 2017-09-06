package gitlab

import "encoding/json"

func Parse(jsonBody string) (r Request) {

	err := json.Unmarshal([]byte(jsonBody), &r)
	if err != nil {
		panic(err)
	}

	return r
}
