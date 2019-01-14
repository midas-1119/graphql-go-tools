package document

// SelectionSet as specified in:
// http://facebook.github.io/graphql/draft/#SelectionSet
type SelectionSet struct {
	Fields          []int
	FragmentSpreads []int
	InlineFragments []int
}

func (s SelectionSet) NodeValueType() ValueType {
	panic("implement me")
}

func (s SelectionSet) NodeValueReference() int {
	panic("implement me")
}

func (s SelectionSet) NodeUnionMemberTypes() []ByteSlice {
	panic("implement me")
}

func (s SelectionSet) NodeSchemaDefinition() SchemaDefinition {
	panic("implement me")
}

func (s SelectionSet) NodeScalarTypeDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeObjectTypeDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeInterfaceTypeDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeUnionTypeDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeEnumTypeDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeInputObjectTypeDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeDirectiveDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeName() string {
	panic("implement me")
}

func (s SelectionSet) NodeAlias() string {
	panic("implement me")
}

func (s SelectionSet) NodeDescription() string {
	panic("implement me")
}

func (s SelectionSet) NodeArguments() []int {
	panic("implement me")
}

func (s SelectionSet) NodeArgumentsDefinition() []int {
	panic("implement me")
}

func (s SelectionSet) NodeDirectives() []int {
	panic("implement me")
}

func (s SelectionSet) NodeEnumValuesDefinition() []int {
	panic("implement me")
}

func (s SelectionSet) NodeFields() []int {
	return s.Fields
}

func (s SelectionSet) NodeFieldsDefinition() []int {
	panic("implement me")
}

func (s SelectionSet) NodeFragmentSpreads() []int {
	return s.FragmentSpreads
}

func (s SelectionSet) NodeInlineFragments() []int {
	return s.InlineFragments
}

func (s SelectionSet) NodeVariableDefinitions() []int {
	panic("implement me")
}

func (s SelectionSet) NodeType() int {
	panic("implement me")
}

func (s SelectionSet) NodeOperationType() OperationType {
	panic("implement me")
}

func (s SelectionSet) NodeValue() int {
	panic("implement me")
}

func (s SelectionSet) NodeDefaultValue() int {
	panic("implement me")
}

func (s SelectionSet) NodeImplementsInterfaces() []ByteSlice {
	panic("implement me")
}

// IsEmpty returns true if fields, fragment spreads and inline fragments are 0
func (s SelectionSet) IsEmpty() bool {
	return len(s.Fields) == 0 &&
		len(s.FragmentSpreads) == 0 &&
		len(s.InlineFragments) == 0
}
