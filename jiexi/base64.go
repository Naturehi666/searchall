package jiexi

import "encoding/base64"

func ProcessFastCodeHistroy(value string) (string, error) {
	decodedValue, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}
	return string(decodedValue), nil
}
