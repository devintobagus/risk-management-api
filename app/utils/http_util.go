package utils

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

func ParseResponse[T any](resp *http.Response) (*T, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ParseFromBytes[T](body)
}

func ParseXmlResponse[T any](resp *http.Response) (*T, error) {
	var data T
	err := xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
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
