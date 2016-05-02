// Copyright (C) 2015 Space Monkey, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitor

import (
	"fmt"
)

type Func struct {
	// sync/atomic things
	FuncStats

	// constructor things
	id    int64
	scope *Scope
	name  string
}

func newFunc(s *Scope, name string) (f *Func) {
	f = &Func{
		id:    NewId(),
		scope: s,
		name:  name,
	}
	initFuncStats(&f.FuncStats)
	return f
}

func (f *Func) ShortName() string { return f.name }

func (f *Func) FullName() string {
	return fmt.Sprintf("%s.%s", f.scope.name, f.name)
}

func (f *Func) Id() int64     { return f.id }
func (f *Func) Scope() *Scope { return f.scope }

func (f *Func) Parents(cb func(f *Func)) {
	f.FuncStats.parents(cb)
}
