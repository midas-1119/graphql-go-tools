package parser

import (
	"bytes"
	. "github.com/franela/goblin"
	"github.com/jensneuse/graphql-go-tools/pkg/document"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"testing"
)

func TestInputValueDefinitionsParser(t *testing.T) {

	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("parser.parseInputValueDefinitions", func() {
		tests := []struct {
			it           string
			input        string
			expectErr    types.GomegaMatcher
			expectValues types.GomegaMatcher
		}{
			{
				it:        "should parse a single InputValueDefinition",
				input:     "inputValue: Int",
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Name: []byte("inputValue"),
						Type: document.NamedType{
							Name: []byte("Int"),
						},
					},
				}),
			},
			{
				it:        "should parse a InputValueDefinition with DefaultValue",
				input:     `inputValue: Int = 2`,
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Name: []byte("inputValue"),
						Type: document.NamedType{
							Name: []byte("Int"),
						},
						DefaultValue: document.IntValue{
							Val: 2,
						},
					},
				}),
			},
			{
				it:        "should parse a InputValueDefinition with Description",
				input:     `"useful description"inputValue: Int = 2`,
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Description: []byte("useful description"),
						Name:        []byte("inputValue"),
						Type: document.NamedType{
							Name: []byte("Int"),
						},
						DefaultValue: document.IntValue{
							Val: 2,
						},
					},
				}),
			},
			{
				it:        "should parse multiple InputValueDefinition",
				input:     `inputValue: Int, outputValue: String`,
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Name: []byte("inputValue"),
						Type: document.NamedType{
							Name: []byte("Int"),
						},
					},
					{
						Name: []byte("outputValue"),
						Type: document.NamedType{
							Name: []byte("String"),
						},
					},
				}),
			},
			{
				it:        "should parse a  multiple InputFieldDefinitions with Descriptions",
				input:     `"this is a inputValue"inputValue: Int, "this is a outputValue"outputValue: String`,
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Description: []byte("this is a inputValue"),
						Name:        []byte("inputValue"),
						Type: document.NamedType{
							Name: []byte("Int"),
						},
					},
					{
						Description: []byte("this is a outputValue"),
						Name:        []byte("outputValue"),
						Type: document.NamedType{
							Name: []byte("String"),
						},
					},
				}),
			},
			{
				it:        "should parse multiple InputFieldDefinitions with Descriptions and DefaultValues",
				input:     `"this is a inputValue"inputValue: Int = 2, "this is a outputValue"outputValue: String = "test"`,
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Description: []byte("this is a inputValue"),
						Name:        []byte("inputValue"),
						Type: document.NamedType{
							Name: []byte("Int"),
						},
						DefaultValue: document.IntValue{
							Val: 2,
						},
					},
					{
						Description: []byte("this is a outputValue"),
						Name:        []byte("outputValue"),
						Type: document.NamedType{
							Name: []byte("String"),
						},
						DefaultValue: document.StringValue{
							Val: []byte("test"),
						},
					},
				}),
			},
			{
				it:        "should parse nonNull Types",
				input:     "inputValue: Int!",
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Name: []byte("inputValue"),
						Type: document.NamedType{
							Name:    []byte("Int"),
							NonNull: true,
						},
					},
				}),
			},
			{
				it:        "should parse list Types",
				input:     "inputValue: [Int]",
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Name: []byte("inputValue"),
						Type: document.ListType{
							Type: document.NamedType{
								Name: []byte("Int"),
							},
						},
					},
				}),
			},
			{
				it:        "should parse an InputValueDefinition with Directives",
				input:     `inputValue: Int @fromTop(to: "bottom") @fromBottom(to: "top")`,
				expectErr: BeNil(),
				expectValues: Equal([]document.InputValueDefinition{
					{
						Name: []byte("inputValue"),
						Type: document.NamedType{
							Name: []byte("Int"),
						},
						Directives: document.Directives{
							document.Directive{
								Name: []byte("fromTop"),
								Arguments: document.Arguments{
									document.Argument{
										Name: []byte("to"),
										Value: document.StringValue{
											Val: []byte("bottom"),
										},
									},
								},
							},
							document.Directive{
								Name: []byte("fromBottom"),
								Arguments: document.Arguments{
									document.Argument{
										Name: []byte("to"),
										Value: document.StringValue{
											Val: []byte("top"),
										},
									},
								},
							},
						},
					},
				}),
			},
		}

		for _, test := range tests {
			test := test

			g.It(test.it, func() {

				reader := bytes.NewReader([]byte(test.input))
				parser := NewParser()
				parser.l.SetInput(reader)

				val, err := parser.parseInputValueDefinitions()
				Expect(err).To(test.expectErr)
				Expect(val).To(test.expectValues)
			})
		}
	})
}
