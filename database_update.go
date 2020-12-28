package models

import "strings"

// DatabaseUpdate type represents an SQL update statement
type DatabaseUpdate struct {
	table      string
	statements []UpdateStatement
	conditions []UpdateCondition
}

// UpdateStatement type represents a single update within an SQL statement
type UpdateStatement struct {
	column string
	value  string
}

// UpdateCondition type represents a condition within an SQL statement
type UpdateCondition struct {
	column   string
	operator string
	value    string
}

// NewDatabaseUpdate method to construct the DatabaseUpdate type
func NewDatabaseUpdate(t string) *DatabaseUpdate {
	return &DatabaseUpdate{
		table:      t,
		statements: []UpdateStatement{},
		conditions: []UpdateCondition{},
	}
}

// AddStatement method to add a single update statement to SQL query
func (du *DatabaseUpdate) AddStatement(c string, v string) {
	newStatement := UpdateStatement{
		column: c,
		value:  v,
	}
	du.statements = append(du.statements, newStatement)
}

// AddCondition method to add a single condition to the SQL query
func (du *DatabaseUpdate) AddCondition(c string, o string, v string) {
	newCondition := UpdateCondition{
		column:   c,
		operator: o,
		value:    v,
	}
	du.conditions = append(du.conditions, newCondition)
}

// Generate method compiles the query and returns the generated string
func (du *DatabaseUpdate) Generate() string {
	statement := Update + Space + du.table + Space + Set
	if len(du.statements) > 0 {
		statement = statement + Space + du.generateUpdateStatements()
	}
	if len(du.conditions) > 0 {
		statement = statement + Space + du.generateUpdateConditions()
	}
	return statement + SemiColon
}

func (du *DatabaseUpdate) generateUpdateStatements() string {
	statements := []string{}
	for i := 0; i < len(du.statements); i++ {
		statements = append(statements, du.statements[i].generateUpdateStatement())
	}
	return strings.Join(statements, Comma+Space)
}

func (s *UpdateStatement) generateUpdateStatement() string {
	return s.column + Space + Equals + Space + Apostrophe + s.value + Apostrophe
}

func (du *DatabaseUpdate) generateUpdateConditions() string {
	conditions := []string{}
	for i := 0; i < len(du.conditions); i++ {
		conditions = append(conditions, du.conditions[i].generateConditionStatement())
	}
	return Where + Space + strings.Join(conditions, Space+And+Space)
}

func (c *UpdateCondition) generateConditionStatement() string {
	return c.column + Space + c.operator + Space + Apostrophe + c.value + Apostrophe
}
