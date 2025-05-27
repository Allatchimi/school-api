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
