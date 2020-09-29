package src

import (
	"context"
	"fmt"

	"github.com/loopcontext/go-graphql-orm/test/gen"
)

const (
	jwtTokenPermissionErrMsg = "You don't have permission to %s the entity %s"
)

// Tasks method
func (r *QueryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.TaskSortType, filter *gen.TaskFilterType) (*gen.TaskResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "tasks", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "tasks")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Tasks to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Tasks(ctx, offset, limit, q, sort, filter)
}

// CreateTask method
func (r *MutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *gen.Task, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "tasks", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "tasks")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Tasks to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateTask(ctx, input)
}

// UpdateTask method
func (r *MutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *gen.Task, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "tasks", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "tasks")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Tasks to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateTask(ctx, id, input)
}

// DeleteTask method
func (r *MutationResolver) DeleteTask(ctx context.Context, id string) (item *gen.Task, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "tasks", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "tasks")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Tasks to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteTask(ctx, id)
}

// DeleteAllTasks method
func (r *MutationResolver) DeleteAllTasks(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "tasks", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "tasks")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Tasks to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllTasks(ctx)
}

// TaskCategories method
func (r *QueryResolver) TaskCategories(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.TaskCategorySortType, filter *gen.TaskCategoryFilterType) (*gen.TaskCategoryResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "task_categories", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "task_categories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope TaskCategories to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.TaskCategories(ctx, offset, limit, q, sort, filter)
}

// CreateTaskCategory method
func (r *MutationResolver) CreateTaskCategory(ctx context.Context, input map[string]interface{}) (item *gen.TaskCategory, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "task_categories", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "task_categories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope TaskCategories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateTaskCategory(ctx, input)
}

// UpdateTaskCategory method
func (r *MutationResolver) UpdateTaskCategory(ctx context.Context, id string, input map[string]interface{}) (item *gen.TaskCategory, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "task_categories", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "task_categories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope TaskCategories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateTaskCategory(ctx, id, input)
}

// DeleteTaskCategory method
func (r *MutationResolver) DeleteTaskCategory(ctx context.Context, id string) (item *gen.TaskCategory, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "task_categories", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "task_categories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope TaskCategories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteTaskCategory(ctx, id)
}

// DeleteAllTaskCategories method
func (r *MutationResolver) DeleteAllTaskCategories(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "task_categories", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "task_categories")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope TaskCategories to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllTaskCategories(ctx)
}

// Companies method
func (r *QueryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.CompanySortType, filter *gen.CompanyFilterType) (*gen.CompanyResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "companies", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "companies")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Companies to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Companies(ctx, offset, limit, q, sort, filter)
}

// CreateCompany method
func (r *MutationResolver) CreateCompany(ctx context.Context, input map[string]interface{}) (item *gen.Company, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "companies", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "companies")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Companies to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateCompany(ctx, input)
}

// UpdateCompany method
func (r *MutationResolver) UpdateCompany(ctx context.Context, id string, input map[string]interface{}) (item *gen.Company, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "companies", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "companies")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Companies to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateCompany(ctx, id, input)
}

// DeleteCompany method
func (r *MutationResolver) DeleteCompany(ctx context.Context, id string) (item *gen.Company, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "companies", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "companies")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Companies to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteCompany(ctx, id)
}

// DeleteAllCompanies method
func (r *MutationResolver) DeleteAllCompanies(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "companies", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "companies")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Companies to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllCompanies(ctx)
}

// Users method
func (r *QueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.UserSortType, filter *gen.UserFilterType) (*gen.UserResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.Users(ctx, offset, limit, q, sort, filter)
}

// CreateUser method
func (r *MutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *gen.User, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreateUser(ctx, input)
}

// UpdateUser method
func (r *MutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *gen.User, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdateUser(ctx, id, input)
}

// DeleteUser method
func (r *MutationResolver) DeleteUser(ctx context.Context, id string) (item *gen.User, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteUser(ctx, id)
}

// DeleteAllUsers method
func (r *MutationResolver) DeleteAllUsers(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "users", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "users")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope Users to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllUsers(ctx)
}

// PlainEntities method
func (r *QueryResolver) PlainEntities(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.PlainEntitySortType, filter *gen.PlainEntityFilterType) (*gen.PlainEntityResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "plain_entities", gen.JWTPermissionConstList) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "plain_entities")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PlainEntities to current jwtClaims.Subject (User.ID)
	return r.GeneratedQueryResolver.PlainEntities(ctx, offset, limit, q, sort, filter)
}

// CreatePlainEntity method
func (r *MutationResolver) CreatePlainEntity(ctx context.Context, input map[string]interface{}) (item *gen.PlainEntity, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "plain_entities", gen.JWTPermissionConstCreate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "plain_entities")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PlainEntities to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.CreatePlainEntity(ctx, input)
}

// UpdatePlainEntity method
func (r *MutationResolver) UpdatePlainEntity(ctx context.Context, id string, input map[string]interface{}) (item *gen.PlainEntity, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "plain_entities", gen.JWTPermissionConstUpdate) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "plain_entities")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PlainEntities to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.UpdatePlainEntity(ctx, id, input)
}

// DeletePlainEntity method
func (r *MutationResolver) DeletePlainEntity(ctx context.Context, id string) (item *gen.PlainEntity, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasPermission(jwtClaims, "plain_entities", gen.JWTPermissionConstDelete) {
		return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "plain_entities")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PlainEntities to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeletePlainEntity(ctx, id)
}

// DeleteAllPlainEntities method
func (r *MutationResolver) DeleteAllPlainEntities(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if !gen.HasRole(jwtClaims, "admin") &&
		!gen.HasPermission(jwtClaims, "plain_entities", gen.JWTPermissionConstDelete) {
		return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "plain_entities")
	}
	// TODO: Insert here any code ETL on your query/mutation for example: scope PlainEntities to current jwtClaims.Subject (User.ID)
	return r.GeneratedMutationResolver.DeleteAllPlainEntities(ctx)
}
