// Package conf library for parsing `.conf` configuration files.
package conf

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

import (
	"fmt"
	"os"

	"go.neonxp.ru/conf/model"
	"go.neonxp.ru/conf/parser"
)

func LoadFile(filename string) (model.Group, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed load file: %w", err)
	}

	return Load(filename, content)
}

func Load(name string, input []byte) (model.Group, error) {
	res, err := parser.Parse(name, input)
	if err != nil {
		return nil, parser.CaretErrors(err, string(input))
	}

	return res.(model.Group), nil
}
