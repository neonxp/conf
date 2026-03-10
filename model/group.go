// Package model implements custom types and methods used in conf parser.
package model

// This file is part of conf library.
// Copyright (C) 2026  Alexander NeonXP Kiryukhin <i@neonxp.ru>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

import "iter"

type Group []Directive

// Get returns first directive with given name.
func (g Group) Get(name Ident) *Directive {
	for _, c := range g {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

// Filter directives by predicate and returns iterator over filtered items.
func (g Group) Filter(predicate func(c *Directive) bool) iter.Seq[*Directive] {
	return func(yield func(*Directive) bool) {
		for _, c := range g {
			if predicate(&c) && !yield(&c) {
				return
			}
		}
	}
}

// Directives returns iterator over Directives by ident.
func (g Group) Directives(ident Ident) iter.Seq[*Directive] {
	return g.Filter(func(c *Directive) bool { return c.Name == ident })
}
