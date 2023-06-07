package imports

import (
	"log"
	"missing"
)

func foo() {
	output, err := missing.Encode()
	if err != nil {
		log.Fatalf("encode error: %v", err)
	}
	log.Print(output)
}
