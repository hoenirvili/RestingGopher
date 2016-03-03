package models

import "fmt"

// ErrSql struct for internal purpose error handleing
type ErrSql struct {
	Message string
}

//Errorimplement error interface
func (e ErrSql) Error() string {
	return fmt.Sprintf("[SQL ERROR] %s\n", e.Message)
}
