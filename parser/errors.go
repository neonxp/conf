package parser

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
			err1, shouldReturn := caretError(e, input, buffer, err)
			if shouldReturn {
				return err1
			}
		}
		return errors.New(buffer.String())
	}
	return err
}

func caretError(e error, input string, buffer bytes.Buffer, err error) (error, bool) {
	if parserErr, ok := e.(ParserError); ok {
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
		fmt.Fprintf(&buffer, "%s\n%s\n%s\n", line, strings.Repeat(" ", pos)+"^", err.Error())

		return err, true
	} else {
		return err, true
	}
	return nil, false
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
