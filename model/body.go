package model

import (
	"errors"
	"fmt"
)

type Body []any

var ErrInvalidType = errors.New("invalid type")

func (d Body) Execute(v Visitor) error {
	for _, it := range d {
		switch it := it.(type) {
		case *Setting:
			if err := v.VisitSetting(it.Key, it.Value); err != nil {
				return err
			}
		case *Directive:
			if err := v.VisitDirective(it.Name, it.Arguments, it.Body); err != nil {
				return err
			}
		default:
			return fmt.Errorf("%w: %t", ErrInvalidType, it)
		}
	}
	return nil
}
