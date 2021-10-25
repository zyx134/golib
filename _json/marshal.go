package _json

import (
	"encoding/json"
	"git.catddm.com/yixing/lib/_byte"
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
