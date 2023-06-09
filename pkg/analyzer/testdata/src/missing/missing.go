package missing

import (
	"bytes"
	"encoding/xml"
)

type Document struct {
	A string `xml:"a"`
}

func Encode() (string, error) {
	var buf bytes.Buffer
	err := xml.NewEncoder(&buf).Encode(Document{ // want "Encoder.Close must be called"
		A: "abc123",
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
