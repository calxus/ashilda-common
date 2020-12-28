package models

import "strings"

// DatabaseInsert type represents an SQL insert statement
type DatabaseInsert struct {
	table   string
	columns []string
	values  []string
}

// NewDatabaseInsert method to construct the DatabaseInsert type
func NewDatabaseInsert(t string) *DatabaseInsert {
	return &DatabaseInsert{
		table:   t,
		columns: []string{},
		values:  []string{},
	}
}

// AddEntry method adds a single value to insert into the table
func (di *DatabaseInsert) AddEntry(c string, v string) {
	di.columns = append(di.columns, c)
	di.values = append(di.values, v)
}

// Generate method compiles the query and returns the generated string
func (di *DatabaseInsert) Generate() string {
	statement := InsertInto + Space + di.table + Space + joinTables(di.columns) + Space + Values + Space + joinValues(di.values) + SemiColon
	return statement
}

func joinTables(l []string) string {
	return LeftParentheses + strings.Join(l, Comma+Space) + RightParentheses
}

func joinValues(l []string) string {
	newList := []string{}
	for _, s := range l {
		newList = append(newList, Apostrophe+s+Apostrophe)
	}
	return LeftParentheses + strings.Join(newList, Comma+Space) + RightParentheses
}
