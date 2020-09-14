package src

import (
	"context"

	"github.com/loopcontext/go-graphql-orm/test/gen"
)

// NewResolver ...
func NewResolver(db *gen.DB, ec *gen.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
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
func (r *MutationResolver) BeginTransaction(ctx context.Context, fn func(context.Context) error) error {
	ctx = gen.EnrichContextWithMutations(ctx, r.GeneratedResolver)
	err := fn(ctx)
	if err != nil {
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

// TaskResultTypeResolver struct
type TaskResultTypeResolver struct {
	*gen.GeneratedTaskResultTypeResolver
}

// TaskResultType ...
func (r *Resolver) TaskResultType() gen.TaskResultTypeResolver {
	return &TaskResultTypeResolver{
		GeneratedTaskResultTypeResolver: &gen.GeneratedTaskResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// TaskResolver struct
type TaskResolver struct {
	*gen.GeneratedTaskResolver
}

// Task ...
func (r *Resolver) Task() gen.TaskResolver {
	return &TaskResolver{
		GeneratedTaskResolver: &gen.GeneratedTaskResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// TaskCategoryResultTypeResolver struct
type TaskCategoryResultTypeResolver struct {
	*gen.GeneratedTaskCategoryResultTypeResolver
}

// TaskCategoryResultType ...
func (r *Resolver) TaskCategoryResultType() gen.TaskCategoryResultTypeResolver {
	return &TaskCategoryResultTypeResolver{
		GeneratedTaskCategoryResultTypeResolver: &gen.GeneratedTaskCategoryResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// TaskCategoryResolver struct
type TaskCategoryResolver struct {
	*gen.GeneratedTaskCategoryResolver
}

// TaskCategory ...
func (r *Resolver) TaskCategory() gen.TaskCategoryResolver {
	return &TaskCategoryResolver{
		GeneratedTaskCategoryResolver: &gen.GeneratedTaskCategoryResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// CompanyResultTypeResolver struct
type CompanyResultTypeResolver struct {
	*gen.GeneratedCompanyResultTypeResolver
}

// CompanyResultType ...
func (r *Resolver) CompanyResultType() gen.CompanyResultTypeResolver {
	return &CompanyResultTypeResolver{
		GeneratedCompanyResultTypeResolver: &gen.GeneratedCompanyResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// CompanyResolver struct
type CompanyResolver struct {
	*gen.GeneratedCompanyResolver
}

// Company ...
func (r *Resolver) Company() gen.CompanyResolver {
	return &CompanyResolver{
		GeneratedCompanyResolver: &gen.GeneratedCompanyResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// UserResultTypeResolver struct
type UserResultTypeResolver struct {
	*gen.GeneratedUserResultTypeResolver
}

// UserResultType ...
func (r *Resolver) UserResultType() gen.UserResultTypeResolver {
	return &UserResultTypeResolver{
		GeneratedUserResultTypeResolver: &gen.GeneratedUserResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// UserResolver struct
type UserResolver struct {
	*gen.GeneratedUserResolver
}

// User ...
func (r *Resolver) User() gen.UserResolver {
	return &UserResolver{
		GeneratedUserResolver: &gen.GeneratedUserResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PlainEntityResultTypeResolver struct
type PlainEntityResultTypeResolver struct {
	*gen.GeneratedPlainEntityResultTypeResolver
}

// PlainEntityResultType ...
func (r *Resolver) PlainEntityResultType() gen.PlainEntityResultTypeResolver {
	return &PlainEntityResultTypeResolver{
		GeneratedPlainEntityResultTypeResolver: &gen.GeneratedPlainEntityResultTypeResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// PlainEntityResolver struct
type PlainEntityResolver struct {
	*gen.GeneratedPlainEntityResolver
}

// PlainEntity ...
func (r *Resolver) PlainEntity() gen.PlainEntityResolver {
	return &PlainEntityResolver{
		GeneratedPlainEntityResolver: &gen.GeneratedPlainEntityResolver{
			GeneratedResolver: r.GeneratedResolver,
		},
	}
}

// ReviewResolver struct
type ReviewResolver struct {
	*gen.GeneratedReviewResolver
}

// CountryResolver struct
type CountryResolver struct {
	*gen.GeneratedCountryResolver
}
