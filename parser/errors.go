// Package parser parses conf language.
package parser

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
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type ErrorLister interface {
	Errors() []error
}

func (e errList) Errors() []error {
	return e
}

// ParserError is the public interface to errors of type parserError
type ParserError interface {
	Error() string
	InnerError() error
	Pos() (int, int, int)
	Expected() []string
}

func (p *parserError) InnerError() error {
	return p.Inner
}

func (p *parserError) Pos() (line, col, offset int) {
	return p.pos.line, p.pos.col, p.pos.offset
}

func (p *parserError) Expected() []string {
	return p.expected
}

func CaretErrors(err error, input string) error {
	if el, ok := err.(ErrorLister); ok {
		var buffer bytes.Buffer
		for _, e := range el.Errors() {
			if err := caretError(e, input); err != nil {
				buffer.WriteString(err.Error())
			}
		}
		return errors.New(buffer.String())
	}
	return err
}

func caretError(err error, input string) error {
	if parserErr, ok := err.(ParserError); ok {
		_, col, off := parserErr.Pos()
		line := extractLine(input, off)
		if col >= len(line) {
			col = len(line) - 1
		} else {
			if col > 0 {
				col--
			}
		}
		if col < 0 {
			col = 0
		}
		pos := col
		for _, chr := range line[:col] {
			if chr == '\t' {
				pos += 7
			}
		}
		return fmt.Errorf("%s\n%s\n%w", line, strings.Repeat(" ", pos)+"^", err)
	}
	return err
}

func extractLine(input string, initPos int) string {
	if initPos < 0 {
		initPos = 0
	}
	if initPos >= len(input) && len(input) > 0 {
		initPos = len(input) - 1
	}
	startPos := initPos
	endPos := initPos
	for ; startPos > 0; startPos-- {
		if input[startPos] == '\n' {
			if startPos != initPos {
				startPos++
				break
			}
		}
	}
	for ; endPos < len(input); endPos++ {
		if input[endPos] == '\n' {
			if endPos == initPos {
				endPos++
			}
			break
		}
	}
	return input[startPos:endPos]
}
