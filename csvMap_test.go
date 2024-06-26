package csvmap

import "testing"

type testStruct struct {
	A int `csv:"a"`
	B int `csv:"B"`
	C int `csv:"C "`
	D int `csv:"-"`
}

func TestMapOne(t *testing.T) {
	s := testStruct{}

	headerRecord := []string{"123", "A", "b", " c", "d"}

	if err := Map(headerRecord, &s); err != nil {
		t.Fatalf("Error mapping: %v", err)
		return
	}

	if s.A != 1 {
		t.Fatalf("A should be 1 and was %d", s.A)
		return
	}
	if s.B != 2 {
		t.Fatalf("B should be 2 and was %d", s.B)
		return
	}
	if s.C != 3 {
		t.Fatalf("C should be 3 and was %d", s.C)
		return
	}
	if s.D != -1 {
		t.Fatalf("D should be -1 and was %d", s.D)
		return
	}
}

func TestMapTwo(t *testing.T) {
	s := testStruct{}

	headerRecord := []string{"c", "d", "a", "%", "b"}

	if err := Map(headerRecord, &s); err != nil {
		t.Fatalf("Error mapping: %v", err)
		return
	}

	if s.A != 2 {
		t.Fatalf("A should be 2 and was %d", s.A)
		return
	}
	if s.B != 4 {
		t.Fatalf("B should be 4 and was %d", s.B)
		return
	}
	if s.C != 0 {
		t.Fatalf("C should be 0 and was %d", s.C)
		return
	}
	if s.D != -1 {
		t.Fatalf("D should be -1 and was %d", s.D)
		return
	}
}

type testMStruct struct {
	A int `csv:"a,x"`
	B int `csv:"B,y"`
	C int `csv:"C,z "`
	D int `csv:"-,w"`
}

func TestMapThree(t *testing.T) {
	s := testMStruct{}

	headerRecord := []string{"x", "y", "z", "w"}

	if err := Map(headerRecord, &s); err != nil {
		t.Fatalf("Error mapping: %v", err)
		return
	}

	if s.A != 0 {
		t.Fatalf("A should be 2 and was %d", s.A)
		return
	}
	if s.B != 1 {
		t.Fatalf("B should be 4 and was %d", s.B)
		return
	}
	if s.C != 2 {
		t.Fatalf("C should be 0 and was %d", s.C)
		return
	}
	if s.D != -1 {
		t.Fatalf("D should be -1 and was %d", s.D)
		return
	}
}

type testRequiredStruct struct {
	A int `csv:"a" csvOption:"required"`
	B int `csv:"B"`
	C int `csv:"C "`
	D int `csv:"-"`
}

func TestMapRequired(t *testing.T) {
	s := testRequiredStruct{}

	headerRecord := []string{"123", "B", " c", "d"}

	err := Map(headerRecord, &s)
	if err == nil || err.Error() != "required field a not found in header row" {
		t.Fatalf("Expected error for required field not found, got %v", err)
	}

	headerRecord = []string{"123", "a", "B", " c", "d"}

	err = Map(headerRecord, &s)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

type testIntStruct struct {
	A int     `csv:"a"`
	B float64 `csv:"B"`
	C string  `csv:"C "`
	D int     `csv:"-"`
}

func TestMapInt(t *testing.T) {
	s := testIntStruct{}

	headerRecord := []string{"123", "a", "B", " c", "d"}

	err := Map(headerRecord, &s)
	if err == nil || err.Error() != "field B is not of type int" {
		t.Fatalf("Expected error for non-int field, got %v", err)
	}

	s = testIntStruct{}
	headerRecord = []string{"123", "a", " c", "d"}

	err = Map(headerRecord, &s)
	if err == nil || err.Error() != "field B is not of type int" {
		t.Fatalf("Expected error for non-int field, got %v", err)
	}
}
