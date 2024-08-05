package dashvector

import (
	"encoding/json"
	"fmt"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
)

type InsertRequest struct {
	Documents []Documents `json:"docs"`
}

type Documents struct {
	Vector []float64 `json:"vector"`
	Fields Fields    `json:"fields"`
}

type Fields struct {
	OriginQuestion string `json:"originQuestion"`
	Content        string `json:"content"`
}

type SearchResponse struct {
	Status    int                    `json:"code"`
	RequestId string                 `json:"request_id"`
	Message   string                 `json:"message"`
	Output    []SearchResponseOutput `json:"output"`
}

type SearchResponseOutput struct {
	ID     string    `json:"id"`
	Score  float64   `json:"score"`
	Fields Fields    `json:"fields"`
	Vector []float64 `json:"vector"`
}

type SearchRequest struct {
	Vector        []float64 `json:"vector"`
	TopK          int       `json:"topk"`
	IncludeVector bool      `json:"include_vector"`
}

func QueryVectorResponse(responseBody []byte, log wrapper.Log) (*SearchResponse, error) {
	var response SearchResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Errorf("[QueryNearestVectorResponse]Unmarshal json error:%s, response:%s.", err, string(responseBody))
		return nil, err
	}
	return &response, nil
}

func GenerateQueryNearestVectorRequest(c string, k string, vector []float64, log wrapper.Log) (string, []byte, [][2]string, error) {
	url := fmt.Sprintf("/v1/collections/%s/query", c)

	requestData := SearchRequest{
		Vector:        vector,
		TopK:          1,
		IncludeVector: true,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Errorf("Marshal json error:%s, data:%s.", err, requestData)
		return "", nil, nil, err
	}

	header := [][2]string{
		{"Content-Type", "application/json"},
		{"dashvector-auth-token", k},
	}

	return url, requestBody, header, nil
}

func GenerateInsertDocumentsRequest(c string, k string, fields Fields, vector []float64, log wrapper.Log) (string, []byte, [][2]string, error) {
	url := fmt.Sprintf("/v1/collections/%s/docs", c)

	DocumentsObject := Documents{
		Fields: fields,
		Vector: vector,
	}

	requestData := InsertRequest{
		Documents: []Documents{DocumentsObject},
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Errorf("Marshal json error:%s, data:%s.", err, requestData)
		return "", nil, nil, err
	}

	header := [][2]string{
		{"Content-Type", "application/json"},
		{"dashvector-auth-token", k},
	}

	return url, requestBody, header, nil
}
