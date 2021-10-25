package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

func UpdateById() {
	es, err := GetEsClient()
	failOnError(err, "Error creating the client")
	// Update updates a document with a script or partial document.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"doc": map[string]interface{}{
			"title":   "更新你看到外面的世界是什么样的？",
			"content": "更新外面的世界真的很精彩",
		},
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.Update(index, "testtttt", &buf, es.Update.WithDocumentType(documentType))
	if err != nil {
		failOnError(err, "Error Update response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func UpdateByQuery() {
	es, err := GetEsClient()
	failOnError(err, "Error creating the client")
	// UpdateByQuery performs an update on every document in the index without changing the source,
	// for example to pick up a mapping change.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
		// 根据搜索条件更新title
		/*
		   "script": map[string]interface{}{
		       "source": "ctx._source['title']='更新你看到外面的世界是什么样的？'",
		   },
		*/
		// 根据搜索条件更新title、content
		/*
		   "script": map[string]interface{}{
		       "source": "ctx._source=params",
		       "params": map[string]interface{}{
		           "title": "外面的世界真的很精彩",
		           "content": "你看到外面的世界是什么样的？",
		       },
		       "lang": "painless",
		   },
		*/
		// 根据搜索条件更新title、content
		"script": map[string]interface{}{
			"source": "ctx._source.title=params.title;ctx._source.content=params.content;",
			"params": map[string]interface{}{
				"title":   "看看外面的世界真的很精彩",
				"content": "他们和你看到外面的世界是什么样的？",
			},
			"lang": "painless",
		},
	}
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		failOnError(err, "Error encoding doc")
	}
	res, err := es.UpdateByQuery(
		indexArr,
		es.UpdateByQuery.WithDocumentType(documentType),
		es.UpdateByQuery.WithBody(&buf),
		es.UpdateByQuery.WithContext(context.Background()),
		es.UpdateByQuery.WithPretty(),
	)
	if err != nil {
		failOnError(err, "Error Update response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
