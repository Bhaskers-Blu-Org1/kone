// Copyright 2018 Google LLC All Rights Reserved.
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

package build

import (
	"path/filepath"
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Recorder composes with another Interface to record the built nodejs apps.
type Recorder struct {
	m       sync.Mutex
	Paths   []string
	Builder Interface
}

// Recorder implements Interface
var _ Interface = (*Recorder)(nil)

// IsSupportedReference implements Interface
func (r *Recorder) IsSupportedReference(base, dir string) *string {
	return r.Builder.IsSupportedReference(base, dir)
}

// Build implements Interface
func (r *Recorder) Build(base, dir string) (v1.Image, error) {
	func() {
		r.m.Lock()
		defer r.m.Unlock()
		r.Paths = append(r.Paths, filepath.Join(base, dir))
	}()
	return r.Builder.Build(base, dir)
}
