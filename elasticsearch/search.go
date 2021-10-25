package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

func SearchById() {
	es, err := GetEsClient()
	failOnError(err, "Error creating the client")
	res, err := es.Get(index, "testtttt")
	if err != nil {
		failOnError(err, "Error get response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func SearchByQuery() {
	es, err := GetEsClient()
	failOnError(err, "Error creating the client")
	// info
	res, err := es.Info()
	failOnError(err, "Error getting response")
	// search - highlight
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "世界",
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font color='red'>"},
			"post_tags": []string{"</font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}
	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		failOnError(err, "Error getting response")
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
