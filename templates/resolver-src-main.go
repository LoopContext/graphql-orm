package templates

// ResolverSrc ...
var ResolverSrc = `package src

import (
	"context"
	"{{.Config.Package}}/gen"
	"github.com/loopcontext/graphql-orm/events"
	"golang.org/x/crypto/bcrypt"
	//"github.com/loopcontext/checkmail"
)

// New ...
func New(db *gen.DB, ec *gen.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedMutationResolver, input map[string]interface{}) (item *gen.Company, err error) {
		// Before save

		// Verify email, for example.
		// err = checkmail.ValidateFormat(item.Email)
		// if err != nil {
		// 	return nil, err
		// }
		return gen.CreateUserHandler(ctx, r, input)
	}
	resolver.Handlers.OnEvent = func(ctx context.Context, r *gen.GeneratedResolver, e *events.Event) (err error) {
		// After save
		if e.Entity == "User" && (e.Type == events.EventTypeCreated || e.Type == events.EventTypeUpdated) {
			// do something...
		}
		return nil
	}

	return resolver
}


// You can extend QueryResolver for adding custom fields in schema
// // Hello world
// func (r *QueryResolver) Hello(ctx context.Context) (string, error) {
// 	return "world", nil
// }
`
