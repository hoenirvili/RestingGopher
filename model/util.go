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
	"encoding/json"
	"fmt"
)

// CategoriesJSON readss from rows object and return all the info in JSON frormat
func CategoriesJSON(rows *sql.Rows) ([]byte, error) {
	data := []Categories{}
	holder := Categories{}
	// read row by row
	for rows.Next() {
		// scan into data
		err := rows.Scan(&holder.ID, &holder.Name)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Cateogries table ")}
		}
		data = append(data, holder)
	}

	// check if the row.Next() exited from internal error not EOF
	err := rows.Err()
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error encountered after looping through rows")}
	}

	// put back the connection in the pool
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	return json.MarshalIndent(NewPayload(data), "", " ")
}

// OneCategoryJSON return one row from Category table JSON format
func OneCategoryJSON(databse DB, queryStmt string, args ...interface{}) ([]byte, error) {
	holder := Categories{}
	h := databse.Handler()
	err := h.QueryRow(queryStmt, args...).Scan(&holder.ID, &holder.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoContent
		}
		return nil, &ErrSQL{Message: fmt.Sprintf("Error read row from single query %s", queryStmt)}
	}
	return json.MarshalIndent(NewPayload(holder), "", " ")
}

// ArticleJSON return article JSON format from joining multiple tables Article, Cateogries, Image, Comment, Users.
func ArticleJSON(database DB) ([]byte, error) {
	const (
		firstQuery = "SELECT art.ID_Article, art.Title, art.Time, art.Author, art.Content, cat.Name FROM Article AS art INNER JOIN Category AS cat ON art.ID_Category = cat.ID_Category ORDER BY art.ID_Article ASC"
	)
	var (
		article []Articles
		holder  Articles
	)

	rows, err := database.Query(firstQuery)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on first query from Joining Category with articles table")}
	}

	for rows.Next() {
		err := rows.Scan(&holder.ID, &holder.Title, &holder.Time, &holder.Author, &holder.Content, &holder.Category)
		if err != nil {
			fmt.Println(err)
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Article tables table ")}
		}
		article = append(article, holder)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
	return json.MarshalIndent(NewPayload(article), "", " ")
}
