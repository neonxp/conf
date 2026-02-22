package model

type Directive struct {
	Name      string
	Arguments Values
	Body      Body
}

type Directives []*Directive
