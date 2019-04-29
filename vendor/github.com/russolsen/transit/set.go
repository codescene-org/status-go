// Copyright 2016 Russ Olsen. All Rights Reserved.
//
// This code is a Go port of the Java version created and maintained by Cognitect, therefore:
//
// Copyright 2014 Cognitect. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS-IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transit

import (
	"fmt"
)

// A set is a very simple minded representation of a set, one that
// does not enforce any redundancy constraints. Just a box
// for sets read from transit.
type Set struct {
	Contents []interface{}
}

func MakeSet(contents ...interface{}) *Set {
	return &Set{Contents: contents}
}

func NewSet(contents []interface{}) *Set {
	return &Set{Contents: contents}
}

func (s Set) String() string {
	return fmt.Sprintf("Set[%v]: %v", len(s.Contents), s.Contents)
}

func (s Set) Contains(value interface{}, mf MatchF) bool {
	for _, element := range s.Contents {
		if mf(element, value) {
			return true
		}
	}
	return false
}

func (s Set) ContainsEq(value interface{}) bool {
	return s.Contains(value, Equals)
}
