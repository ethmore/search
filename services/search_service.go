package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"search-service/dotEnv"
	"time"
)

type QueryString struct {
	Query string `json:"query"`
}

type Query struct {
	QueryString QueryString `json:"query_string"`
}

type SearchQuery struct {
	Query Query `json:"query"`
}

type ResultShards struct {
	Total      int
	Successful int
	Skipped    int
	Failed     int
}

type Source struct {
	ID          string
	Title       string
	Price       string
	Description string
	Image       string
	Stock       string
}

type Total struct {
	Value    int
	Relation string
}

type Hits struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float32 `json:"_score"`
	Source Source  `json:"_source"`
}

type HitsParent struct {
	Total    Total
	MaxScore float32 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}

type SearchResult struct {
	Took     int          `json:"took"`
	TimedOut bool         `json:"timed_out"`
	Shards   ResultShards `json:"_shards"`
	Hits     HitsParent   `json:"hits"`
}

func Search(keyword string) ([]Source, error) {
	var query SearchQuery
	query.Query.QueryString.Query = keyword
	body, _ := json.Marshal(query)
	jsonBody := []byte(body)
	bodyReader := bytes.NewReader(jsonBody)
	requestURL := dotEnv.GoDotEnvVariable("SEARCH_PRODUCT")

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

	var response SearchResult
	if unmarshalErr := json.Unmarshal([]byte(b), &response); err != nil {
		return nil, unmarshalErr
	}

	var products []Source
	for _, hits := range response.Hits.Hits {
		products = append(products, hits.Source)
	}

	fmt.Println(products)
	return products, nil
}
