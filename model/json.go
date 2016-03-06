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

//Payload json fromat
type Payload struct {
	Data interface{}
}

//NewPayload return new JSON response payload
func NewPayload(d interface{}) Payload {
	return Payload{
		Data: d,
	}
}

// Categories struct serializable json response
type Categories struct {
	ID   int    `json : "ID"`
	Name string `json : "Name"`
}

// Image struct serializable json response
type Image struct {
	ID   int    `json : "ID"`
	Link string `json : "Link"`
}

// Comment struct serializable json response
type Comment struct {
	ID      int    `json : "ID"`
	Time    string `json : "Time"`
	Content string `json : "Content"`
}

// User struct serializable json response
type User struct {
	ID    int    `json: "ID"`
	Name  string `json: "Name"`
	Email string `json: "Email:"`
}

// Articles struct serializable json response
type Articles struct {
	ID       int        `json: "ID"`
	Title    string     `json: "Title"`
	Time     []byte     `json: "Time"`
	Author   string     `json: "Author"`
	Content  string     `json: "Content"`
	Category Categories `json: "Category"`
	Image    []Image    `json: "Image"`
	Comments []Comment  `json "Comments"`
}
