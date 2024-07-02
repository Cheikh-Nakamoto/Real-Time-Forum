package database

import "database/sql"

func IsEmpty(db *sql.DB, table,Query string) bool {
	var count uint
	if db != nil && len(table) != 0 {
		err := db.QueryRow(Query + table).Scan(&count)
		if err != nil {
			return false
		}
	}
	return count == 0
}
