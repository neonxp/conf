package conf

import (
	"fmt"
	"os"

	"go.neonxp.ru/conf/internal/ast"
	"go.neonxp.ru/conf/internal/parser"
	"go.neonxp.ru/conf/model"
)

func New() *Conf {
	return &Conf{
		root: model.Body{},
	}
}

type Conf struct {
	root model.Body
}

func (c *Conf) LoadFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed load file: %w", err)
	}

	return c.Load(filename, content)
}

func (c *Conf) Load(name string, input []byte) error {
	p := &parser.Parser{}
	astSlice, err := p.Parse(name, input)
	if err != nil {
		return fmt.Errorf("failed parse conf content: %w", err)
	}

	astTree := ast.Parse(p, astSlice)

	doc, err := ast.ToDoc(astTree[0])
	if err != nil {
		return fmt.Errorf("failed build Doc: %w", err)
	}

	c.root = doc

	return nil
}

func (c *Conf) Process(visitor model.Visitor) error {
	return c.root.Execute(visitor)
}
