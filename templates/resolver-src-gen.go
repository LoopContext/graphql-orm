package templates

// ResolverSrcGen ...
var ResolverSrcGen = `package src

import (
	"{{.Config.Package}}/gen"
)

// NewResolver ...
func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{gen.NewGeneratedResolver(handlers, db, ec)}
}

// Resolver ...
type Resolver struct {
	*gen.GeneratedResolver
}

// MutationResolver ...
type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

// BeginTransaction ...
func (r * MutationResolver) BeginTransaction(ctx context.Context,fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err!=nil{
		tx := r.GeneratedResolver.GetDB(ctx)
		tx.Rollback()
		return err
	}
	return gen.FinishMutationContext(ctx, r.GeneratedResolver)
}

// QueryResolver ...
type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

// Mutation ...
func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{
		GeneratedMutationResolver: &gen.GeneratedMutationResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// Query ...
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{
		GeneratedQueryResolver: &gen.GeneratedQueryResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}


{{range $obj := .Model.ObjectEntities}}
	// {{$obj.Name}}ResultTypeResolver struct
	type {{$obj.Name}}ResultTypeResolver struct {
		*gen.Generated{{$obj.Name}}ResultTypeResolver
	}
	// {{$obj.Name}}ResultType ...
	func (r *Resolver) {{$obj.Name}}ResultType() gen.{{$obj.Name}}ResultTypeResolver {
		return &{{$obj.Name}}ResultTypeResolver{
			Generated{{$obj.Name}}ResultTypeResolver: &gen.Generated{{$obj.Name}}ResultTypeResolver{
				GeneratedResolver: r.GeneratedResolver,
			},
		}
	}
	{{if $obj.NeedsQueryResolver}}
		// {{$obj.Name}}Resolver struct
		type {{$obj.Name}}Resolver struct {
			*gen.Generated{{$obj.Name}}Resolver
		}
		// {{$obj.Name}} ...
		func (r *Resolver) {{$obj.Name}}() gen.{{$obj.Name}}Resolver {
			return &{{$obj.Name}}Resolver{
				Generated{{$obj.Name}}Resolver: &gen.Generated{{$obj.Name}}Resolver{
					GeneratedResolver: r.GeneratedResolver,
				},
			}
		}
	{{end}}
{{end}}
{{range $ext := .Model.ObjectExtensions}}
	{{$obj := $ext.Object}}
	{{if not $ext.ExtendsLocalObject}}
		// {{$obj.Name}}Resolver struct
		type {{$obj.Name}}Resolver struct {
			*gen.Generated{{$obj.Name}}Resolver
		}
		{{if or $obj.HasAnyRelationships $obj.HasReadonlyColumns $ext.HasAnyNonExternalField}}
			// {{$obj.Name}} ...
			func (r *Resolver) {{$obj.Name}}() gen.{{$obj.Name}}Resolver {
				return &{{$obj.Name}}Resolver{
					Generated{{$obj.Name}}Resolver: &gen.Generated{{$obj.Name}}Resolver{
						GeneratedResolver: r.GeneratedResolver,
					},
				}
			}
		{{end}}
	{{end}}
{{end}}
`
