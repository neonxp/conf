package model

import (
	"iter"
)

type Ident string

type Command struct {
	Name  Ident
	Args  []any
	Group Group
}

// Value returns first argument of command.
func (c *Command) Value() any {
	if len(c.Args) > 0 {
		return c.Args[0]
	}
	return nil
}

type Group []Command

// Get returns first command with given name.
func (g Group) Get(name Ident) *Command {
	for _, c := range g {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

// Filter commands by predicate and returns iterator over filtered commands.
func (g Group) Filter(predicate func(c *Command) bool) iter.Seq[*Command] {
	return func(yield func(*Command) bool) {
		for _, c := range g {
			if predicate(&c) && !yield(&c) {
				return
			}
		}
	}
}
