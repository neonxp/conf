package conf

import (
	"fmt"
	"os"

	"go.neonxp.ru/conf/model"
	"go.neonxp.ru/conf/parser"
)

func LoadFile(filename string) (model.Group, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed load file: %w", err)
	}

	return Load(filename, content)
}

func Load(name string, input []byte) (model.Group, error) {

	res, err := parser.Parse(name, input)
	if err != nil {
		return nil, parser.CaretErrors(err, string(input))
	}

	return res.(model.Group), nil
}
