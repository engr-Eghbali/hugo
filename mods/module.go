// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mods

import (
	"github.com/gohugoio/hugo/config"
)

type Module interface {
	// Returns the path to this module.
	// This will either be the module path, e.g. "github.com/gohugoio/myshortcodes",
	// or the path below your /theme folder, e.g. "mytheme".
	Path() string

	// Directory holding files for this module.
	Dir() string

	// Returns whether Dir points below the _vendor dir.
	Vendor() bool

	// Returns whether this is a Go Module.
	IsGoMod() bool

	// The module version, "none" if not applicable.
	Version() string

	// In the dependency tree, this is the first module that defines this module
	// as a dependency.
	Owner() Module

	// Optional configuration filename (e.g. "/themes/mytheme/config.json").
	// This will be added to the special configuration watch list when in
	// server mode.
	ConfigFilename() string

	// Optional config read from the configFilename above.
	Cfg() config.Provider
}

var _ Module = (*moduleAdapter)(nil)

type moduleAdapter struct {
	// Set if not a Go module.
	path string
	dir  string

	// Set if a Go module.
	gomod *GoModule

	// May be set for all.
	vendor         bool
	owner          Module
	configFilename string
	cfg            config.Provider
}

func (m *moduleAdapter) Path() string {
	if m.gomod != nil {
		return m.gomod.Path
	}
	return m.path
}

func (m *moduleAdapter) Dir() string {
	// This may point to the _vendor dir.
	return m.dir
}

func (m *moduleAdapter) Vendor() bool {
	return m.vendor
}

func (m *moduleAdapter) IsGoMod() bool {
	return m.gomod != nil
}

func (m *moduleAdapter) Version() string {
	if m.gomod != nil {
		return m.gomod.Version
	}
	return "none"
}

func (m *moduleAdapter) Owner() Module {
	return m.owner
}

func (m *moduleAdapter) ConfigFilename() string {
	return m.configFilename
}

func (m *moduleAdapter) Cfg() config.Provider {
	return m.cfg
}

type Modules []Module

type Config struct {
	// Decode: support :default =>
	// ^assets$|
	IncludeDirs string
}
