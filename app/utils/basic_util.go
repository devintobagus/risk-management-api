package utils

import "encoding/json"

func Parse[Dst any](data any) (*Dst, error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return ParseFromBytes[Dst](raw)

}

func ParseFromBytes[Dst any](src []byte) (*Dst, error) {
	var data_ Dst
	if err := json.Unmarshal(src, &data_); err != nil {
		return nil, err
	}
	return &data_, nil
}
