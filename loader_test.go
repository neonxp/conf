// Package conf_test tests.
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
package conf_test

import (
	"fmt"

	"go.neonxp.ru/conf"
)

func ExampleLoadFile() {
	result, err := conf.LoadFile("./example.conf")
	if err != nil {
		panic("\n" + err.Error())
	}

	out(
		result.Get("group1").
			Group().Get("group2").
			Group().Get("group3").
			Group().Get("key").Value(),
	) // → value
	out(
		result.Get("group1").
			Group().Get("group2").
			Group().Get("group3").
			Group().Get("int_key").Value(),
	) // → 123
	out(
		result.Get("group1").
			Group().Get("group2").
			Group().Get("group3").
			Group().Get("float_key").Value(),
	) // → 123.321
	out(
		result.Get("group1").
			Group().Get("group2").
			Group().Get("group3").
			Group().Get("bool_key").Value(),
	) // → true

	out(
		result.Get("group1").
			Group().Get("group2").
			Group().Get("group3").Value(),
	) // → param3

	out(
		result.Get("group1").
			Group().Get("group2").
			Group().Get("external_vars").
			StringExt(" ", func(key string) (string, bool) {
				if key == "EXTERNAL_VAR" {
					return "external variable", true
				}
				return "unknnown var", false
			}),
	) // → Concatenate with external variable and integer 123

	// Output: value (string)
	// 123 (int)
	// 123.321 (float64)
	// true (bool)
	// param3 (model.Ident)
	// Concatenate with external variable and integer 123 (string)
}

func out(a any) {
	fmt.Printf("%v (%T)\n", a, a)
}
