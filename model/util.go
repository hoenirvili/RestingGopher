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

type articleComment struct {
	IDComment uint64
	UserName  string
	UserMail  string
}

/*

 */
// ArticleJSON return article JSON format from joining multiple tables Article, Cateogries, Image, Comment, Users.
func ArticleJSON(database DB) ([]byte, error) {
	const (
		firstQuery  = "SELECT art.ID_Article, art.Title, art.Time, art.Author, art.Content, cat.Name FROM Article AS art INNER JOIN Category AS cat ON art.ID_Category = cat.ID_Category ORDER BY art.ID_Article ASC"
		secondQeury = "SELECT art.ID_Article , img.Link FROM Image AS img INNER JOIN ArticleImage AS aimg ON aimg.ID_Image = img.ID_Image INNER JOIN Article as art ON aimg.ID_Article = art.ID_Article ORDER BY art.ID_Article ASC"
		thirdQuery  = "SELECT art.ID_Article, com.Time, com.Content	FROM Comment AS com INNER JOIN ArticleComment acom ON acom.ID_Comment = com.ID_Comment INNER JOIN Article as art on acom.ID_Article = art.ID_Article ORDER BY art.ID_Article ASC "
		fourthQuery = "SELECT com.ID_Comment, usr.Name, usr.Email FROM User AS usr INNER JOIN UserComment ucom ON ucom.ID_User = usr.ID_User INNER JOIN Comment as com ON ucom.ID_Comment = com.ID_Comment ORDER BY com.ID_Comment ASC"
	)
	var (
		article           []Articles
		holder            Articles
		timeHolder        []byte
		imageHolder       Image
		commentHolder     Comment
		artcomSliceHolder []articleComment
		artcomHolder      articleComment
		idHolder          int
		indexes           []int
	)
	//first query merge
	rows, err := database.Query(firstQuery)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on first query from Joining Category with Articles table")}
	}

	for rows.Next() {
		err := rows.Scan(&holder.ID, &holder.Title, &timeHolder, &holder.Author, &holder.Content, &holder.Category)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Article table ")}
		}
		// parse time like a string from byte slice
		holder.Time = string(timeHolder)
		article = append(article, holder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	if err = rows.Close(); err != nil {
		panic(err)
	}

	// second query merge
	rows, err = database.Query(secondQeury)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on second query from joining Image with Artices")}
	}
	for rows.Next() {
		err := rows.Scan(&imageHolder.ID, &imageHolder.Link)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Image,ImageArticle tables")}
		}
		article[imageHolder.ID-1].Image = append(article[imageHolder.ID-1].Image, imageHolder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	if err = rows.Close(); err != nil {
		panic(err)
	}

	// third row
	rows, err = database.Query(thirdQuery)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on third query from joining Article with Comments")}
	}
	for rows.Next() {
		err := rows.Scan(&commentHolder.ID, &timeHolder, &commentHolder.Content)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Comment,ArticleComment tables")}
		}
		commentHolder.Time = string(timeHolder)
		article[commentHolder.ID-1].Comments = append(article[commentHolder.ID-1].Comments, commentHolder)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	if err = rows.Close(); err != nil {
		panic(err)
	}

	// fourth query
	rows, err = database.Query(fourthQuery)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on fourth query from joinging Comment with User")}
	}
	for rows.Next() {
		err := rows.Scan(&artcomHolder.IDComment, &artcomHolder.UserName, &artcomHolder.UserMail)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Comment User tables")}
		}
		artcomSliceHolder = append(artcomSliceHolder, artcomHolder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	if err = rows.Close(); err != nil {
		panic(err)
	}

	// suplimentar query
	rows, err = database.Query("SELECT ID_Article FROM ArticleComment ORDER BY ID_Comment")
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on fourth query from joining Comment with User")}
	}
	for rows.Next() {
		err := rows.Scan(&idHolder)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Comment User tables")}
		}
		indexes = append(indexes, idHolder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	if err = rows.Close(); err != nil {
		panic(err)
	}

	// this is the result from very bad sql-ing
	for i := 0; i < len(article); i++ {
		for j := 0; j < len(indexes); j++ {
			if article[i].ID == indexes[i] {
				for k := 0; k < len(article[i].Comments); k++ {
					for l := 0; l < len(artcomSliceHolder); l++ {
						if article[i].Comments[k].ID == artcomSliceHolder[l].IDComment {
							article[i].Comments[k].User.Name = artcomSliceHolder[l].UserName
							article[i].Comments[k].User.Email = artcomSliceHolder[l].UserMail
						}
					}
				}
			}
		}
	}

	return json.MarshalIndent(NewPayload(article), "", " ")
}

