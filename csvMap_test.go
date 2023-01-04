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

	Map(headerRecord, &s)

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

	Map(headerRecord, &s)

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

	Map(headerRecord, &s)

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
