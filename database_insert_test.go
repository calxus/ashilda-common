package models

import "testing"

func TestNewDatabaseInsert(t *testing.T) {
	di := NewDatabaseInsert("table_name")
	if di.table != "table_name" {
		t.Fail()
	}
}

func TestAddInsertEntry(t *testing.T) {
	di := NewDatabaseInsert("table_name")
	di.AddEntry("column_name_1", "value_name_1")
	di.AddEntry("column_name_2", "value_name_2")
	if di.columns[0] != "column_name_1" {
		t.Fail()
	}
	if di.columns[1] != "column_name_2" {
		t.Fail()
	}
	if di.values[0] != "value_name_1" {
		t.Fail()
	}
	if di.values[1] != "value_name_2" {
		t.Fail()
	}
}

func TestGenerateZeroEntries(t *testing.T) {
	di := NewDatabaseInsert("table_name")
	statement := di.Generate()
	if statement != "INSERT INTO table_name () VALUES ();" {
		t.Error(statement)
	}
}

func TestGenerateOneEntry(t *testing.T) {
	di := NewDatabaseInsert("table_name")
	di.AddEntry("column_name_1", "value_name_1")
	statement := di.Generate()
	if statement != "INSERT INTO table_name (column_name_1) VALUES ('value_name_1');" {
		t.Error(statement)
	}
}

func TestGenerateMultipleEntries(t *testing.T) {
	di := NewDatabaseInsert("table_name")
	di.AddEntry("column_name_1", "value_name_1")
	di.AddEntry("column_name_2", "value_name_2")
	statement := di.Generate()
	if statement != "INSERT INTO table_name (column_name_1, column_name_2) VALUES ('value_name_1', 'value_name_2');" {
		t.Error(statement)
	}
}