// OneArticleJSON makes query and responds with one article in json format
func OneArticleJSON(database DB, id uint64) ([]byte, error) {
	const (
		firstQuery  = "SELECT art.ID_Article, art.Title, art.Time, art.Author, art.Content, cat.Name FROM Article AS art INNER JOIN Category AS cat ON art.ID_Category = cat.ID_Category WHERE art.ID_Article = ?"
		secondQeury = "SELECT art.ID_Article , img.Link FROM Image AS img INNER JOIN ArticleImage AS aimg ON aimg.ID_Image = img.ID_Image INNER JOIN Article as art ON aimg.ID_Article = art.ID_Article WHERE art.ID_Article = 1"
		thirdQuery  = "SELECT art.ID_Article, com.Time, com.Content FROM Comment AS com INNER JOIN ArticleComment acom ON acom.ID_Comment = com.ID_Comment INNER JOIN Article as art ON acom.ID_Article = art.ID_Article WHERE art.ID_Article = 1;"
		fourthQuery = "SELECT com.ID_Comment, usr.Name, usr.Email FROM User AS usr INNER JOIN UserComment ucom ON ucom.ID_User = usr.ID_User INNER JOIN Comment as com ON ucom.ID_Comment = com.ID_Comment ORDER BY com.ID_Comment ASC"
	)

	var (
		article           Articles
		timeHolder        []byte
		imageHolder       Image
		commentHolder     Comment
		artcomSliceHolder []articleComment
		artcomHolder      articleComment
		idHolder          int
		indexes           []int
	)

	handler := database.Handler()

	//first query merge
	err := handler.QueryRow(firstQuery, id).Scan(&article.ID, &article.Title, &timeHolder, &article.Author, &article.Content, &article.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoContent
		}
		return nil, &ErrSQL{Message: fmt.Sprintf("Error read row from single query %s", firstQuery)}
	}
	article.Time = string(timeHolder)

	// second query merge
	rows, err := database.Query(secondQeury)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on second query from joining Image with Artices")}
	}
	for rows.Next() {
		err := rows.Scan(&imageHolder.ID, &imageHolder.Link)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Image,ImageArticle tables")}
		}
		article.Image = append(article.Image, imageHolder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	if err = rows.Close(); err != nil {
		panic(err)
	}

	// second query merge
	rows, err = database.Query(thirdQuery)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on second query from joining Image with Artices")}
	}
	for rows.Next() {
		err := rows.Scan(&commentHolder.ID, &timeHolder, &commentHolder.Content)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Image,ImageArticle tables")}
		}
		article.Comments = append(article.Comments, commentHolder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	if err = rows.Close(); err != nil {
		panic(err)
	}

	// fourth query
	rows, err = database.Query(fourthQuery)
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on fourth query from joinging Comment with User")}
	}
	for rows.Next() {
		err := rows.Scan(&artcomHolder.IDComment, &artcomHolder.UserName, &artcomHolder.UserMail)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Comment User tables")}
		}
		artcomSliceHolder = append(artcomSliceHolder, artcomHolder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	if err = rows.Close(); err != nil {
		panic(err)
	}

	// suplimentar query
	rows, err = database.Query("SELECT ID_Article FROM ArticleComment ORDER BY ID_Comment")
	if err != nil {
		return nil, &ErrSQL{Message: fmt.Sprintf("Error on fourth query from joining Comment with User")}
	}
	for rows.Next() {
		err := rows.Scan(&idHolder)
		if err != nil {
			return nil, &ErrSQL{Message: fmt.Sprintf("Error row read from Comment User tables")}
		}
		indexes = append(indexes, idHolder)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	if err = rows.Close(); err != nil {
		panic(err)
	}

	// this is the result from very bad sql-ing
	for j := 0; j < len(indexes); j++ {
		if article.ID == indexes[j] {
			for k := 0; k < len(article.Comments); k++ {
				for l := 0; l < len(artcomSliceHolder); l++ {
					if article.Comments[k].ID == artcomSliceHolder[l].IDComment {
						article.Comments[k].User.Name = artcomSliceHolder[l].UserName
						article.Comments[k].User.Email = artcomSliceHolder[l].UserMail
					}
				}
			}
		}
	}
	return json.MarshalIndent(NewPayload(article), "", " ")
}
