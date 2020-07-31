package utils

import (
	"encoding/json"
)

func ParseToString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	s := string(data)
	return s, nil
}
