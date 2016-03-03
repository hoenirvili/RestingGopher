package models

import (
	"database/sql"
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
	return DB{db}, err
}

// Close the handler
// use with defer stmts
func (d *DB) Close() {
	if err := d.handler.Close(); err != nil {
		//TODO better way to handle unexpected error on close
		panic(err)
	}
}

// Query the databse returing in a serialized json format
func (d DB) Query(queryStmt string, data interface{}) {
	// rows, err := d.handler.Query(queryStmt)
	// if err != nil {
	//
	// }
}

// func StartConnSql() error {
// 	db, err := sql.Open("mysql", username+password+dbName)
// 	if err != nil {
// 		return &ErrSql{fmt.Sprintf("%s\n", err)}
// 	}
// 	// Validate connection
// 	err = db.Ping()
// 	if err != nil {
// 		return &ErrSql{fmt.Sprintf("%s\n", err)}
// 	}
//
// 	defer func() {
// 		if err = db.Close(); err != nil {
// 			panic(err)
// 		}
// 	}()
//
// 	prep, errPrep := db.Query("SELECT *FROM Category")
//
// 	if errPrep != nil {
// 		panic(errPrep)
// 	}
//
// 	for prep.Next() {
// 		var idCateg int
// 		var name string
// 		err = prep.Scan(&idCateg, &name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(idCateg)
// 		fmt.Println(name)
// 	}
//
// 	return nil
//
// }
