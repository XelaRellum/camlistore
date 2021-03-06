/*
Copyright 2012 The Camlistore Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package readerutil provides and operates on io.Readers.
package readerutil // import "camlistore.org/pkg/readerutil"

import (
	"bytes"
	"io"
	"os"
)

// ReaderSize tries to determine the length of r.
func ReaderSize(r io.Reader) (size int64, ok bool) {
	switch rt := r.(type) {
	case io.Seeker:
		pos, err := rt.Seek(0, os.SEEK_CUR)
		if err != nil {
			return
		}
		end, err := rt.Seek(0, os.SEEK_END)
		if err != nil {
			return
		}
		size = end - pos
		pos1, err := rt.Seek(pos, os.SEEK_SET)
		if err != nil || pos1 != pos {
			msg := "failed to restore seek position"
			if err != nil {
				msg += ": " + err.Error()
			}
			panic(msg)
		}
		return size, true
	case *bytes.Buffer:
		return int64(rt.Len()), true
	}
	return
}
