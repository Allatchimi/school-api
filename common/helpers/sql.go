package helpers

import (
	"fmt"
	"strings"
)

func AppendWhereClause(where string, append string) string {
	newWhere := where
	if strings.HasPrefix(where, "WHERE") {
		newWhere = fmt.Sprintf("%s AND (%s)", where, append)
	} else {
		newWhere = fmt.Sprintf("WHERE %s", append)
	}
	return newWhere
}

func AppendSafeWhereClause(
	where,
	clauseFormat string,
	value any,
	args []any,
	index int,
) (string, []any, int) {
	clause := fmt.Sprintf(clauseFormat, index)
	if strings.HasPrefix(where, "WHERE") {
		where = fmt.Sprintf("%s AND (%s)", where, clause)
	} else {
		where = fmt.Sprintf("WHERE %s", clause)
	}
	if value != nil {
		args = append(args, value)
	}
	return where, args, index + 1
}
