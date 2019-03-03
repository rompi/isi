package isi

import (
	"encoding/json"
)

const tagName = "default"

// Unmarshal with JSON
func Unmarshal(data []byte, v interface{}) error {
	return unmarshalJSON(data, &v)
}

// UnmarshalDefault set default value by tag if value is empty
func UnmarshalDefault(data []byte, val interface{}) error {
	err := unmarshalJSON(data, &val)
	setDefaultValue(val)
	return err
}

func unmarshalJSON(data []byte, val interface{}) error {
	err := json.Unmarshal(data, &val)
	return err
}
