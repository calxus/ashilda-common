package models

import "strings"

type DatabaseUpdate struct {
	table string
	statements []UpdateStatement
	conditions []UpdateCondition
}

type UpdateStatement struct {
	column string
	value string
}

type UpdateCondition struct {
	column string
	operator string
	value string
}

func NewDatabaseUpdate(t string) *DatabaseUpdate {
	return &DatabaseUpdate {
		table: t,
		statements: []UpdateStatement{},
		conditions: []UpdateCondition{},
	}
}

func (du *DatabaseUpdate) AddStatement(c string, v string) {
	newStatement := UpdateStatement {
		column: c,
		value: v,
	}
	du.statements = append(du.statements, newStatement)
}

func (du *DatabaseUpdate) AddCondition(c string, o string, v string) {
	newCondition := UpdateCondition {
		column: c,
		operator: o,
		value: v,
	}
	du.conditions = append(du.conditions, newCondition)
}

func (du *DatabaseUpdate) Generate() string {
	statement := UPDATE + SPACE + du.table + SPACE + SET
	if (len(du.statements) > 0) {
		statement = statement + SPACE + du.generateUpdateStatements()
	}
	if (len(du.conditions) > 0) {
		statement = statement + SPACE + du.generateUpdateConditions()
	}
	return statement + SEMI_COLON
}

func (du *DatabaseUpdate) generateUpdateStatements() string {
	statements := []string{}
	for i := 0; i < len(du.statements); i++ {
		statements = append(statements, du.statements[i].generateUpdateStatement())
	}
	return strings.Join(statements, COMMA + SPACE)
}

func (s *UpdateStatement) generateUpdateStatement() string {
	return s.column + SPACE + EQUALS + SPACE + APOSTROPHE + s.value + APOSTROPHE
}

func (du *DatabaseUpdate) generateUpdateConditions() string {
	conditions := []string{}
	for i := 0; i < len(du.conditions); i++ {
		conditions = append(conditions, du.conditions[i].generateConditionStatement())
	}
	return WHERE + SPACE + strings.Join(conditions, SPACE + AND + SPACE)
}

func (c *UpdateCondition) generateConditionStatement() string {
	return c.column + SPACE + c.operator + SPACE + APOSTROPHE + c.value + APOSTROPHE
}