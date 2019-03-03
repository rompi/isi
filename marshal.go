package isi

import (
	"encoding/json"
)

// Marshal to JSON value
func Marshal(val interface{}) ([]byte, error) {
	return marshalJSON(val)
}

// MarshalDefault set default value by tag if value is empty
func MarshalDefault(val interface{}) ([]byte, error) {
	setDefaultValue(val)
	return marshalJSON(val)
}

func marshalJSON(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}
