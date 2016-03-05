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
	if err := rows.Err(); err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error encountered after looping through rows")}
	}

	// put back the connection in the pool
	if err := rows.Close(); err != nil {
		panic(err)
	}

	return json.MarshalIndent(NewPayload(data), "", " ")
}

// OneCategoryJSON return one row from Category table JSON format
func OneCategoryJSON(databse DB, queryStmt string, arg interface{}) ([]byte, error) {
	holder := Categories{}
	h := databse.Handler()
	err := h.QueryRow(queryStmt, arg).Scan(&holder.ID, &holder.Name)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error read row from single query %s", queryStmt)}
	}
	return json.MarshalIndent(NewPayload(holder), "", " ")
}
