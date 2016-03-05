// Copyright [2016] [hoenir]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"database/sql"
	"fmt"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB struct pointer to sql.DB
// sql.Open return handler for database
// database/sql package manages a pool of connections
// in the background and dosen't open any connections
// util you need them.
// Before making queries we test with handler.Ping() or with
// err, _ = db.Exec("DO 1")
//	if err != nil {
//		panic(err.Error())
//	}
type DB struct {
	handler *sql.DB
}

// NewOpen returns a new db type object
// basically returns a db object to make a connection to
// mysql/maria db
func NewOpen(dbtype, auth string) (DB, error) {

	db, err := sql.Open(dbtype, auth)
	if err != nil {
		return DB{db}, &ErrSQL{Message: fmt.Sprintf("Can't make new open handler to mysql database...")}
	}

	err = db.Ping()
	if err != nil {
		return DB{db}, &ErrSQL{Message: fmt.Sprintf("Connection not set")}
	}

	return DB{db}, err
}

// Handler get *sql.DB handler
func (d DB) Handler() *sql.DB {
	return d.handler
}

// Close the handler
// use with defer stmts
func (d *DB) Close() {
	if err := d.handler.Close(); err != nil {
		panic(err)
	}
}

// Query the databse returing the ptr to sql.Rows object to consume it
func (d DB) Query(queryStmt string, args ...interface{}) (*sql.Rows, error) {
	var (
		rows *sql.Rows
		err  error
	)
	if len(args) > 0 {
		rows, err = d.handler.Query(queryStmt, args)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error query : %s", queryStmt)}
		}
	} else {
		rows, err = d.handler.Query(queryStmt)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error query : %s", queryStmt)}
		}
	}

	return rows, err
}
