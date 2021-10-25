package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//需要指定_id，_id已存在返回409
func CreateById() {
	es, err := GetEsClient()
	id := "testtttt"
	failOnError(err, "Error creating the client")
	// Create creates a new document in the index.
	// Returns a 409 response when a document with a same ID already exists in the index.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title":   "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.Create(index, id, &buf, es.Create.WithDocumentType(documentType))
	if err != nil {
		failOnError(err, "Error create response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
