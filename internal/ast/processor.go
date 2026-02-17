package ast

import (
	"fmt"
	"strconv"

	"go.neonxp.ru/conf/internal/parser"
	"go.neonxp.ru/conf/model"
)

func ToDoc(config *Node) (*model.Doc, error) {
	if len(config.Children) < 1 {
		return nil, fmt.Errorf("invalid ast tree")
	}

	doc := config.Children[0]

	return processDoc(doc), nil
}

func processDoc(docNode *Node) *model.Doc {
	doc := model.New(len(docNode.Children))
	for _, stmt := range docNode.Children {
		processStmt(doc, stmt)
	}
	return doc
}

func processStmt(doc *model.Doc, stmt *Node) {
	ident := extractIdent(stmt.Children[0])
	nodeBody := stmt.Children[1]
	switch nodeBody.Symbol {
	case parser.Command:
		doc.AppendCommand(processCommand(ident, nodeBody))
	case parser.Assignment:
		doc.AppendAssignment(processAssignment(ident, nodeBody))
	}
}

func processCommand(ident string, command *Node) *model.Command {
	result := &model.Command{
		Name: ident,
	}

	for _, child := range command.Children {
		// Can be arguments OR body OR both
		switch child.Symbol {
		case parser.Values:
			result.Arguments = extractValues(child)
		case parser.Body:
			// Children[0] = '{', Children[1] = Body, Children[2] = '}'
			result.Body = processDoc(child.Children[1])
		}
	}

	return result
}

func processAssignment(ident string, assignment *Node) *model.Assignment {
	result := &model.Assignment{
		Key:   ident,
		Value: extractValues(assignment.Children[1]), // Children[0] = '=', Children[1] = Values, Children[2] = ';'
	}

	return result
}

func extractIdent(word *Node) string {
	return word.Children[0].Source
}

func extractValues(args *Node) []model.Value {
	result := make([]model.Value, len(args.Children))
	for i, child := range args.Children {
		v, err := extractValue(child)
		if err != nil {
			result[i] = err
			continue
		}
		result[i] = v
	}

	return result
}

func extractValue(value *Node) (any, error) {
	v := value.Children[0]
	s := v.Children[0].Source
	switch v.Symbol {
	case parser.Word:
		return model.Word(s), nil
	case parser.String:
		return unquote(s), nil
	case parser.Number:
		d, err := strconv.Atoi(s)
		if err == nil {
			return d, nil
		}
		fl, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return fl, nil
		}
		return nil, fmt.Errorf("invalid number: %s (%s)", v.Source, s)
	case parser.Boolean:
		return s == "true", nil
	default:
		return nil, fmt.Errorf("unknown value type: %s (%s)", v.Symbol, s)
	}
}

func unquote(str string) string {
	if len(str) == 0 {
		return ""
	}
	if str[0:1] == `"` || str[0:1] == "'" || str[0:1] == "`" {
		return str[1 : len(str)-1]
	}
	return str
}
