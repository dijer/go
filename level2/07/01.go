package main

import (
	"fmt"
	"log"
	"reflect"
)

func PrintStruct(in interface{}) error {
	if in == nil {
		return fmt.Errorf("Empty input")
	}
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("value is not struct")
	}
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		if typeField.Type.Kind() == reflect.Struct {
			log.Printf("nested field: %v", typeField.Name)
			PrintStruct(val.Field(i).Interface())
			continue
		}
	}

	return nil
}
