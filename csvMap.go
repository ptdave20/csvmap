package csvmap

import (
	"errors"
	"reflect"
	"strings"
)

const csvTag = "csv"
const csvOptionTag = "csvOption"

func Map(headerRow []string, csvStruct interface{}) error {
	v := reflect.ValueOf(csvStruct)
	var t reflect.Type
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	} else {
		t = reflect.TypeOf(csvStruct)
	}

	for fieldIndex := 0; fieldIndex < t.NumField(); fieldIndex++ {
		field := t.Field(fieldIndex)
		tag := field.Tag.Get(csvTag)
		optionTag := field.Tag.Get(csvOptionTag)

		if v.Field(fieldIndex).Kind() != reflect.Int {
			return errors.New("field " + tag + " is not of type int")
		}

		v.Field(fieldIndex).SetInt(int64(-1))

		possibleHeaders := strings.Split(tag, ",")
		if possibleHeaders[0] == "-" {
			continue
		}
		for v := range possibleHeaders {
			possibleHeaders[v] = strings.TrimSpace(possibleHeaders[v])
			possibleHeaders[v] = strings.ToLower(possibleHeaders[v])
		}

		found := false
		for headerColumn := range headerRow {
			column := strings.ToLower(headerRow[headerColumn])
			column = strings.TrimSpace(column)

			for h := range possibleHeaders {
				if column == possibleHeaders[h] {
					v.Field(fieldIndex).SetInt(int64(headerColumn))
					found = true
					break
				}
			}
		}

		if optionTag == "required" && !found {
			return errors.New("required field " + tag + " not found in header row")
		}
	}
	return nil
}
