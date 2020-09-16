package templates

// ResolverSrcExt template
var ResolverSrcExt = `package src

import (
	"context"
	"{{.Config.Package}}/gen"
	"github.com/rs/zerolog/log"
)

{{range $obj := .Model.ObjectEntities}}
// {{$obj.PluralName}} method
func (r *QueryResolver) {{$obj.PluralName}}(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.{{$obj.Name}}SortType, filter *gen.{{$obj.Name}}FilterType) (*gen.{{$obj.Name}}ResultType, error) {
	log.Debug().Msgf("JWT Claims from Context: %#v", gen.GetJWTClaimsFromContext(ctx))
	// Check permissions, role, or whatever from JWT claims from context
	return r.GeneratedQueryResolver.{{$obj.PluralName}}(ctx, offset, limit, q, sort, filter)
}

// Create{{$obj.Name}} method
func (r *MutationResolver) Create{{$obj.Name}}(ctx context.Context, input map[string]interface{}) (item *gen.{{$obj.Name}}, err error) {
	log.Debug().Msgf("Creating {{$obj.Name}} if I'm allowed to - %#v", input)
	return r.GeneratedMutationResolver.Create{{$obj.Name}}(ctx, input)
}

// Update{{$obj.Name}} method
func (r *MutationResolver) Update{{$obj.Name}}(ctx context.Context, id string, input map[string]interface{}) (item *gen.{{$obj.Name}}, err error) {
	log.Debug().Msgf("Updating {{$obj.Name}} if I'm allowed to | %s | %#v", id, input)
	return r.GeneratedMutationResolver.Update{{$obj.Name}}(ctx, id, input)
}

// Delete{{$obj.Name}} method
func (r *MutationResolver) Delete{{$obj.Name}}(ctx context.Context, id string) (item *gen.{{$obj.Name}}, err error) {
	log.Debug().Msgf("Deleting {{$obj.Name}} if I'm allowed to - %s", id)
	return r.GeneratedMutationResolver.Delete{{$obj.Name}}(ctx, id)
}

// DeleteAll{{$obj.PluralName}} method
func (r *MutationResolver) DeleteAll{{$obj.PluralName}}(ctx context.Context) (ok bool, err error) {
	log.Debug().Msg("Deleting {{$obj.PluralName}} if I'm allowed to")
	return r.GeneratedMutationResolver.DeleteAll{{$obj.PluralName}}(ctx)
}

{{end}}
`
