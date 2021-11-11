package model

import (
	"fmt"
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/loopcontext/anycase"
)

// ObjectRelationship struct
type ObjectRelationship struct {
	Def *ast.FieldDefinition
	Obj *Object
}

// Name ...
func (o *ObjectRelationship) Name() string {
	return o.Def.Name.Value
}

// MethodName ...
func (o *ObjectRelationship) MethodName() string {
	return anycase.ToCamel(o.Def.Name.Value)
}

// ValueForRelationshipDirectiveAttribute ...
func (o *ObjectRelationship) ValueForRelationshipDirectiveAttribute(name string) (val interface{}, ok bool) {
	for _, d := range o.Def.Directives {
		if d.Name.Value == "relationship" {
			for _, arg := range d.Arguments {
				if arg.Name.Value == name {
					val = arg.Value.GetValue()
					ok = true
					return
				}
			}
		}
	}
	return
}

// StringForRelationshipDirectiveAttribute ...
func (o *ObjectRelationship) StringForRelationshipDirectiveAttribute(name string) (val string, ok bool) {
	value, ok := o.ValueForRelationshipDirectiveAttribute(name)
	if !ok {
		return
	}
	val, ok = value.(string)
	if !ok {
		panic(fmt.Sprintf("invalid %s value for %s->%s relationship", name, o.Obj.Name(), o.Name()))
	}
	return
}

// BoolForRelationshipDirectiveAttribute ...
func (o *ObjectRelationship) BoolForRelationshipDirectiveAttribute(name string) (val bool, ok bool) {
	value, ok := o.ValueForRelationshipDirectiveAttribute(name)
	if !ok {
		return
	}
	val, ok = value.(bool)
	if !ok {
		panic(fmt.Sprintf("invalid %s value for %s->%s relationship", name, o.Obj.Name(), o.Name()))
	}
	return
}

// InverseRelationshipName ...
func (o *ObjectRelationship) InverseRelationshipName() string {
	val, ok := o.StringForRelationshipDirectiveAttribute("inverse")
	if !ok {
		panic(fmt.Sprintf("missing inverse value for %s->%s relationship", o.Obj.Name(), o.Name()))
	}
	return val
}

// Preload ...
func (o *ObjectRelationship) Preload() bool {
	val, _ := o.BoolForRelationshipDirectiveAttribute("preload")
	return val
}

// Target ...
func (o *ObjectRelationship) Target() *Object {
	target := o.Obj.Model.Object(o.TargetType())
	return &target
}

// InverseRelationship ...
func (o *ObjectRelationship) InverseRelationship() *ObjectRelationship {
	return o.Target().Relationship(o.InverseRelationshipName())
}

// IsToMany ...
func (o *ObjectRelationship) IsToMany() bool {
	t := getNullableType(o.Def.Type)
	return isListType(t)
}

// IsToOne ...
func (o *ObjectRelationship) IsToOne() bool {
	return !o.IsToMany()
}

// IsManyToMany ...
func (o *ObjectRelationship) IsManyToMany() bool {
	return o.IsToMany() && o.InverseRelationship().IsToMany()
}

// IsManyToOne ...
func (o *ObjectRelationship) IsManyToOne() bool {
	return o.IsToMany() && !o.InverseRelationship().IsToMany()
}

// IsOneToMany ...
func (o *ObjectRelationship) IsOneToMany() bool {
	return !o.IsToMany() && o.InverseRelationship().IsToMany()
}

// IsSelfReferencing ...
func (o *ObjectRelationship) IsSelfReferencing() bool {
	inv := o.InverseRelationship()
	return o.Obj.Name() == inv.Obj.Name() && o.Name() == inv.Name()
}

// IsMainRelationshipForManyToMany ...
func (o *ObjectRelationship) IsMainRelationshipForManyToMany() bool {
	main := o.MainRelationshipForManyToMany()
	return o.Obj.Name() == main.Obj.Name() && o.Name() == main.Name()
}

// IsNonNull ...
func (o *ObjectRelationship) IsNonNull() bool {
	return isNonNullType(o.Def.Type)
}

// ReturnType ...
func (o *ObjectRelationship) ReturnType() string {
	nt := getNamedType(o.Def.Type).(*ast.Named)
	if o.IsToMany() {
		return fmt.Sprintf("[]*%s", nt.Name.Value)
	}
	return fmt.Sprintf("*%s", nt.Name.Value)
}

// TargetType ...
func (o *ObjectRelationship) TargetType() string {
	nt := getNamedType(o.Def.Type).(*ast.Named)
	return nt.Name.Value
}

// GoType ...
func (o *ObjectRelationship) GoType() string {
	return o.ReturnType()
}

// ChangesName ...
func (o *ObjectRelationship) ChangesName() string {
	name := o.MethodName()
	if o.IsToMany() {
		name += "IDs"
	} else {
		name += "ID"
	}
	return name
}

// ChangesType ...
func (o *ObjectRelationship) ChangesType() string {
	if o.IsToMany() {
		return "[]*string"
	}
	return "*string"
}

