package models

import "strings"

// DatabaseDelete type represents an SQL delete statement
type DatabaseDelete struct {
	table      string
	conditions []DeleteCondition
}

// DeleteCondition type represents a condition within an SQL statement
type DeleteCondition struct {
	column   string
	operator string
	value    string
}

// NewDatabaseDelete method to construct the DatabaseDelete type
func NewDatabaseDelete(t string) *DatabaseDelete {
	return &DatabaseDelete{
		table:      t,
		conditions: []DeleteCondition{},
	}
}

// AddCondition method to add a single condition to the SQL query
func (dd *DatabaseDelete) AddCondition(c string, o string, v string) {
	newCondition := DeleteCondition{
		column:   c,
		operator: o,
		value:    v,
	}
	dd.conditions = append(dd.conditions, newCondition)
}

// Generate method compiles the query and returns the generated string
func (dd *DatabaseDelete) Generate() string {
	statement := Delete + Space + From + Space + dd.table
	if len(dd.conditions) > 0 {
		statement = statement + Space + dd.generateDeleteConditions()
	}
	return statement + SemiColon
}

func (dd *DatabaseDelete) generateDeleteConditions() string {
	conditions := []string{}
	for i := 0; i < len(dd.conditions); i++ {
		conditions = append(conditions, dd.conditions[i].generateConditionStatement())
	}
	return Where + Space + strings.Join(conditions, Space+And+Space)
}

func (c *DeleteCondition) generateConditionStatement() string {
	return c.column + Space + c.operator + Space + Apostrophe + c.value + Apostrophe
}
