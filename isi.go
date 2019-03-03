package isi

import "reflect"

func setDefaultValue(val interface{}) {
	// TypeOf returns the reflection Type that represents the dynamic type of variable.
	// If variable is a nil interface value, TypeOf returns nil.
	t := reflect.TypeOf(val).Elem()

	for i := 0; i < t.NumField(); i++ {
		if _, exists := t.Field(i).Tag.Lookup(tagName); !exists {
			continue
		}
		if reflect.ValueOf(val).Elem().Field(i).String() == "" || reflect.ValueOf(val).Elem().Field(i).String() == "0" {
			tagVal := reflect.ValueOf(t.Field(i).Tag.Get(tagName))
			reflect.ValueOf(val).Elem().Field(i).Set(tagVal)
		}
	}
}
