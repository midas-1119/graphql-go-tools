package parser

import (
	"github.com/jensneuse/graphql-go-tools/pkg/document"
	"github.com/jensneuse/graphql-go-tools/pkg/lexing/keyword"
)

func (p *Parser) parseField(index *[]int) (err error) {

	var field document.Field
	p.initField(&field)

	firstIdent := p.l.Read()
	field.Name = firstIdent.Literal

	hasAlias := p.peekExpect(keyword.COLON, true)

	if hasAlias {
		field.Alias = field.Name
		fieldName, err := p.readExpect(keyword.IDENT, "parseField")
		if err != nil {
			return err
		}

		field.Name = fieldName.Literal
	}

	err = p.parseArguments(&field.Arguments)
	if err != nil {
		return
	}

	err = p.parseDirectives(&field.Directives)
	if err != nil {
		return
	}

	err = p.parseSelectionSet(&field.SelectionSet)

	*index = append(*index, p.putField(field))

	return
}
