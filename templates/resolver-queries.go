package templates

// ResolverQueries ...
var ResolverQueries = `package gen

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
	"github.com/loopcontext/go-graphql-orm/events"
	"github.com/vektah/gqlparser/v2/ast"
)

// GeneratedQueryResolver struct
type GeneratedQueryResolver struct{ *GeneratedResolver }

{{range $obj := .Model.ObjectEntities}}
	// Query{{$obj.Name}}HandlerOptions struct
	type Query{{$obj.Name}}HandlerOptions struct {
		ID *string
		Q      *string
		Filter *{{$obj.Name}}FilterType
	}

	// {{$obj.Name}} ...
	func (r *GeneratedQueryResolver) {{$obj.Name}}(ctx context.Context, id *string, q *string, filter *{{$obj.Name}}FilterType) (*{{$obj.Name}}, error) {
		opts := Query{{$obj.Name}}HandlerOptions{
			ID: id,
			Q: q,
			Filter: filter,
		}
		return r.Handlers.Query{{$obj.Name}}(ctx, r.GeneratedResolver, opts)
	}
	// Query{{$obj.Name}}Handler handler
	func Query{{$obj.Name}}Handler(ctx context.Context, r *GeneratedResolver, opts Query{{$obj.Name}}HandlerOptions) (*{{$obj.Name}}, error) {
		selection := []ast.Selection{}
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			selection = append(selection, f.Field)
		}
		selectionSet := ast.SelectionSet(selection)

		query := {{$obj.Name}}QueryFilter{opts.Q}
		offset := 0
		limit := 1
		rt := &{{$obj.Name}}ResultType{
			EntityResultType: EntityResultType{
				Offset: &offset,
				Limit:  &limit,
				Query:  &query,
				Filter: opts.Filter,
				SelectionSet: &selectionSet,
			},
		}
		qb := r.GetDB(ctx)
		if opts.ID != nil {
			qb = qb.Where(TableName("{{$obj.TableName}}") + ".id = ?", *opts.ID)
		}

		var items []*{{$obj.Name}}
		giOpts := GetItemsOptions{
			Alias:TableName("{{$obj.TableName}}"),
			Preloaders:[]string{ {{range $r := $obj.PreloadableRelationships}}
				"{{$r.MethodName}}",{{end}}
			},
		}
		err := rt.GetItems(ctx, qb, giOpts, &items)
		if err != nil {
			return nil, err
		}
		if len(items) == 0 {
			return nil, nil
		}
		return items[0], err
	}

	// Query{{$obj.PluralName}}HandlerOptions struct
	type Query{{$obj.PluralName}}HandlerOptions struct {
		Offset *int
		Limit  *int
		Q      *string
		Sort   []*{{$obj.Name}}SortType
		Filter *{{$obj.Name}}FilterType
	}
	// {{$obj.PluralName}} handler options
	func (r *GeneratedQueryResolver) {{$obj.PluralName}}(ctx context.Context, offset *int, limit *int, q *string, sort []*{{$obj.Name}}SortType, filter *{{$obj.Name}}FilterType) (*{{$obj.Name}}ResultType, error) {
		opts := Query{{$obj.PluralName}}HandlerOptions{
			Offset: offset,
			Limit: limit,
			Q: q,
			Sort: sort,
			Filter: filter,
		}
		return r.Handlers.Query{{$obj.PluralName}}(ctx, r.GeneratedResolver, opts)
	}
	// {{$obj.PluralName}}Items handler
	func (r *GeneratedResolver) {{$obj.PluralName}}Items(ctx context.Context, opts Query{{$obj.PluralName}}HandlerOptions) (res []*{{$obj.Name}}, err error) {
		resultType, err := r.Handlers.Query{{$obj.PluralName}}(ctx, r, opts)
		if err != nil {
			return
		}
		err = resultType.GetItems(ctx, r.GetDB(ctx), GetItemsOptions{
			Alias: TableName("{{$obj.TableName}}"),
		}, &res)
		if err != nil {
			return
		}
		return
	}
	// {{$obj.PluralName}}Count handler
	func (r *GeneratedResolver) {{$obj.PluralName}}Count(ctx context.Context, opts Query{{$obj.PluralName}}HandlerOptions) (count int, err error) {
		resultType, err := r.Handlers.Query{{$obj.PluralName}}(ctx, r, opts)
		if err != nil {
			return
		}
		return resultType.GetCount(ctx, r.GetDB(ctx), GetItemsOptions{
			Alias: TableName("{{$obj.TableName}}"),
		}, &{{$obj.Name}}{})
	}
	// Query{{$obj.PluralName}}Handler handler
	func Query{{$obj.PluralName}}Handler(ctx context.Context, r *GeneratedResolver, opts Query{{$obj.PluralName}}HandlerOptions) (*{{$obj.Name}}ResultType, error) {
		query := {{$obj.Name}}QueryFilter{opts.Q}

		var selectionSet *ast.SelectionSet
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			if f.Field.Name == "items" {
				selectionSet = &f.Field.SelectionSet
			}
		}

		_sort := []EntitySort{}
		for _, sort := range opts.Sort {
			_sort = append(_sort, sort)
		}

		return &{{$obj.Name}}ResultType{
			EntityResultType: EntityResultType{
				Offset: opts.Offset,
				Limit:  opts.Limit,
				Query:  &query,
				Sort: _sort,
				Filter: opts.Filter,
				SelectionSet: selectionSet,
			},
		}, nil
	}

	// Generated{{$obj.Name}}ResultTypeResolver struct
	type Generated{{$obj.Name}}ResultTypeResolver struct{ *GeneratedResolver }

	// Items ...
	func (r *Generated{{$obj.Name}}ResultTypeResolver) Items(ctx context.Context, obj *{{$obj.Name}}ResultType) (items []*{{$obj.Name}}, err error) {
		opts := GetItemsOptions{
			Alias:TableName("{{$obj.TableName}}"),
			Preloaders:[]string{ {{range $r := $obj.PreloadableRelationships}}
				"{{$r.MethodName}}",{{end}}
			},
		}
		err = obj.GetItems(ctx, r.GetDB(ctx), opts, &items)
		{{if $obj.HasPreloadableRelationships}}
			for _, item := range items {
				{{range $rel := $obj.PreloadableRelationships}}
				item.{{$rel.MethodName}}Preloaded = true{{end}}
			}
		{{end}}

		uniqueItems := []*{{$obj.Name}}{}
		idMap := map[string]bool{}
		for _, item := range items {
			if _,ok := idMap[item.ID]; !ok {
				idMap[item.ID] = true
				uniqueItems = append(uniqueItems, item)
			}
		}
		items = uniqueItems
		return
	}

	// Count ...
	func (r *Generated{{$obj.Name}}ResultTypeResolver) Count(ctx context.Context, obj *{{$obj.Name}}ResultType) (count int, err error) {
		opts := GetItemsOptions{
			Alias:TableName("{{$obj.TableName}}"),
			Preloaders:[]string{ {{range $r := $obj.PreloadableRelationships}}
				"{{$r.MethodName}}",{{end}}
			},
		}
		return obj.GetCount(ctx, r.GetDB(ctx), opts, &{{$obj.Name}}{})
	}

	{{if $obj.HasAggregableColumn}}
	// Aggregations ...
	func (r *Generated{{$obj.Name}}ResultTypeResolver) Aggregations(ctx context.Context, obj *{{$obj.Name}}ResultType) (res *{{$obj.Name}}ResultAggregations, err error) {
		aggregationsMap := map[string]GetAggregationsAggregationField{
			{{range $agg := $obj.AggregationsByField}}
			"{{$agg.FieldName}}": GetAggregationsAggregationField{"{{$agg.Field}}","{{$agg.Name}}"},{{end}}
		}
		aggregationFieldsMap := map[string]string{
			{{range $agg := $obj.AggregationsByField}}
			"{{$agg.FieldName}}": "{{$agg.Field}}",{{end}}
		}
		fieldsMap := map[string]struct{}{}
		fields := []string{}
		aggregationFields := []GetAggregationsAggregationField{}
		for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
			aggregationFields = append(aggregationFields, aggregationsMap[f.Field.Name])
			fieldsMap[aggregationFieldsMap[f.Field.Name]] = struct{}{}
		}
		for key, _ := range fieldsMap {
			fields = append(fields, key)
		}

		opts := GetAggregationsOptions{
			Alias:             TableName("{{$obj.TableName}}"),
			Fields:            fields,
			AggregationFields: aggregationFields,
		}
		res = &{{$obj.Name}}ResultAggregations{}
		err = obj.GetAggregations(ctx, r.GetDB(ctx), opts, &{{$obj.Name}}{}, res)
		return
	}
	{{end}}

	{{if $obj.NeedsQueryResolver}}
		// Generated{{$obj.Name}}Resolver struct
		type Generated{{$obj.Name}}Resolver struct { *GeneratedResolver }

		{{range $col := $obj.Fields}}
			{{if $col.NeedsQueryResolver}}
				// {{$col.MethodName}} ...
				func (r *Generated{{$obj.Name}}Resolver) {{$col.MethodName}}(ctx context.Context, obj *{{$obj.Name}}) (res {{$col.GoResultType}}, err error) {
					return r.Handlers.{{$obj.Name}}{{$col.MethodName}}(ctx, r.GeneratedResolver, obj)
				}
				// {{$obj.Name}}.{{$col.MethodName}}Handler handler
				func {{$obj.Name}}{{$col.MethodName}}Handler(ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (res {{$col.GoResultType}}, err error) {
					{{if and (not $col.IsList) $col.HasTargetTypeWithIDField ($obj.HasColumn (print $col.Name "Id"))}}
						if obj.{{$col.MethodName}}ID != nil {
							res = &{{$col.TargetType}}{ID: *obj.{{$col.MethodName}}ID}
						}
					{{else}}
						{{if $col.IsEmbeddedColumn}}
							{{if $col.IsOptional}}
								if obj.{{$col.MethodName}} != nil && *obj.{{$col.MethodName}} != "" {
									err = json.Unmarshal([]byte(*obj.{{$col.MethodName}}), &res)
								}
							{{else}}
								if obj.{{$col.MethodName}} != "" {
									err = json.Unmarshal([]byte(obj.{{$col.MethodName}}), &res)
								}
							{{end}}
						{{else}}
							err = fmt.Errorf("Resolver handler for {{$obj.Name}}{{$col.MethodName}} not implemented")
						{{end}}
					{{end}}
					return
				}
			{{end}}
		{{end}}

		{{range $index, $rel := $obj.Relationships}}
			// {{$rel.MethodName}} ...
			func (r *Generated{{$obj.Name}}Resolver) {{$rel.MethodName}}(ctx context.Context, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error) {
				return r.Handlers.{{$obj.Name}}{{$rel.MethodName}}(ctx, r.GeneratedResolver, obj)
			}
			// {{$obj.Name}}.{{$rel.MethodName}}Handler handler
			func {{$obj.Name}}{{$rel.MethodName}}Handler(ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error) {
				{{if $rel.Preload}}
				if obj.{{$rel.MethodName}}Preloaded {
					res = obj.{{$rel.MethodName}}
				}else {
				{{end}}
					{{if $rel.IsToMany}}
							items := []*{{$rel.TargetType}}{}
							db := r.GetDB(ctx)
							err = db.Model(obj).Related(&items, "{{$rel.MethodName}}").Error
							res = items
					{{else}}
						loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
						if obj.{{$rel.MethodName}}ID != nil {
							item, _err := loaders["{{$rel.Target.Name}}"].Load(ctx, dataloader.StringKey(*obj.{{$rel.MethodName}}ID))()
							res, _ = item.({{$rel.ReturnType}})
							{{if $rel.IsNonNull}}
							if res == nil {
								_err = fmt.Errorf("{{$rel.Target.Name}} with id '%s' not found",*obj.{{$rel.MethodName}}ID)
							}{{end}}
							err = _err
						}
					{{end}}
				{{if $rel.Preload}}
				}
				{{end}}
				return
			}
			{{if $rel.IsToMany}}
				// {{$rel.MethodName}}Ids ...
				func (r *Generated{{$obj.Name}}Resolver) {{$rel.MethodName}}Ids(ctx context.Context, obj *{{$obj.Name}}) (ids []string, err error) {
					ids = []string{}

					items := []*{{$rel.TargetType}}{}
					db := r.GetDB(ctx)
					err = db.Model(obj).Select(TableName("{{$rel.Target.TableName}}") + ".id").Related(&items, "{{$rel.MethodName}}").Error

					for _, item := range items {
						ids = append(ids, item.ID)
					}

					return
				}

				// {{$rel.MethodName}}Connection method
				func (r *Generated{{$obj.Name}}Resolver) {{$rel.MethodName}}Connection(ctx context.Context, obj *{{$obj.Name}}, offset *int, limit *int, q *string, sort []*{{$rel.TargetType}}SortType, filter *{{$rel.TargetType}}FilterType) (res *{{$rel.TargetType}}ResultType, err error) {
					f := &{{$rel.TargetType}}FilterType{
						{{$rel.InverseRelationship.MethodName}}: &{{$obj.Name}}FilterType{
							ID: &obj.ID,
						},
					}
					if filter == nil {
						filter = f
					} else {
						filter = &{{$rel.TargetType}}FilterType{
							And: []*{{$rel.TargetType}}FilterType{
								filter,
								f,
							},
						}
					}
					opts := Query{{$rel.Target.PluralName}}HandlerOptions{
						Offset: offset,
						Limit:  limit,
						Q:      q,
						Sort:   sort,
						Filter: filter,
					}
					return r.Handlers.Query{{$rel.Target.PluralName}}(ctx, r.GeneratedResolver, opts)
				}
			{{end}}

		{{end}}

	{{end}}

{{end}}
`
