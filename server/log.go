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
	"os"
	"time"
)

// Logger struct has methods for logging warnings and erros
// while the app is runing
type logger struct {
	file *os.File
}

// NewLogger instance
func newLogger() logger {
	f, err := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("[Error] %v", err)
	}
	log.SetOutput(f)

	return logger{file: f}
}

// Close closes the file of the log instance
func (l *logger) Close() {
	err := l.file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// Add writes the msg to log file
func (l logger) Add(logmsg string) {
	log.Println(fmt.Sprintf("[SERVER %s] %s", time.Now().Format(time.RFC850), logmsg))
}
