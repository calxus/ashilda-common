package models

import "testing"

func TestNewDatabaseSelect(t *testing.T) {
	ds := NewDatabaseSelect("table_name")
	if (ds.table != "table_name") {
		t.Fail()
	}
}

func TestAddColumn(t *testing.T) {
	ds := NewDatabaseSelect("table_name")
	ds.AddColumn("column_name_1")
	ds.AddColumn("column_name_2")
	if (ds.columns[0] != "column_name_1") {
		t.Fail()
	}
	if (ds.columns[1] != "column_name_2") {
		t.Fail()
	}
}

func TestAddCondition(t *testing.T) {
	ds := NewDatabaseSelect("table_name")
	ds.AddCondition("column", "operator", "value")
	if (ds.conditions[0].column != "column") {
		t.Fail()
	}
	if (ds.conditions[0].operator != "operator") {
		t.Fail()
	}
	if (ds.conditions[0].value != "value") {
		t.Fail()
	}
}

func TestGenerateZeroConditions(t *testing.T) {
	ds := NewDatabaseSelect("table_name")
	ds.AddColumn("column_name")
	if (ds.Generate() != "SELECT column_name FROM table_name;") {
		t.Error(ds.Generate())
	}
}

func TestGenerateOneCondition(t *testing.T) {
	ds := NewDatabaseSelect("table_name")
	ds.AddColumn("column_name")
	ds.AddCondition("column_name", "=", "1")
	if (ds.Generate() != "SELECT column_name FROM table_name WHERE column_name = '1';") {
		t.Error(ds.Generate())
	}
}

func TestGenerateMultipleCondition(t *testing.T) {
	ds := NewDatabaseSelect("table_name")
	ds.AddColumn("column_name_1")
	ds.AddColumn("column_name_2")
	ds.AddCondition("column_name_1", "=", "1")
	ds.AddCondition("column_name_2", "=", "2")
	if (ds.Generate() != "SELECT column_name_1, column_name_2 FROM table_name WHERE column_name_1 = '1' AND column_name_2 = '2';") {
		t.Error(ds.Generate())
	}
}
