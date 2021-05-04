package tokens

import "strings"

type Header struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (h *Header) Set(key, value string) {
	parts := strings.Split(value, " ")
	if len(parts) > 1 {
		h.Key = key
		h.Type = parts[0]
		h.Value = parts[1]
	}
}

