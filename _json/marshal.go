package _json

import (
	"encoding/json"

	"github.com/zyx134/golib/_byte"
)

func ToJsonByte(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
func ToJsonString(data interface{}) (string, error) {
	b, err := ToJsonByte(data)
	if err != nil {
		return "", err
	}
	return _byte.ToString(b), nil
}
