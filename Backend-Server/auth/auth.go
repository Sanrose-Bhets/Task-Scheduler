package authentication

import (
	"errors"
	"net/http"
	"strings"
)

// Getting a api key from the headers of an http request
// authorization: Apikey {insert api key}(header is like this)
func getAPIKey(headers http.Header) (string, error) {

	key := headers.Get("Authorization")
	if key == "" {
		return "", errors.New("Authorization header is missing")
	}

	apikey := strings.Split(key, "")
	if len(apikey) != 2 || apikey[0] != "Apikey" {
		return "", errors.New("Authorization header is not in the correct format")
	}

	return apikey[1], nil
}
