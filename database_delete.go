package models

import "strings"

type DatabaseDelete struct {
	table string
	conditions []DeleteCondition
}

type DeleteCondition struct {
	column string
	operator string
	value string
}

func NewDatabaseDelete(t string) *DatabaseDelete {
	return &DatabaseDelete {
		table: t,
		conditions: []DeleteCondition{},
	}
}

func (dd *DatabaseDelete) AddCondition(c string, o string, v string) {
	newCondition := DeleteCondition {
		column: c,
		operator: o,
		value: v,
	}
	dd.conditions = append(dd.conditions, newCondition)
}

func (dd *DatabaseDelete) Generate() string {
	statement := DELETE + SPACE + FROM + SPACE + dd.table
	if (len(dd.conditions) > 0) {
		statement = statement + SPACE + dd.generateDeleteConditions()
	}
	return statement + SEMI_COLON
}

func (dd *DatabaseDelete) generateDeleteConditions() string {
	conditions := []string{}
	for i := 0; i < len(dd.conditions); i++ {
		conditions = append(conditions, dd.conditions[i].generateConditionStatement())
	}
	return WHERE + SPACE + strings.Join(conditions, SPACE + AND + SPACE)
}

func (c *DeleteCondition) generateConditionStatement() string {
	return c.column + SPACE + c.operator + SPACE + APOSTROPHE + c.value + APOSTROPHE
}