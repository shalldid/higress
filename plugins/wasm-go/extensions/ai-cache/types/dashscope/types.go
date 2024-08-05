package dashscope

import (
	"encoding/json"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
)

type EmbeddingRequest struct {
	Model      string     `json:"Model"`
	Input      string     `json:"input"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	TextType string `json:"text_type"`
}

type EmbeddingResponse struct {
	RequestId string                  `json:"request_id"`
	Usage     Usage                   `json:"usage"`
	Output    EmbeddingResponseOutput `json:"output"`
}

type Usage struct {
	TotalTokens int `json:"total_tokens"`
}

type EmbeddingResponseOutput struct {
	Embeddings []EmbeddingData `json:"embeddings"`
}

type EmbeddingData struct {
	Embedding []float64 `json:"embedding"`
	TextIndex int       `json:"text_index"`
}

func GenerateTextEmbeddingsRequest(texts string, log wrapper.Log) (string, []byte, [][2]string) {
	url := "/api/v1/services/embeddings/text-embedding"

	data := EmbeddingRequest{
		Model: "zpoint",
		Input: texts,
		Parameters: Parameters{
			TextType: "query",
		},
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

func TextEmbeddingsVectorResponse(responseBody []byte, log wrapper.Log) (*EmbeddingResponse, error) {
	var response EmbeddingResponse
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Errorf("[TextEmbeddingsVectorResponse] Unmarshal json error:%s, response:%s.", err, string(responseBody))
		return nil, err
	}
	return &response, nil
}
