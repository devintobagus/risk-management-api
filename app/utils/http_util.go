package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func ParseResponse[T any](resp *http.Response) (*T, error) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ParseFromBytes[T](body)
}

func Client() http.Client {
	return http.Client{
		Timeout: 10 * time.Second,
	}
}

func MustJSONReader(v any) *bytes.Reader {
	b, _ := json.Marshal(v)
	return bytes.NewReader(b)
}
