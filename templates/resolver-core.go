package templates

// ResolverCore ...
var ResolverCore = `package gen

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
	"github.com/loopcontext/go-graphql-orm/events"
	"github.com/vektah/gqlparser/v2/ast"
)

// ResolutionHandlers struct
type ResolutionHandlers struct {
	OnEvent        func(ctx context.Context, r *GeneratedResolver, e *events.Event) error
	{{range $obj := .Model.ObjectEntities}}
		Create{{$obj.Name}} func (ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *{{$obj.Name}}, err error)
		Update{{$obj.Name}} func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *{{$obj.Name}}, err error)
		Delete{{$obj.Name}} func(ctx context.Context, r *GeneratedResolver, id string) (item *{{$obj.Name}}, err error)
		DeleteAll{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver) (bool, error)
		Query{{$obj.Name}} func (ctx context.Context, r *GeneratedResolver, opts Query{{$obj.Name}}HandlerOptions) (*{{$obj.Name}}, error)
		Query{{$obj.PluralName}} func (ctx context.Context, r *GeneratedResolver, opts Query{{$obj.PluralName}}HandlerOptions) (*{{$obj.Name}}ResultType, error)
		{{range $col := $obj.Fields}}{{if $col.NeedsQueryResolver}}
			{{$obj.Name}}{{$col.MethodName}} func (ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (res {{$col.GoResultType}}, err error)
		{{end}}{{end}}
		{{range $rel := $obj.Relationships}}
			{{$obj.Name}}{{$rel.MethodName}} func (ctx context.Context,r *GeneratedResolver, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error)
		{{end}}
	{{end}}
}

// DefaultResolutionHandlers ...
func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },
		{{range $obj := .Model.ObjectEntities}}
			Create{{$obj.Name}}: Create{{$obj.Name}}Handler,
			Update{{$obj.Name}}: Update{{$obj.Name}}Handler,
			Delete{{$obj.Name}}: Delete{{$obj.Name}}Handler,
			DeleteAll{{$obj.PluralName}}: DeleteAll{{$obj.PluralName}}Handler,
			Query{{$obj.Name}}: Query{{$obj.Name}}Handler,
			Query{{$obj.PluralName}}: Query{{$obj.PluralName}}Handler,
			{{range $col := $obj.Fields}}{{if $col.NeedsQueryResolver}}
				{{$obj.Name}}{{$col.MethodName}}: {{$obj.Name}}{{$col.MethodName}}Handler,
			{{end}}{{end}}
			{{range $rel := $obj.Relationships}}
				{{$obj.Name}}{{$rel.MethodName}}: {{$obj.Name}}{{$rel.MethodName}}Handler,
			{{end}}
		{{end}}
	}
	return handlers
}

// GeneratedResolver struct
type GeneratedResolver struct {
	Handlers ResolutionHandlers
	db *DB
	EventController *EventController
}

// NewGeneratedResolver ...
func NewGeneratedResolver(handlers ResolutionHandlers, db *DB, ec *EventController) *GeneratedResolver {
	return &GeneratedResolver{Handlers: handlers, db: db, EventController: ec}
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.db.Query()
	}
	return db
}
`
