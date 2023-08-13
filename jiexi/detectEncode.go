package jiexi

import (
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
	"unicode/utf8"
)

func DetectEncoding(content []byte) (encoding.Encoding, error) {
	if utf8.Valid(content) {
		return encoding.Nop, nil
	}
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest(content)
	if err != nil {
		return nil, err
	}

	enc, err := htmlindex.Get(result.Charset)
	if err != nil {
		return nil, err
	}

	return enc, nil
}