// ModelTags ...
func (o *ObjectRelationship) ModelTags() string {
	tags := fmt.Sprintf(`json:"%s"`, o.Name())
	invrel := o.InverseRelationship()
	if o.IsManyToOne() {
		tags += fmt.Sprintf(" gorm:\"foreignkey:%sID\"", invrel.MethodName())
	} else if o.IsManyToMany() {
		rel := o.MainRelationshipForManyToMany()
		if o.IsSelfReferencing() {
			tags += fmt.Sprintf(" gorm:\"many2many:%s;jointable_foreignkey:%sId;association_jointable_foreignkey:%sId\"", rel.ManyToManyJoinTable(), inflection.Singular(strings.ToLower(o.Obj.Name())), inflection.Singular(o.InverseRelationshipName()))
		} else {
			tags += fmt.Sprintf(" gorm:\"many2many:%s;jointable_foreignkey:%sId;association_jointable_foreignkey:%sId\"", rel.ManyToManyJoinTable(), inflection.Singular(o.InverseRelationshipName()), inflection.Singular(o.Name()))
		}
	}
	return tags
}

// ManyToManyJoinTable ...
func (o *ObjectRelationship) ManyToManyJoinTable() string {
	m := o.MainRelationshipForManyToMany()
	return m.Obj.LowerName() + "_" + m.Name()
}

// ManyToManyObjectName ...
func (o *ObjectRelationship) ManyToManyObjectName() string {
	m := o.MainRelationshipForManyToMany()
	return m.Obj.Name() + "_" + m.Name()
}

// ManyToManyObjectNameCC ...
func (o *ObjectRelationship) ManyToManyObjectNameCC() string {
	m := o.MainRelationshipForManyToMany()
	return anycase.ToCamel(m.Obj.Name() + "_" + m.Name())
}

// MainRelationshipForManyToMany ...
func (o *ObjectRelationship) MainRelationshipForManyToMany() *ObjectRelationship {
	inversed := o.InverseRelationship()
	if inversed.Name() > o.Name() {
		return inversed
	}
	return o
}

// JoinString ...
func (o *ObjectRelationship) JoinString() string {
	join := ""
	if o.IsManyToMany() {
		joinTable := o.ManyToManyJoinTable()
		join += fmt.Sprintf("\"LEFT JOIN \"+dialect.Quote(TableName(\"%[1]s\"))+\" \"+dialect.Quote(_alias+\"_jointable\")+\" ON \"+dialect.Quote(alias)+\".id = \"+dialect.Quote(_alias+\"_jointable\")+\".\"+dialect.Quote(\"%[3]sId\")+\" LEFT JOIN \"+dialect.Quote(TableName(\"%[2]s\"))+\" \"+dialect.Quote(_alias)+\" ON \"+dialect.Quote(_alias+\"_jointable\")+\".\"+dialect.Quote(\"%[4]sId\")+\" = \"+dialect.Quote(_alias)+\".id\"", joinTable, o.Target().TableName(), inflection.Singular(o.InverseRelationshipName()), inflection.Singular(o.Name()))
	} else if o.IsToOne() {
		join += fmt.Sprintf("\"LEFT JOIN \"+dialect.Quote(TableName(\"%[1]s\"))+\" \"+dialect.Quote(_alias)+\" ON \"+dialect.Quote(_alias)+\".id = \"+dialect.Quote(alias)+\".\"+dialect.Quote(\"%[2]sId\")", o.Target().TableName(), o.Name())
	} else if o.IsToMany() {
		join += fmt.Sprintf("\"LEFT JOIN \"+dialect.Quote(TableName(\"%[1]s\"))+\" \"+dialect.Quote(_alias)+\" ON \"+dialect.Quote(_alias)+\".\"+dialect.Quote(\"%[3]sId\")+\" = \"+dialect.Quote(alias)+\".id\"", o.Target().TableName(), o.Name(), o.InverseRelationshipName())
	}
	return join
}

// ForeignKeyDestinationColumn ...
func (o *ObjectRelationship) ForeignKeyDestinationColumn() string {
	if o.IsToOne() {
		return "id"
	}
	if o.IsManyToMany() {
		return inflection.Singular(o.InverseRelationshipName()) + "Id"
	}
	if o.IsToMany() {
		return inflection.Singular(o.InverseRelationshipName()) + "Id"
	}
	return ""
}

// ForeignKeyDestinationColumnCC ...
func (o *ObjectRelationship) ForeignKeyDestinationColumnCC() string {
	if o.IsToOne() {
		return "ID"
	}
	if o.IsManyToMany() {
		return anycase.ToCamel(inflection.Singular(o.InverseRelationshipName())) + "ID"
	}
	return ""
}

// OnDelete ...
func (o *ObjectRelationship) OnDelete(def string) string {
	str, exists := o.StringForRelationshipDirectiveAttribute("onDelete")
	if !exists {
		return def
	}
	return str
}

// OnUpdate ...
func (o *ObjectRelationship) OnUpdate(def string) string {
	str, exists := o.StringForRelationshipDirectiveAttribute("onUpdate")
	if !exists {
		return def
	}
	return str
}
