// Copyright 2019 eBay Inc.
// Primary authors: Simon Fell, Diego Ongaro,
//                  Raymond Kroeker, and Sathish Kandasamy.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package unicode contains Unicode text functionality for Akutan store.
package unicode

import "golang.org/x/text/unicode/norm"

// Normalize converts strings to normalized Unicode string. Both write
// and read operations to the Akutan store should use this function to normalize
// all strings before writing and reading respectively, so that canonical
// equivalent strings (ex., 'Beyonce\u0301' and 'Beyonc\u00e9') are treated
// equal.
func Normalize(s string) string {
	return norm.NFC.String(s)
}
