package models

import "testing"

func TestNewDatabaseRecord(t *testing.T) {
	dr := NewDatabaseRecord("table_name")
	if (dr.table != "table_name") {
		t.Fail()
	}
}

func TestAddRecordEntry(t *testing.T) {
	dr := NewDatabaseRecord("table_name")
	dr.AddEntry("column_name", "value_name")
	if (dr.entries[0].column != "column_name") {
		t.Fail()
	}
	if (dr.entries[0].value != "value_name") {
		t.Fail()
	}
}