package zyecs

import (
	"encoding/json"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
)

type QueryRequest struct {
	Model string `json:"Model"`
	Input string `json:"input"`
}

func GenerateQueryRequest(texts string, log wrapper.Log) (string, []byte, [][2]string) {
	url := "/api/v1/services/embeddings/search"

	data := QueryRequest{
		Model: "zpoint",
		Input: texts,
	}

	requestBody, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Marshal json error:%s, data:%s.", err, data)
		return "", nil, nil
	}

	headers := [][2]string{
		{"Content-Type", "application/json"},
	}
	return url, requestBody, headers
}
