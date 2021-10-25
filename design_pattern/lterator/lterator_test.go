package lterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	arr := []interface{}{"a", "b", "c", "d"}
	arrayContainer := &ArrayContainer{arrayData: arr}
	iterator := arrayContainer.GetIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next().(string))
	}
}
