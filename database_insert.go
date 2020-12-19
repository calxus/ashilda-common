package models

import "strings"

type DatabaseInsert struct {
	table string
	columns []string
	values []string
}

func NewDatabaseInsert(t string) *DatabaseInsert {
	return &DatabaseInsert {
		table: t,
		columns: []string{},
		values: []string{},
	}
}

func (di *DatabaseInsert) AddEntry(c string, v string) {
	di.columns = append(di.columns, c)
	di.values = append(di.values, v)
}

func (di *DatabaseInsert) Generate() string {
	statement := INSERT_INTO + SPACE + di.table + SPACE + joinList(di.columns) + SPACE + VALUES + SPACE + joinList(di.values) + SEMI_COLON
	return statement
}

func joinList(l []string) string {
	return LEFT_PARENTHESES + strings.Join(l, COMMA + SPACE) + RIGHT_PARENTHESES
}