package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"search/dotEnv"
	"time"
)

type Shard struct {
	Total      int
	Successful int
	Failed     int
}

type ElasticsearchResponse struct {
	Index      string `json:"_index"`
	Type       string `json:"_type"`
	Id         string `json:"_id"`
	Version    int    `json:"_version"`
	Result     string
	Shards     Shard `json:"_shards"`
	SeqNo      int   `json:"_seq_no"`
	PrimarTerm int   `json:"_primary_term"`
}

func AddProduct(p Product) (*ElasticsearchResponse, error) {
	body, _ := json.Marshal(p)
	jsonBody := []byte(body)
	bodyReader := bytes.NewReader(jsonBody)
	requestURL := dotEnv.GoDotEnvVariable("ADD_PRODUCT")

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}

	//Read response
	b, readErr := io.ReadAll(res.Body)
	if err != nil {
		return nil, readErr
	}
	defer res.Body.Close()

	var response ElasticsearchResponse
	if unmarshalErr := json.Unmarshal([]byte(b), &response); err != nil {
		return nil, unmarshalErr
	}
	fmt.Println(response)
	return &response, nil
}
