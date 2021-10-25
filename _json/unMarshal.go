package _json

import (
	"encoding/json"
	"git.catddm.com/yixing/lib/_string"
)

func ByteUnMarshal(b []byte, ret *map[string]interface{}) {
	json.Unmarshal(b, ret)
}
func StringUnMarshal(s string, ret *map[string]interface{}) {
	b := _string.ToBytes(s)
	ByteUnMarshal(b, ret)
}
