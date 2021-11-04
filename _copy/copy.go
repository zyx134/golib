package copy

import (
	"bytes"
	"encoding/gob"
)

//深拷贝
//第二个值必须是指针
//err := Copy(a, &b)
//fmt.Println(err)
func Copy(a, b interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(a); err != nil {
		return err
	}
	if err := dec.Decode(b); err != nil {
		return err
	}
	return nil
}
