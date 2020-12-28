package models

import "testing"

func TestNewDatabaseDelete(t *testing.T) {
	dd := NewDatabaseDelete("table_name")
	if dd.table != "table_name" {
		t.Fail()
	}
}

func TestAddDeleteCondition(t *testing.T) {
	dd := NewDatabaseDelete("table_name")
	dd.AddCondition("column", "operator", "value")
	if dd.conditions[0].column != "column" {
		t.Fail()
	}
	if dd.conditions[0].operator != "operator" {
		t.Fail()
	}
	if dd.conditions[0].value != "value" {
		t.Fail()
	}
}

func TestDeleteGenerateZeroConditions(t *testing.T) {
	dd := NewDatabaseDelete("table_name")
	if dd.Generate() != "DELETE FROM table_name;" {
		t.Error(dd.Generate())
	}
}

func TestDeleteGenerateOneCondition(t *testing.T) {
	dd := NewDatabaseDelete("table_name")
	dd.AddCondition("column_name", "=", "1")
	if dd.Generate() != "DELETE FROM table_name WHERE column_name = '1';" {
		t.Error(dd.Generate())
	}
}

func TestDeleteGenerateMultipleCondition(t *testing.T) {
	dd := NewDatabaseDelete("table_name")
	dd.AddCondition("column_name_1", "=", "1")
	dd.AddCondition("column_name_2", "=", "2")
	if dd.Generate() != "DELETE FROM table_name WHERE column_name_1 = '1' AND column_name_2 = '2';" {
		t.Error(dd.Generate())
	}
}
