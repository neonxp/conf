package visitor

import (
	"errors"
	"fmt"
	"strings"

	"go.neonxp.ru/conf/model"
)

var (
	ErrEmptyQuery = errors.New("empty query")
	ErrNoChildKey = errors.New("no child key")
)

func NewDefault() *Default {
	return &Default{
		vars:     map[string]model.Values{},
		children: map[string]*Default{},
		args:     model.Values{},
	}
}

// Default просто собирает рекурсивно все переменные в дерево.
// На самом деле, для большинства сценариев конфигов его должно хватить.
type Default struct {
	vars     map[string]model.Values
	children map[string]*Default
	args     model.Values
}

func (p *Default) VisitDirective(ident string, args model.Values, body model.Body) error {
	p.children[ident] = NewDefault()
	p.children[ident].args = args
	return body.Execute(p.children[ident])
}

func (p *Default) VisitSetting(key string, values model.Values) error {
	p.vars[key] = values

	return nil
}

func (p *Default) Get(path string) (model.Values, error) {
	splitPath := strings.SplitN(path, ".", 2)
	switch len(splitPath) {
	case 1:
		if v, ok := p.vars[splitPath[0]]; ok {
			return v, nil
		}
		if child, ok := p.children[splitPath[0]]; ok {
			return child.args, nil
		}
		return nil, fmt.Errorf("%w: %s", ErrNoChildKey, splitPath[0])
	case 2:
		if child, ok := p.children[splitPath[0]]; ok {
			return child.Get(splitPath[1])
		}
		return nil, fmt.Errorf("%w: %s", ErrNoChildKey, splitPath[0])
	default:
		return nil, ErrEmptyQuery
	}
}
