package models

type DatabaseRecord struct {
	table string
	entries []DatabaseEntry
}

type DatabaseEntry struct {
	column string
	value string
}

func NewDatabaseRecord(t string) *DatabaseRecord {
	return &DatabaseRecord {
		table: t,
		entries: []DatabaseEntry{},
	}
}

func (dr *DatabaseRecord) AddEntry(c string, v string) {
	entry := DatabaseEntry {
		column: c,
		value: v,
	}
	dr.entries = append(dr.entries, entry)
}