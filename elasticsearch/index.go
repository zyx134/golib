package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//index 在索引中创建或更新文档
//索引不存在的情况下，会自动创建索引。
//默认的_type（类型）是_doc，下面是指定doc类型创建添加的。
func Index() {
	es, err := GetEsClient()
	failOnError(err, "Error creating the client")
	// Index creates or updates a document in an index
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title":   "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.Index(index, &buf, es.Index.WithDocumentType(documentType))
	if err != nil {
		failOnError(err, "Error Index response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
