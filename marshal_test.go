package isi_test

import (
	"testing"

	"github.com/rompi/isi"
)

func TestDefaultMarshal(t *testing.T) {
	t.Run("test positive without default value", func(t *testing.T) {
		val := &Example{
			First: "TONO",
		}

		byt, err := isi.Marshal(val)
		if len(byt) == 0 {
			t.Errorf("Byte value")
		}

		if err != nil {
			t.Errorf("Error marshal")
		}

		res := &Example{}
		err = isi.Unmarshal(byt, res)

		if err != nil {
			t.Errorf("Unmarshal json")
		}

		if res.First != "TONO" {
			t.Errorf("Unmarshal first_field")
		}

		if res.Second != "" {
			t.Errorf("Unmarshal default value")
		}
	})

	t.Run("test positive with default value", func(t *testing.T) {
		val := &Example{
			First: "TONO",
		}

		byt, err := isi.MarshalDefault(val)
		if len(byt) == 0 {
			t.Errorf("Byte value")
		}

		if err != nil {
			t.Errorf("Error marshal")
		}

		res := &Example{}
		err = isi.Unmarshal(byt, res)

		if err != nil {
			t.Errorf("Unmarshal json")
		}

		if res.First != "TONO" {
			t.Errorf("Unmarshal first_field")
		}

		if res.Second != "kedua" {
			t.Errorf("Unmarshal default value")
		}
	})
}
