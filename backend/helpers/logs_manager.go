package helpers

import (
	"database/sql"
)

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func insertLog(exec Executor, query string, args ...interface{}) error {
	_, err := exec.Exec(query, args...)
	return err
}

func InsertLogsError(exec Executor, tableName string, message string) error {
	query := `
        INSERT INTO logs_error (table_name, message)
        VALUES ($1, $2)`
	return insertLog(exec, query, tableName, message)
}

func InsertLogs(exec Executor, action string, tableName string, recordId int, description string) error {
	query := `
        INSERT INTO log_actions (action, table_name, record_id, description)
        VALUES ($1, $2, $3, $4)`
	return insertLog(exec, query, action, tableName, recordId, description)
}
