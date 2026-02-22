package model

type Visitor interface {
	VisitDirective(ident string, args Values, body Body) error
	VisitSetting(key string, values Values) error
}
