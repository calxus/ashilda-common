package models

import "testing"

func TestNewDatabaseUpdate(t *testing.T) {
	du := NewDatabaseUpdate("table_name")
	if (du.table != "table_name") {
		t.Fail()
	}
}

func TestAddUpdateStatement(t *testing.T) {
	du := NewDatabaseUpdate("table_name")
	du.AddStatement("column", "value")
	if (du.statements[0].column != "column") {
		t.Fail()
	}
	if (du.statements[0].value != "value") {
		t.Fail()
	}
}

func TestAddUpdateCondition(t *testing.T) {
	du := NewDatabaseUpdate("table_name")
	du.AddCondition("column", "operator", "value")
	if (du.conditions[0].column != "column") {
		t.Fail()
	}
	if (du.conditions[0].operator != "operator") {
		t.Fail()
	}
	if (du.conditions[0].value != "value") {
		t.Fail()
	}
}

func TestUpdateGenerateZeroConditions(t *testing.T) {
	du := NewDatabaseUpdate("table_name")
	du.AddStatement("column_name", "value_name")
	if (du.Generate() != "UPDATE table_name SET column_name = 'value_name';") {
		t.Error(du.Generate())
	}
}

func TestUpdateGenerateOneCondition(t *testing.T) {
	du := NewDatabaseUpdate("table_name")
	du.AddStatement("column_name", "value_name")
	du.AddCondition("column_name", "=", "1")
	if (du.Generate() != "UPDATE table_name SET column_name = 'value_name' WHERE column_name = '1';") {
		t.Error(du.Generate())
	}
}

func TestUpdateGenerateMultipleConditions(t *testing.T) {
	du := NewDatabaseUpdate("table_name")
	du.AddStatement("column_name_1", "value_name_1")
	du.AddStatement("column_name_2", "value_name_2")
	du.AddCondition("column_name_1", "=", "1")
	du.AddCondition("column_name_2", "=", "1")
	if (du.Generate() != "UPDATE table_name SET column_name_1 = 'value_name_1', column_name_2 = 'value_name_2' WHERE column_name_1 = '1' AND column_name_2 = '1';") {
		t.Error(du.Generate())
	}
}
