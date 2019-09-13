package stringtable

import (
	"reflect"
	"testing"
)

var texDataString = "foo\x00bar\x00baz\x00bat"
var texDataData = []int32{
	0,
	4,
	8,
	12,
}

func TestNewTable(t *testing.T) {
	expected := StringTable{}
	actual := *NewFromExistingStringTableData(texDataString, texDataData)
	if reflect.TypeOf(actual) != reflect.TypeOf(expected) {
		t.Errorf("Unexpected type. Expected: %s, actual: %s", reflect.TypeOf(expected), reflect.TypeOf(actual))
	}
}

func TestStringTable_AddOrFindString(t *testing.T) {
	table := NewFromExistingStringTableData(texDataString, texDataData)
	if r := table.AddOrFindString("baz"); r != 2 {
		t.Errorf("Failed to find string. Expected: 2, actual: %d", r)
	}

	if r := table.AddOrFindString("newstring"); r != 4 {
		t.Errorf("Failed to add new string. Expected: 4, actual: %d", r)
	}
}

func TestStringTable_FindString(t *testing.T) {
	table := NewFromExistingStringTableData(texDataString, texDataData)
	if r, _ := table.FindString(0); r != "foo" {
		t.Errorf("Failed to get string. Expected: foo, actual: %s", r)
	}

	if r, _ := table.FindString(3); r != "bat" {
		t.Errorf("Failed to get string. Expected: bat, actual: %s", r)
	}

	if r, _ := table.FindString(999); r != "" {
		t.Errorf("Failed to get string. Expected: '', actual: %s", r)
	}
}
