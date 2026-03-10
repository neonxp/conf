// Package model implements custom types and methods used in conf parser.
//
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
package model

import (
	"strconv"
	"strings"
)

type Directive struct {
	Name      Ident
	RawValues []any
}

// Group returns group of directive.
func (c *Directive) Group() Group {
	if c == nil {
		return Group{}
	}
	if len(c.RawValues) > 0 {
		last := c.RawValues[len(c.RawValues)-1]
		if g, ok := last.(Group); ok {
			return g
		}
	}
	return Group{}
}

// Value returns first argument of directive.
func (c *Directive) Value() any {
	args := c.Values()
	if len(args) > 0 {
		return args[0]
	}
	return nil
}

// String returns first argument of directive as string.
func (c *Directive) String() string {
	args := c.Values()
	if len(args) > 0 {
		if s, ok := args[0].(string); ok {
			return s
		}
	}
	return ""
}

// Int returns first argument of directive as int.
func (c *Directive) Int() int {
	args := c.Values()
	if len(args) > 0 {
		if s, ok := args[0].(int); ok {
			return s
		}
	}
	return 0
}

// Float returns first argument of directive as int.
func (c *Directive) Float() float64 {
	args := c.Values()
	if len(args) > 0 {
		if s, ok := args[0].(float64); ok {
			return s
		}
	}
	return 0
}

// Bool returns first argument of directive as boolean.
func (c *Directive) Bool() bool {
	args := c.Values()
	if len(args) > 0 {
		if s, ok := args[0].(bool); ok {
			return s
		}
	}
	return false
}

// Values returns directive values without group.
func (c *Directive) Values() []any {
	if c == nil {
		return []any{}
	}
	result := make([]any, 0, len(c.RawValues))
	for _, rv := range c.RawValues {
		if _, ok := rv.(Group); ok {
			continue
		}
		result = append(result, rv)
	}
	return result
}

type Lookup func(key string) (string, bool)

// StringExt joins all values of directive in single string.
// Idents passes to identLookup func, e.g.`os.LookupEnv`.
func (c *Directive) StringExt(sep string, identLookup Lookup) string {
	args := c.Values()
	stringSl := make([]string, len(args))
	for i, it := range args {
		switch it := it.(type) {
		case string:
			stringSl[i] = it
		case float64:
			stringSl[i] = strconv.FormatFloat(it, 'g', 5, 64)
		case int:
			stringSl[i] = strconv.Itoa(it)
		case bool:
			stringSl[i] = strconv.FormatBool(it)
		case Ident:
			if newVal, isset := identLookup(string(it)); isset {
				stringSl[i] = newVal
			}
		}
	}

	return strings.Join(stringSl, sep)
}
