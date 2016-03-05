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

package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hoenirvili/RestingGopher/model"
)

// unexported const var
const (
	// server settings
	port   = ":8080"
	domain = "localhost"
	// db settings
	dbName   = "@/gophercloud"
	password = "gopherinthecloud"
	username = "golang:"
)

var (
	// Database export glbal connection
	Database model.DB
	// unexported  err var
	err error
	//Logger global way to log things
	Logger logger
)

// Start func start server process and init all
// config server options
func Start() {
	fmt.Println("[ Logger ]\tInit logger, target: server.log")
	// make new logger global instance
	Logger = newLogger()

	fmt.Println("[ DataBS ]\tStarting mysql with " + username + " pass " + password + " database " + dbName)
	// make new database global instance
	Database, err = model.NewOpen("mysql", username+password+dbName)
	if err != nil {
		Logger.Add("Can't open a new Database Handler")
	}

	fmt.Println("[ Server ]\tStarting on port " + port)
	// start the server
	err = http.ListenAndServe(domain+port, getRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
