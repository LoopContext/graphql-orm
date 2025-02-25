package model

import (
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/jinzhu/inflection"
	"github.com/loopcontext/anycase"

	"github.com/graphql-go/graphql/language/ast"
)

func queryDefinition(m *Model) *ast.ObjectDefinition {
	fields := []*ast.FieldDefinition{
		createFederationServiceQueryField(),
	}

	for _, obj := range m.ObjectEntities() {
		fields = append(fields, fetchFieldDefinition(obj), listFieldDefinition(obj))
	}
	return &ast.ObjectDefinition{
		Kind: kinds.ObjectDefinition,
		Name: &ast.Name{
			Kind:  kinds.Name,
			Value: "Query",
		},
		Fields: fields,
	}
}

func fetchFieldDefinition(obj Object) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind: kinds.FieldDefinition,
		Name: nameNode(inflection.Singular(anycase.ToLowerCamel(obj.Name()))),
		Type: namedType(obj.Name()),
		Arguments: []*ast.InputValueDefinition{
			&ast.InputValueDefinition{
				Kind:        kinds.InputValueDefinition,
				Name:        nameNode("id"),
				Description: &ast.StringValue{Kind: kinds.StringValue, Value: "Input for searching by object ID"},
				Type:        namedType("ID"),
			},
			&ast.InputValueDefinition{
				Kind:        kinds.InputValueDefinition,
				Name:        nameNode("q"),
				Description: &ast.StringValue{Kind: kinds.StringValue, Value: "Input for textual searching across selected field (string only)"},
				Type:        namedType("String"),
			},
			&ast.InputValueDefinition{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("filter"),
				Type: namedType(obj.Name() + "FilterType"),
			},
		},
	}
}

func listFieldDefinition(obj Object) *ast.FieldDefinition {
	return listFieldResultTypeDefinition(obj, inflection.Plural(anycase.ToLowerCamel(obj.Name())))
}
func listFieldResultTypeDefinition(obj Object, name string) *ast.FieldDefinition {
	return &ast.FieldDefinition{
		Kind: kinds.FieldDefinition,
		Name: nameNode(name),
		Type: nonNull(namedType(obj.Name() + "ResultType")),
		Arguments: []*ast.InputValueDefinition{
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("offset"),
				Type: namedType("Int"),
			},
			{
				Kind:         kinds.InputValueDefinition,
				Name:         nameNode("limit"),
				DefaultValue: &ast.IntValue{Kind: kinds.IntValue, Value: "30"},
				Type:         namedType("Int"),
			},
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("q"),
				Type: namedType("String"),
			},
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("sort"),
				Type: listType(nonNull(namedType(obj.Name() + "SortType"))),
			},
			{
				Kind: kinds.InputValueDefinition,
				Name: nameNode("filter"),
				Type: namedType(obj.Name() + "FilterType"),
			},
		},
	}
}
