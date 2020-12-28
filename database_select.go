package models

import "strings"

// DatabaseSelect type represents an SQL select statement
type DatabaseSelect struct {
	table      string
	columns    []string
	conditions []SelectCondition
}

// SelectCondition type represents a single SQL condition statement
type SelectCondition struct {
	column   string
	operator string
	value    string
}

// NewDatabaseSelect method to construct the DatabaseSelect type
func NewDatabaseSelect(t string) *DatabaseSelect {
	return &DatabaseSelect{
		table:      t,
		columns:    []string{},
		conditions: []SelectCondition{},
	}
}

// AddColumn method adds a single column to return from SQL statement
func (ds *DatabaseSelect) AddColumn(c string) {
	ds.columns = append(ds.columns, c)
}

// AddCondition method to add a single condition to the SQL query
func (ds *DatabaseSelect) AddCondition(c string, o string, v string) {
	newCondition := SelectCondition{
		column:   c,
		operator: o,
		value:    v,
	}
	ds.conditions = append(ds.conditions, newCondition)
}

// Generate method compiles the query and returns the generated string
func (ds *DatabaseSelect) Generate() string {
	statement := Select + Space + ds.generateSelectColumns() + Space + From + Space + ds.table
	if len(ds.conditions) > 0 {
		statement = statement + Space + ds.generateSelectConditions()
	}
	return statement + SemiColon
}

func (ds *DatabaseSelect) generateSelectColumns() string {
	return strings.Join(ds.columns, Comma+Space)
}

func (ds *DatabaseSelect) generateSelectConditions() string {
	conditions := []string{}
	for i := 0; i < len(ds.conditions); i++ {
		conditions = append(conditions, ds.conditions[i].generateConditionStatement())
	}
	return Where + Space + strings.Join(conditions, Space+And+Space)
}

func (c *SelectCondition) generateConditionStatement() string {
	return c.column + Space + c.operator + Space + Apostrophe + c.value + Apostrophe
}
