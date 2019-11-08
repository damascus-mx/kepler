package utils

import "encoding/json"

// JSON Struct for JSON encoding/decoding
type JSON struct {
}

type payload struct {
	message string
}

// MarshalMessage Returns a message in JSON object
func (j JSON) MarshalMessage(msg string) []byte {
	msgPayload := payload{msg}

	msgJSON, err := json.Marshal(msgPayload)
	if err != nil {
		panic(err)
	}

	return msgJSON
}
