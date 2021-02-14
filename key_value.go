package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// KeyValue stores a collection of key value pairs.
type KeyValue map[string]string

// Find tries to return a value from a keyvalue.
func (k KeyValue) Find(key string) string {
	return k[key]
}

// Set the value for given key in KeyValue
func (k KeyValue) Set(key, value string) {
	k[key] = value
}

// NewKeyValue creates a keyvalue from JSON.
func NewKeyValue(rdr io.Reader) (KeyValue, error) {
	var keyvalue map[string]string
	err := json.NewDecoder(rdr).Decode(&keyvalue)

	if err != nil {
		err = fmt.Errorf("problem parsing keyvalue, %v", err)
	}

	return keyvalue, err
}
