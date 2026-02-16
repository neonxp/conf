package model

type Doc []any

type Assignment struct {
	Key   string
	Value []Value
}

type Command struct {
	Name      string
	Arguments []Value
	Body      Doc
}

type Value any

type Word string
