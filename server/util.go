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

import "strconv"

// This will throw an error. Be careful if you use uint64 values,
// as they may start out small and work without error, but increment
// over time and start throwing errors.
func resourceID(id string) (uint64, error) {
	var (
		intID uint64
		err   error
	)
	// if we have valid id
	if len(id) > 0 {
		intID, err = strconv.ParseUint(id, 10, 64)
		if err != nil {
			return 0, err
		}
		// test if the high bit is set
		if toHighSet(intID) {
			return 0, errHighBitSet
		}
	} else {
		return 0, errParamNotSet
	}

	// everything is fine 0,nil
	return intID, nil
}

func logIT(err error) {
	if err != nil {
		Logger.Add(err.Error())
	}
}

func toHighSet(n uint64) bool {
	flag := false
	if n&(1<<63) != 0x0 {
		flag = true
	}
	return flag
}
