package testsgo

import (
	"testing"
)

func TestExecutableDefinitionsRule(t *testing.T) {

	expectErrors := func(queryStr string) ResultCompare {
		return ExpectValidationErrors("ExecutableDefinitionsRule", queryStr)
	}

	expectValid := func(queryStr string) {
		expectErrors(queryStr)([]Err{})
	}

	t.Run("Validate: Executable definitions", func(t *testing.T) {
		t.Run("with only operation", func(t *testing.T) {
			expectValid(`
      query Foo {
        dog {
          name
        }
      }
    `)
		})

		t.Run("with operation and fragment", func(t *testing.T) {
			expectValid(`
      query Foo {
        dog {
          name
          ...Frag
        }
      }

      fragment Frag on Dog {
        name
      }
    `)
		})

		t.Run("with type definition", func(t *testing.T) {
			expectErrors(`
      query Foo {
        dog {
          name
        }
      }

      type Cow {
        name: String
      }

      extend type Dog {
        color: String
      }
    `)([]Err{
				{
					message:   `The "Cow" definition is not executable.`,
					locations: []Loc{{line: 8, column: 7}},
				},
				{
					message:   `The "Dog" definition is not executable.`,
					locations: []Loc{{line: 12, column: 7}},
				},
			})
		})

		t.Run("with schema definition", func(t *testing.T) {
			expectErrors(`
      schema {
        query: Query
      }

      type Query {
        test: String
      }

      extend schema @directive
    `)([]Err{
				{
					message:   "The schema definition is not executable.",
					locations: []Loc{{line: 2, column: 7}},
				},
				{
					message:   `The "Query" definition is not executable.`,
					locations: []Loc{{line: 6, column: 7}},
				},
				{
					message:   "The schema definition is not executable.",
					locations: []Loc{{line: 10, column: 7}},
				},
			})
		})
	})

}
