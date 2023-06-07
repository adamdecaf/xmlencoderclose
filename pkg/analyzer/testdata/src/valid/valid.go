package valid

import (
	"encoding/xml"
	"log"
	"os"
)

type Document struct {
	A string `xml:"a"`
}

func valid() {
	// First case
	w := xml.NewEncoder(os.Stdout)
	err := w.Encode(Document{
		A: "abc123",
	})
	w.Close()
	if err != nil {
		log.Fatalf("problem encoding: %v", err)
	}

	// Second case, with defer
	w = xml.NewEncoder(os.Stdout)
	defer w.Close()

	err = w.Encode(Document{
		A: "abc123",
	})
	if err != nil {
		log.Fatalf("problem encoding: %v", err)
	}
}
