package isi_test

import (
	"testing"

	"github.com/rompi/isi"
)

type Example struct {
	First  string `json:"first,omitempty"`
	Second string `json:"second,omitempty" default:"kedua"`
}

func TestDefaultUnmarshalString(t *testing.T) {
	ex := []byte(`{"first":"first_field"}`)
	t.Run("test positive without default", func(t *testing.T) {
		res := &Example{}
		err := isi.Unmarshal(ex, res)

		if err != nil {
			t.Errorf("Unmarshal json")
		}

		if res.First != "first_field" {
			t.Errorf("Unmarshal first_field")
		}

		if res.Second != "" {
			t.Errorf("Unmarshal default value")
		}
	})

	t.Run("test positive with default", func(t *testing.T) {
		res := &Example{}
		err := isi.UnmarshalDefault(ex, res)

		if err != nil {
			t.Errorf("Unmarshal json")
		}

		if res.First != "first_field" {
			t.Errorf("Unmarshal first_field")
		}

		if res.Second != "kedua" {
			t.Errorf("Unmarshal default value")
		}
	})
}
