package models

import "strings"

type DatabaseSelect struct {
	table string
	columns []string
	conditions []SelectCondition
}

type SelectCondition struct {
	column string
	operator string
	value string
}

func NewDatabaseSelect(t string) *DatabaseSelect {
	return &DatabaseSelect {
		table: t,
		columns: []string{},
		conditions: []SelectCondition{},
	}
}

func (ds *DatabaseSelect) AddColumn(c string) {
	ds.columns = append(ds.columns, c)
}

func (ds *DatabaseSelect) AddCondition(c string, o string, v string) {
	newCondition := SelectCondition {
		column: c,
		operator: o,
		value: v,
	}
	ds.conditions = append(ds.conditions, newCondition)
}

func (ds *DatabaseSelect) Generate() string {
	statement := SELECT + SPACE + ds.generateSelectColumns() + SPACE + FROM + SPACE + ds.table
	if (len(ds.conditions) > 0) {
		statement = statement + SPACE + ds.generateSelectConditions()
	}
	return statement + SEMI_COLON
}

func (ds *DatabaseSelect) generateSelectColumns() string {
	return strings.Join(ds.columns, COMMA + SPACE)
}

func (ds *DatabaseSelect) generateSelectConditions() string {
	conditions := []string{}
	for i := 0; i < len(ds.conditions); i++ {
		conditions = append(conditions, ds.conditions[i].generateConditionStatement())
	}
	return WHERE + SPACE + strings.Join(conditions, SPACE + AND + SPACE)
}

func (c *SelectCondition) generateConditionStatement() string {
	return c.column + SPACE + c.operator + SPACE + APOSTROPHE + c.value + APOSTROPHE
}