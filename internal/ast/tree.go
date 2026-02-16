package ast

import (
	"go.neonxp.ru/conf/internal/parser"
	"modernc.org/scanner"
)

func Parse(p *parser.Parser, data []int32) []*Node {
	nodes := make([]*Node, 0, 2)
	for len(data) != 0 {
		next := int32(1)
		var node *Node
		switch n := data[0]; {
		case n < 0:
			next = 2 + data[1]
			node = &Node{
				Symbol:   parser.Symbol(-data[0]),
				Children: Parse(p, data[2:next]),
			}
		default:
			tok := p.Token(n)
			node = &Node{
				Token:  tok,
				Symbol: parser.Symbol(tok.Ch),
				Source: tok.Src(),
				Col:    tok.Position().Column,
				Line:   tok.Position().Line,
			}
		}
		nodes = append(nodes, node)
		data = data[next:]
	}
	return nodes
}

type Node struct {
	Symbol   parser.Symbol
	Token    scanner.Token
	Children []*Node
	Source   string
	Col      int
	Line     int
}
