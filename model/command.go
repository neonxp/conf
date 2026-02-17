package model

type Command struct {
	Name      string
	Arguments Values
	Body      *Doc
}

type Commands []*Command
