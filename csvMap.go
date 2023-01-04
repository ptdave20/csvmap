package csvmap

import (
	"reflect"
	"strings"
)

const csvTag = "csvHeader"

func Map(headerRow []string, csvStruct interface{}) {
	v := reflect.ValueOf(csvStruct)
	var t reflect.Type
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	} else {
		t = reflect.TypeOf(csvStruct)
	}

	for fieldIndex := 0; fieldIndex < t.NumField(); fieldIndex++ {
		v.Field(fieldIndex).SetInt(int64(-1))
		field := t.Field(fieldIndex)
		tag := field.Tag.Get(csvTag)

		possibleHeaders := strings.Split(tag, ",")
		if possibleHeaders[0] == "-" {
			continue
		}
		for v := range possibleHeaders {
			possibleHeaders[v] = strings.TrimSpace(possibleHeaders[v])
			possibleHeaders[v] = strings.ToLower(possibleHeaders[v])
		}

		for headerColumn := range headerRow {
			column := strings.ToLower(headerRow[headerColumn])
			column = strings.TrimSpace(column)

			for h := range possibleHeaders {
				if column == possibleHeaders[h] {
					v.Field(fieldIndex).SetInt(int64(headerColumn))
					break
				}
			}

		}
	}
}
