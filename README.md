# CSV Map

CSV Map is a Go library that provides a simple way to map CSV headers to struct fields. It supports custom tags for struct fields to specify CSV headers and options.

## Features

- Map CSV headers to struct fields using `csv` tag.
- Support for optional CSV headers using `csvOption` tag.
- Error handling for required fields and non-integer fields.

## Usage

Import the `csvmap` package in your Go file:

```go
import "csvmap"
```

Define a struct with `csv` and `csvOption` tags:

```go
type MyStruct struct {
	A int `csv:"a" csvOption:"required"`
	B int `csv:"B"`
	C int `csv:"C "`
	D int `csv:"-"`
}
```

Use the Map function to map CSV headers to struct fields:

```go
headerRecord := []string{"123", "a", "B", " c", "d"}
s := MyStruct{}

err := csvmap.Map(headerRecord, &s)
if err != nil {
	// handle error
}
```

# Testing
This project includes test for the Map function. To run the tests, use the go test command:
```
go test
```

# Contributing
Contributions are welcome, Please submit a pull request or create an issue to discuss the changes.

