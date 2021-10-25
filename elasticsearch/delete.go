package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func DeleteById() {
	es, err := GetEsClient()
	failOnError(err, "Error creating the client")
	// Delete removes a document from the index
	res, err := es.Delete(index, "POcKSHIBX-ZyL96-ywQO")
	if err != nil {
		failOnError(err, "Error delete by id response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func DeleteByQuery() {
	es, err := GetEsClient()
	failOnError(err, "Error creating the client")
	// DeleteByQuery deletes documents matching the provided query
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}
	res, err := es.DeleteByQuery(indexArr, &buf)
	if err != nil {
		failOnError(err, "Error delete by query response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
