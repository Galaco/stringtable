package stringtable

import (
	"errors"
	"strings"
)

const (
	// stringTableNullTerminator string table entries are always null terminated
	stringTableNullTerminator = "\x00"
)

var (
	// ErrorStringIdOutOfBounds occurs when a provided lookup index does not exist
	ErrorStringIdOutOfBounds = errors.New("string id is out of stringtable bounds")
	// ErrorCorruptedStringTableData occurs when the underlying string data mismatches with the lookup table
	ErrorCorruptedStringTableData = errors.New("lookup table includes out of bounds references")
)

// StringTable is a string storage object. It consists of a single
// string that is itself a collection of null-separated strings, as
// well as a lookup table that stores indexes to the start of each stored
// string.
// The purpose is to be able to reference a string by a simple 
// incremental index, regardless of the underlying structure of the string data
type StringTable struct {
	data        string
	lookupTable []int32
}

// FindString finds a string by StringID. StringID comes from StringData, and
// is a lookup to StringTable null escaped strings
func (table *StringTable) FindString(stringID int) (string, error) {
	// Ensure that we can't go out of bounds
	if stringID >= len(table.lookupTable) || stringID < 0 {
		return "", ErrorStringIdOutOfBounds
	}
	if table.lookupTable[stringID] >= int32(len(table.data)) {
		return "", ErrorCorruptedStringTableData
	}
	end := strings.Index(table.data[table.lookupTable[stringID]:], stringTableNullTerminator)
	if end == -1 {
		return table.data[table.lookupTable[stringID]:], nil
	}
	return strings.Split(table.data[table.lookupTable[stringID]:], stringTableNullTerminator)[0], nil
}

// AddString adds a new string to the table and returns 
// its lookup index
func (table *StringTable) AddString(s string) int {
	// @TODO validate this
	// this may be invalid
	table.data += stringTableNullTerminator + s
	outOffset := len(table.data)
	table.lookupTable = append(table.lookupTable, int32(outOffset))

	return len(table.lookupTable) - 1
}

// AddOrFindString adds a new string to the Table, unless it exists
// Returns the stringID of the newly added string, or the existing one if found
func (table *StringTable) AddOrFindString(s string) int {
	for i := 0; i < len(table.lookupTable); i++ {
		end := strings.Index(table.data[table.lookupTable[i]:], stringTableNullTerminator)
		if end > 0 {
			if strings.EqualFold(s, table.data[table.lookupTable[i]:int(table.lookupTable[i])+end]) {
				return i
			}
		} else {
			if strings.EqualFold(s, table.data[table.lookupTable[i]:]) {
				return i
			}
		}
	}

	return table.AddString(s)
}

// NewFromExistingStringTableData creates a stringtable from existing stringtable data.
// stringdata is a single null-terminator separated string.
// lookupTable is a slice of indexes to the start of each string in the stringdata
// string.
func NewFromExistingStringTableData(stringdata string, lookupTable []int32) *StringTable {
	return &StringTable{
		data:        stringdata,
		lookupTable: lookupTable,
	}
}

// New creates a new Stringtable
func New() *StringTable {
	return &StringTable{
		data:        "",
		lookupTable: make([]int32, 0),
	}
}