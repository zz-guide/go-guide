package e5

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v5"
	"github.com/elastic/go-elasticsearch/v5/esapi"

	"fmt"
	"os"
)

var es5Client *elasticsearch.Client

func init() {
	var err error
	config := elasticsearch.Config{}
	config.Addresses = []string{"http://baiyin789.top:9200"}
	es5Client, err = elasticsearch.NewClient(config)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CreateIndex() {
	body := map[string]interface{}{
		"mappings": map[string]interface{}{
			"test_type": map[string]interface{}{
				"properties": map[string]interface{}{
					"str": map[string]interface{}{
						"type": "keyword", // 表示这个字段不分词
					},
				},
			},
		},
	}
	jsonBody, _ := json.Marshal(body)
	req := esapi.IndicesCreateRequest{
		Index: "test_index",
		Body:  bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), es5Client)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

func InsertSingle() {
	body := map[string]interface{}{
		"num": 0,
		"v":   0,
		"str": "test",
	}
	jsonBody, _ := json.Marshal(body)

	req := esapi.CreateRequest{ // 如果是esapi.IndexRequest则是插入/替换
		Index:        "test_index",
		DocumentType: "test_type",
		DocumentID:   "test_1",
		Body:         bytes.NewReader(jsonBody),
	}
	res, err := req.Do(context.Background(), es5Client)
	checkError(err)
	defer res.Body.Close()
	fmt.Println(res.String())
}

func SelectBySql() {

}
