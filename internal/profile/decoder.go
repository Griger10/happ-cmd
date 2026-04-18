package profile

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

func DecodeUrl(url string) (string, error) {
	if !strings.HasPrefix(url, "happ://routing/") {
		return "", errors.New("invalid happ url")
	}

	strippedString := url[strings.LastIndex(url, "/")+1:]
	decodedString, err := base64.StdEncoding.DecodeString(strippedString)
	if err != nil {
		return "", errors.New("invalid base64 string")
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, decodedString, "", "  ")
	if err != nil {
		return "", errors.New("invalid JSON")
	}
	return prettyJSON.String(), nil
}
