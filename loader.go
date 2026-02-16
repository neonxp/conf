package conf

import (
	"fmt"
	"os"

	"go.neonxp.ru/conf/internal/ast"
	"go.neonxp.ru/conf/internal/parser"
	"go.neonxp.ru/conf/model"
)

func LoadFile(filename string) (model.Doc, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed load file: %w", err)
	}

	return Load(filename, content)
}

func Load(name string, input []byte) (model.Doc, error) {
	p := &parser.Parser{}
	astSlice, err := p.Parse(name, input)
	if err != nil {
		return nil, fmt.Errorf("failed parse conf content: %w", err)
	}

	astTree := ast.Parse(p, astSlice)

	doc, err := ast.ToDoc(astTree[0])
	if err != nil {
		return nil, fmt.Errorf("failed build Doc: %w", err)
	}

	return doc, nil
}
