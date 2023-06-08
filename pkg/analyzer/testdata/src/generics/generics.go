package generics

import (
	"encoding/xml"
	"log"
	"os"
)

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Document[T signed] struct {
	Value T `xml:"value"`
}

func generic() {
	// int
	w := xml.NewEncoder(os.Stdout)
	err := w.Encode(Document[int]{
		Value: 123,
	})
	w.Close()
	if err != nil {
		log.Fatalf("problem encoding: %v", err)
	}

	// int64
	w = xml.NewEncoder(os.Stdout)
	err = w.Encode(Document[int64]{
		Value: int64(123),
	})
	w.Close()
	if err != nil {
		log.Fatalf("problem encoding: %v", err)
	}
}
