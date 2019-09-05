package parser

import (
	"github.com/jensneuse/graphql-go-tools/legacy/document"
	"github.com/jensneuse/graphql-go-tools/pkg/lexer/keyword"
)

func (p *Parser) parseVariableDefinitions(index *[]int) (err error) {

	if open := p.peekExpect(keyword.LPAREN, true); !open {
		return
	}

	for {
		next := p.l.Peek(true)

		switch next {
		case keyword.VARIABLE:

			variable := p.l.Read()

			variableDefinition := document.VariableDefinition{
				DefaultValue: -1,
				Variable:     variable.Literal,
			}

			variableDefinition.Position.MergeStartIntoStart(variable.TextPosition)

			_, err = p.readExpect(keyword.COLON, "parseVariableDefinitions")
			if err != nil {
				return err
			}

			err = p.parseType(&variableDefinition.Type)
			if err != nil {
				return err
			}

			variableDefinition.DefaultValue, err = p.parseDefaultValue()
			if err != nil {
				return err
			}

			variableDefinition.Position.MergeStartIntoEnd(p.TextPosition())
			*index = append(*index, p.putVariableDefinition(variableDefinition))

		case keyword.RPAREN:
			p.l.Read()
			return err
		default:
			invalid := p.l.Read()
			return newErrInvalidType(invalid.TextPosition, "parseVariableDefinitions", "variable/bracket close", invalid.Keyword.String())
		}
	}
}
