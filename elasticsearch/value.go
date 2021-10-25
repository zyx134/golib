package elasticsearch

import "github.com/elastic/go-elasticsearch/v6"

var (
	host         = []string{"http://diediediedie.asuscomm.com:9200"}
	index        = "elastic-search-test_v1"
	indexArr     = []string{index}
	documentType = "test_v1"
)

func GetEsClient() (es *elasticsearch.Client, err error) {
	config := elasticsearch.Config{
		Addresses: host,
		Username:  "",
		Password:  "",
		CloudID:   "",
		APIKey:    "",
	}
	// new client
	es, err = elasticsearch.NewClient(config)
	return
}
