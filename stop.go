// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package gse

import (
	"path"
	"strings"

	"github.com/go-vgo/gt/file"
)

// StopWordMap default contains some stop words.
var StopWordMap = map[string]bool{
	"啊": true,
}

// LoadStop load stop word files add token to map
func (seg *Segmenter) LoadStop(files ...string) error {
	if len(files) <= 0 {
		var (
			dictDir  = path.Join(path.Dir(getCurrentFilePath()), "data")
			dictPath = path.Join(dictDir, "dict/stop_word.txt")
		)

		files = append(files, dictPath)
	}

	name := strings.Split(files[0], ", ")
	for i := 0; i < len(name); i++ {
		s, err := file.Read(name[i])
		if err != nil {
			return err
		}

		ns := strings.Split(s, "\n")
		for h := 0; h < len(ns); h++ {
			StopWordMap[ns[h]] = true
		}
	}

	return nil
}

// AddStop adds a token into StopWord dictionary.
func (seg *Segmenter) AddStop(text string) {
	StopWordMap[text] = true
}

// IsStop checks if a given word is stop word.
func (seg *Segmenter) IsStop(s string) bool {
	// _, ok := StopWordMap[s]
	// return ok
	return StopWordMap[s]
}
