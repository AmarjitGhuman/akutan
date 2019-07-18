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

package plandef

import (
	"testing"

	"github.com/ebay/akutan/query/parser"
	"github.com/ebay/akutan/util/cmp"
	"github.com/stretchr/testify/assert"
)

func Test_HashJoin(t *testing.T) {
	vars := VarSet{&Variable{Name: "alice"}}
	join := &HashJoin{Variables: vars}
	assert.Equal(t, "HashJoin (inner) ?alice", join.String())
	assert.Equal(t, "HashJoin (inner) ?alice", cmp.GetKey(join))
	join = &HashJoin{Variables: vars, Specificity: parser.MatchOptional}
	assert.Equal(t, "HashJoin (left) ?alice", join.String())
	assert.Equal(t, "HashJoin (left) ?alice", cmp.GetKey(join))
}

func Test_LoopJoin(t *testing.T) {
	vars := VarSet{&Variable{Name: "alice"}}
	join := &LoopJoin{Variables: vars}
	assert.Equal(t, "LoopJoin (inner) ?alice", join.String())
	assert.Equal(t, "LoopJoin (inner) ?alice", cmp.GetKey(join))
	join = &LoopJoin{Variables: vars, Specificity: parser.MatchOptional}
	assert.Equal(t, "LoopJoin (left) ?alice", join.String())
	assert.Equal(t, "LoopJoin (left) ?alice", cmp.GetKey(join))
}
