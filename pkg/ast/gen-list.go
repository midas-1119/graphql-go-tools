// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package ast

import "github.com/jensneuse/graphql-go-tools/pkg/lexing/position"

type DirectiveList struct {
	Open          position.Position
	Close         position.Position
	current       Directive
	currentRef    int
	nextRef       int
	isInitialized bool
}

type DirectiveGetter interface {
	GetDirective(ref int) (node Directive, nextRef int)
}

func NewDirectiveList(first int) DirectiveList {
	nodeList := DirectiveList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *DirectiveList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *DirectiveList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *DirectiveList) Next(getter DirectiveGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetDirective(n.nextRef)
	return true
}

func (n *DirectiveList) Value() (Directive, int) {
	return n.current, n.currentRef
}

type FieldDefinitionList struct {
	Open          position.Position
	Close         position.Position
	current       FieldDefinition
	currentRef    int
	nextRef       int
	isInitialized bool
}

type FieldDefinitionGetter interface {
	GetFieldDefinition(ref int) (node FieldDefinition, nextRef int)
}

func NewFieldDefinitionList(first int) FieldDefinitionList {
	nodeList := FieldDefinitionList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *FieldDefinitionList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *FieldDefinitionList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *FieldDefinitionList) Next(getter FieldDefinitionGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetFieldDefinition(n.nextRef)
	return true
}

func (n *FieldDefinitionList) Value() (FieldDefinition, int) {
	return n.current, n.currentRef
}

type RootOperationTypeDefinitionList struct {
	Open          position.Position
	Close         position.Position
	current       RootOperationTypeDefinition
	currentRef    int
	nextRef       int
	isInitialized bool
}

type RootOperationTypeDefinitionGetter interface {
	GetRootOperationTypeDefinition(ref int) (node RootOperationTypeDefinition, nextRef int)
}

func NewRootOperationTypeDefinitionList(first int) RootOperationTypeDefinitionList {
	nodeList := RootOperationTypeDefinitionList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *RootOperationTypeDefinitionList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *RootOperationTypeDefinitionList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *RootOperationTypeDefinitionList) Next(getter RootOperationTypeDefinitionGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetRootOperationTypeDefinition(n.nextRef)
	return true
}

func (n *RootOperationTypeDefinitionList) Value() (RootOperationTypeDefinition, int) {
	return n.current, n.currentRef
}

type ArgumentList struct {
	Open          position.Position
	Close         position.Position
	current       Argument
	currentRef    int
	nextRef       int
	isInitialized bool
}

type ArgumentGetter interface {
	GetArgument(ref int) (node Argument, nextRef int)
}

func NewArgumentList(first int) ArgumentList {
	nodeList := ArgumentList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *ArgumentList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *ArgumentList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *ArgumentList) Next(getter ArgumentGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetArgument(n.nextRef)
	return true
}

func (n *ArgumentList) Value() (Argument, int) {
	return n.current, n.currentRef
}

type TypeList struct {
	Open          position.Position
	Close         position.Position
	current       Type
	currentRef    int
	nextRef       int
	isInitialized bool
}

type TypeGetter interface {
	GetType(ref int) (node Type, nextRef int)
}

func NewTypeList(first int) TypeList {
	nodeList := TypeList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *TypeList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *TypeList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *TypeList) Next(getter TypeGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetType(n.nextRef)
	return true
}

func (n *TypeList) Value() (Type, int) {
	return n.current, n.currentRef
}

type InputValueDefinitionList struct {
	Open          position.Position
	Close         position.Position
	current       InputValueDefinition
	currentRef    int
	nextRef       int
	isInitialized bool
}

type InputValueDefinitionGetter interface {
	GetInputValueDefinition(ref int) (node InputValueDefinition, nextRef int)
}

func NewInputValueDefinitionList(first int) InputValueDefinitionList {
	nodeList := InputValueDefinitionList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *InputValueDefinitionList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *InputValueDefinitionList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *InputValueDefinitionList) Next(getter InputValueDefinitionGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetInputValueDefinition(n.nextRef)
	return true
}

func (n *InputValueDefinitionList) Value() (InputValueDefinition, int) {
	return n.current, n.currentRef
}

type EnumValueDefinitionList struct {
	Open          position.Position
	Close         position.Position
	current       EnumValueDefinition
	currentRef    int
	nextRef       int
	isInitialized bool
}

type EnumValueDefinitionGetter interface {
	GetEnumValueDefinition(ref int) (node EnumValueDefinition, nextRef int)
}

func NewEnumValueDefinitionList(first int) EnumValueDefinitionList {
	nodeList := EnumValueDefinitionList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *EnumValueDefinitionList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *EnumValueDefinitionList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *EnumValueDefinitionList) Next(getter EnumValueDefinitionGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetEnumValueDefinition(n.nextRef)
	return true
}

func (n *EnumValueDefinitionList) Value() (EnumValueDefinition, int) {
	return n.current, n.currentRef
}

type ValueList struct {
	Open          position.Position
	Close         position.Position
	current       Value
	currentRef    int
	nextRef       int
	isInitialized bool
}

type ValueGetter interface {
	GetValue(ref int) (node Value, nextRef int)
}

func NewValueList(first int) ValueList {
	nodeList := ValueList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *ValueList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *ValueList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *ValueList) Next(getter ValueGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetValue(n.nextRef)
	return true
}

func (n *ValueList) Value() (Value, int) {
	return n.current, n.currentRef
}

type ObjectFieldList struct {
	Open          position.Position
	Close         position.Position
	current       ObjectField
	currentRef    int
	nextRef       int
	isInitialized bool
}

type ObjectFieldGetter interface {
	GetObjectField(ref int) (node ObjectField, nextRef int)
}

func NewObjectFieldList(first int) ObjectFieldList {
	nodeList := ObjectFieldList{}
	nodeList.SetFirst(first)
	return nodeList
}

func (n *ObjectFieldList) SetFirst(first int) {
	n.nextRef = first
	n.isInitialized = first != -1
}

func (n *ObjectFieldList) HasNext() bool {
	return n.isInitialized && n.nextRef != -1
}

func (n *ObjectFieldList) Next(getter ObjectFieldGetter) bool {
	if !n.isInitialized || n.nextRef == -1 {
		return false
	}
	n.currentRef = n.nextRef
	n.current, n.nextRef = getter.GetObjectField(n.nextRef)
	return true
}

func (n *ObjectFieldList) Value() (ObjectField, int) {
	return n.current, n.currentRef
}
