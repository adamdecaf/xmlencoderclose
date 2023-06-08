# xmlencoderclose

[![GoDoc](https://godoc.org/github.com/adamdecaf/xmlencoderclose?status.svg)](https://godoc.org/github.com/adamdecaf/xmlencoderclose)
[![Build Status](https://github.com/adamdecaf/xmlencoderclose/workflows/Go/badge.svg)](https://github.com/adamdecaf/xmlencoderclose/actions)
[![Coverage Status](https://codecov.io/gh/adamdecaf/xmlencoderclose/branch/master/graph/badge.svg)](https://codecov.io/gh/adamdecaf/xmlencoderclose)
[![Go Report Card](https://goreportcard.com/badge/github.com/adamdecaf/xmlencoderclose)](https://goreportcard.com/report/github.com/adamdecaf/xmlencoderclose)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/adamdecaf/xmlencoderclose/master/LICENSE)

xmlencoderclose is a Go linter to check that `encoding/xml.Encoder` type has its `Close()` method called. This linter is similar to [`bodyclose`](https://github.com/timakin/bodyclose) and [`sqlclosecheck`](https://github.com/ryanrolds/sqlclosecheck) with inspiration from [gostaticanalysis/sqlrows](https://github.com/gostaticanalysis/sqlrows)

## Install

```
go get github.com/adamdecaf/xmlencoderclose
```

## Example

Given a few example games where players score goals we can combine those to find their total scores.

```go
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
```

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
