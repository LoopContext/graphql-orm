package gen

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/graph-gophers/dataloader"
	"github.com/vektah/gqlparser/v2/ast"
)

// GeneratedQueryResolver struct
type GeneratedQueryResolver struct{ *GeneratedResolver }

// QueryTaskHandlerOptions struct
type QueryTaskHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *TaskFilterType
}

// Task ...
func (r *GeneratedQueryResolver) Task(ctx context.Context, id *string, q *string, filter *TaskFilterType) (*Task, error) {
	opts := QueryTaskHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryTask(ctx, r.GeneratedResolver, opts)
}

// QueryTaskHandler handler
func QueryTaskHandler(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TaskQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &TaskResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("tasks")+".id = ?", *opts.ID)
	}

	var items []*Task
	giOpts := GetItemsOptions{
		Alias: TableName("tasks"),
		Preloaders: []string{
			"Assignee",
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

// QueryTasksHandlerOptions struct
type QueryTasksHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*TaskSortType
	Filter *TaskFilterType
}

// Tasks ...
func (r *GeneratedQueryResolver) Tasks(ctx context.Context, offset *int, limit *int, q *string, sort []*TaskSortType, filter *TaskFilterType) (*TaskResultType, error) {
	opts := QueryTasksHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}

// QueryTasksHandler handler
func QueryTasksHandler(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error) {
	query := TaskQueryFilter{opts.Q}

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

	return &TaskResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedTaskResultTypeResolver struct
type GeneratedTaskResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedTaskResultTypeResolver) Items(ctx context.Context, obj *TaskResultType) (items []*Task, err error) {
	otps := GetItemsOptions{
		Alias: TableName("tasks"),
		Preloaders: []string{
			"Assignee",
		},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	for _, item := range items {

		item.AssigneePreloaded = true
	}

	uniqueItems := []*Task{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedTaskResultTypeResolver) Count(ctx context.Context, obj *TaskResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias: TableName("tasks"),
		Preloaders: []string{
			"Assignee",
		},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Task{})
}

// GeneratedTaskResolver struct
type GeneratedTaskResolver struct{ *GeneratedResolver }

// Metas ...
func (r *GeneratedTaskResolver) Metas(ctx context.Context, obj *Task) (res []*TaskMeta, err error) {
	return r.Handlers.TaskMetas(ctx, r.GeneratedResolver, obj)
}

// TaskMetasHandler handler
func TaskMetasHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res []*TaskMeta, err error) {

	if obj.Metas != nil && *obj.Metas != "" {
		err = json.Unmarshal([]byte(*obj.Metas), &res)
	}

	return
}

// Meta ...
func (r *GeneratedTaskResolver) Meta(ctx context.Context, obj *Task) (res *TaskMeta, err error) {
	return r.Handlers.TaskMeta(ctx, r.GeneratedResolver, obj)
}

// TaskMetaHandler handler
func TaskMetaHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res *TaskMeta, err error) {

	if obj.Meta != nil && *obj.Meta != "" {
		err = json.Unmarshal([]byte(*obj.Meta), &res)
	}

	return
}

// Assignee ...
func (r *GeneratedTaskResolver) Assignee(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskAssignee(ctx, r.GeneratedResolver, obj)
}

// TaskAssigneeHandler handler
func TaskAssigneeHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error) {

	if obj.AssigneePreloaded {
		res = obj.Assignee
	} else {

		loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
		if obj.AssigneeID != nil {
			item, _err := loaders["User"].Load(ctx, dataloader.StringKey(*obj.AssigneeID))()
			res, _ = item.(*User)

			err = _err
		}

	}

	return
}

// Owner ...
func (r *GeneratedTaskResolver) Owner(ctx context.Context, obj *Task) (res *User, err error) {
	return r.Handlers.TaskOwner(ctx, r.GeneratedResolver, obj)
}

// TaskOwnerHandler handler
func TaskOwnerHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.OwnerID != nil {
		item, _err := loaders["User"].Load(ctx, dataloader.StringKey(*obj.OwnerID))()
		res, _ = item.(*User)

		if res == nil {
			_err = fmt.Errorf("User with id '%s' not found", *obj.OwnerID)
		}
		err = _err
	}

	return
}

// ParentTask ...
func (r *GeneratedTaskResolver) ParentTask(ctx context.Context, obj *Task) (res *Task, err error) {
	return r.Handlers.TaskParentTask(ctx, r.GeneratedResolver, obj)
}

// TaskParentTaskHandler handler
func TaskParentTaskHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res *Task, err error) {

	loaders := ctx.Value(KeyLoaders).(map[string]*dataloader.Loader)
	if obj.ParentTaskID != nil {
		item, _err := loaders["Task"].Load(ctx, dataloader.StringKey(*obj.ParentTaskID))()
		res, _ = item.(*Task)

		err = _err
	}

	return
}

// Subtasks ...
func (r *GeneratedTaskResolver) Subtasks(ctx context.Context, obj *Task) (res []*Task, err error) {
	return r.Handlers.TaskSubtasks(ctx, r.GeneratedResolver, obj)
}

// TaskSubtasksHandler handler
func TaskSubtasksHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res []*Task, err error) {

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Subtasks").Error
	res = items

	return
}

// SubtasksIds ...
func (r *GeneratedTaskResolver) SubtasksIds(ctx context.Context, obj *Task) (ids []string, err error) {
	ids = []string{}

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("tasks")+".id").Related(&items, "Subtasks").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// SubtasksConnection method
func (r *GeneratedTaskResolver) SubtasksConnection(ctx context.Context, obj *Task, offset *int, limit *int, q *string, sort []*TaskSortType, filter *TaskFilterType) (res *TaskResultType, err error) {
	f := &TaskFilterType{
		ParentTask: &TaskFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &TaskFilterType{
			And: []*TaskFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryTasksHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}

// Categories ...
func (r *GeneratedTaskResolver) Categories(ctx context.Context, obj *Task) (res []*TaskCategory, err error) {
	return r.Handlers.TaskCategories(ctx, r.GeneratedResolver, obj)
}

// TaskCategoriesHandler handler
func TaskCategoriesHandler(ctx context.Context, r *GeneratedResolver, obj *Task) (res []*TaskCategory, err error) {

	items := []*TaskCategory{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Categories").Error
	res = items

	return
}

// CategoriesIds ...
func (r *GeneratedTaskResolver) CategoriesIds(ctx context.Context, obj *Task) (ids []string, err error) {
	ids = []string{}

	items := []*TaskCategory{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("task_categories")+".id").Related(&items, "Categories").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// CategoriesConnection method
func (r *GeneratedTaskResolver) CategoriesConnection(ctx context.Context, obj *Task, offset *int, limit *int, q *string, sort []*TaskCategorySortType, filter *TaskCategoryFilterType) (res *TaskCategoryResultType, err error) {
	f := &TaskCategoryFilterType{
		Tasks: &TaskFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &TaskCategoryFilterType{
			And: []*TaskCategoryFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryTaskCategoriesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTaskCategories(ctx, r.GeneratedResolver, opts)
}

// QueryTaskCategoryHandlerOptions struct
type QueryTaskCategoryHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *TaskCategoryFilterType
}

// TaskCategory ...
func (r *GeneratedQueryResolver) TaskCategory(ctx context.Context, id *string, q *string, filter *TaskCategoryFilterType) (*TaskCategory, error) {
	opts := QueryTaskCategoryHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryTaskCategory(ctx, r.GeneratedResolver, opts)
}

// QueryTaskCategoryHandler handler
func QueryTaskCategoryHandler(ctx context.Context, r *GeneratedResolver, opts QueryTaskCategoryHandlerOptions) (*TaskCategory, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := TaskCategoryQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &TaskCategoryResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("task_categories")+".id = ?", *opts.ID)
	}

	var items []*TaskCategory
	giOpts := GetItemsOptions{
		Alias:      TableName("task_categories"),
		Preloaders: []string{},
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

// QueryTaskCategoriesHandlerOptions struct
type QueryTaskCategoriesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*TaskCategorySortType
	Filter *TaskCategoryFilterType
}

// TaskCategories ...
func (r *GeneratedQueryResolver) TaskCategories(ctx context.Context, offset *int, limit *int, q *string, sort []*TaskCategorySortType, filter *TaskCategoryFilterType) (*TaskCategoryResultType, error) {
	opts := QueryTaskCategoriesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTaskCategories(ctx, r.GeneratedResolver, opts)
}

// QueryTaskCategoriesHandler handler
func QueryTaskCategoriesHandler(ctx context.Context, r *GeneratedResolver, opts QueryTaskCategoriesHandlerOptions) (*TaskCategoryResultType, error) {
	query := TaskCategoryQueryFilter{opts.Q}

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

	return &TaskCategoryResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedTaskCategoryResultTypeResolver struct
type GeneratedTaskCategoryResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedTaskCategoryResultTypeResolver) Items(ctx context.Context, obj *TaskCategoryResultType) (items []*TaskCategory, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("task_categories"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*TaskCategory{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedTaskCategoryResultTypeResolver) Count(ctx context.Context, obj *TaskCategoryResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("task_categories"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &TaskCategory{})
}

// GeneratedTaskCategoryResolver struct
type GeneratedTaskCategoryResolver struct{ *GeneratedResolver }

// Tasks ...
func (r *GeneratedTaskCategoryResolver) Tasks(ctx context.Context, obj *TaskCategory) (res []*Task, err error) {
	return r.Handlers.TaskCategoryTasks(ctx, r.GeneratedResolver, obj)
}

// TaskCategoryTasksHandler handler
func TaskCategoryTasksHandler(ctx context.Context, r *GeneratedResolver, obj *TaskCategory) (res []*Task, err error) {

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

// TasksIds ...
func (r *GeneratedTaskCategoryResolver) TasksIds(ctx context.Context, obj *TaskCategory) (ids []string, err error) {
	ids = []string{}

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("tasks")+".id").Related(&items, "Tasks").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// TasksConnection method
func (r *GeneratedTaskCategoryResolver) TasksConnection(ctx context.Context, obj *TaskCategory, offset *int, limit *int, q *string, sort []*TaskSortType, filter *TaskFilterType) (res *TaskResultType, err error) {
	f := &TaskFilterType{
		Categories: &TaskCategoryFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &TaskFilterType{
			And: []*TaskFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryTasksHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}

// QueryCompanyHandlerOptions struct
type QueryCompanyHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *CompanyFilterType
}

// Company ...
func (r *GeneratedQueryResolver) Company(ctx context.Context, id *string, q *string, filter *CompanyFilterType) (*Company, error) {
	opts := QueryCompanyHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryCompany(ctx, r.GeneratedResolver, opts)
}

// QueryCompanyHandler handler
func QueryCompanyHandler(ctx context.Context, r *GeneratedResolver, opts QueryCompanyHandlerOptions) (*Company, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := CompanyQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &CompanyResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("companies")+".id = ?", *opts.ID)
	}

	var items []*Company
	giOpts := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
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

// QueryCompaniesHandlerOptions struct
type QueryCompaniesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*CompanySortType
	Filter *CompanyFilterType
}

// Companies ...
func (r *GeneratedQueryResolver) Companies(ctx context.Context, offset *int, limit *int, q *string, sort []*CompanySortType, filter *CompanyFilterType) (*CompanyResultType, error) {
	opts := QueryCompaniesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryCompanies(ctx, r.GeneratedResolver, opts)
}

// QueryCompaniesHandler handler
func QueryCompaniesHandler(ctx context.Context, r *GeneratedResolver, opts QueryCompaniesHandlerOptions) (*CompanyResultType, error) {
	query := CompanyQueryFilter{opts.Q}

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

	return &CompanyResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedCompanyResultTypeResolver struct
type GeneratedCompanyResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedCompanyResultTypeResolver) Items(ctx context.Context, obj *CompanyResultType) (items []*Company, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*Company{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedCompanyResultTypeResolver) Count(ctx context.Context, obj *CompanyResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("companies"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &Company{})
}

// GeneratedCompanyResolver struct
type GeneratedCompanyResolver struct{ *GeneratedResolver }

// Country ...
func (r *GeneratedCompanyResolver) Country(ctx context.Context, obj *Company) (res *Country, err error) {
	return r.Handlers.CompanyCountry(ctx, r.GeneratedResolver, obj)
}

// CompanyCountryHandler handler
func CompanyCountryHandler(ctx context.Context, r *GeneratedResolver, obj *Company) (res *Country, err error) {

	if obj.CountryID != nil {
		res = &Country{ID: *obj.CountryID}
	}

	return
}

// Reviews ...
func (r *GeneratedCompanyResolver) Reviews(ctx context.Context, obj *Company) (res []*Review, err error) {
	return r.Handlers.CompanyReviews(ctx, r.GeneratedResolver, obj)
}

// CompanyReviewsHandler handler
func CompanyReviewsHandler(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*Review, err error) {

	err = fmt.Errorf("Resolver handler for CompanyReviews not implemented")

	return
}

// Employees ...
func (r *GeneratedCompanyResolver) Employees(ctx context.Context, obj *Company) (res []*User, err error) {
	return r.Handlers.CompanyEmployees(ctx, r.GeneratedResolver, obj)
}

// CompanyEmployeesHandler handler
func CompanyEmployeesHandler(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*User, err error) {

	items := []*User{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Employees").Error
	res = items

	return
}

// EmployeesIds ...
func (r *GeneratedCompanyResolver) EmployeesIds(ctx context.Context, obj *Company) (ids []string, err error) {
	ids = []string{}

	items := []*User{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("users")+".id").Related(&items, "Employees").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// EmployeesConnection method
func (r *GeneratedCompanyResolver) EmployeesConnection(ctx context.Context, obj *Company, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (res *UserResultType, err error) {
	f := &UserFilterType{
		Employers: &CompanyFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &UserFilterType{
			And: []*UserFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}

// QueryUserHandlerOptions struct
type QueryUserHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *UserFilterType
}

// User ...
func (r *GeneratedQueryResolver) User(ctx context.Context, id *string, q *string, filter *UserFilterType) (*User, error) {
	opts := QueryUserHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryUser(ctx, r.GeneratedResolver, opts)
}

// QueryUserHandler handler
func QueryUserHandler(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := UserQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &UserResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("users")+".id = ?", *opts.ID)
	}

	var items []*User
	giOpts := GetItemsOptions{
		Alias: TableName("users"),
		Preloaders: []string{
			"Employers",
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

// QueryUsersHandlerOptions struct
type QueryUsersHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*UserSortType
	Filter *UserFilterType
}

// Users ...
func (r *GeneratedQueryResolver) Users(ctx context.Context, offset *int, limit *int, q *string, sort []*UserSortType, filter *UserFilterType) (*UserResultType, error) {
	opts := QueryUsersHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryUsers(ctx, r.GeneratedResolver, opts)
}

// QueryUsersHandler handler
func QueryUsersHandler(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error) {
	query := UserQueryFilter{opts.Q}

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

	return &UserResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedUserResultTypeResolver struct
type GeneratedUserResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedUserResultTypeResolver) Items(ctx context.Context, obj *UserResultType) (items []*User, err error) {
	otps := GetItemsOptions{
		Alias: TableName("users"),
		Preloaders: []string{
			"Employers",
		},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	for _, item := range items {

		item.EmployersPreloaded = true
	}

	uniqueItems := []*User{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedUserResultTypeResolver) Count(ctx context.Context, obj *UserResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias: TableName("users"),
		Preloaders: []string{
			"Employers",
		},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &User{})
}

// GeneratedUserResolver struct
type GeneratedUserResolver struct{ *GeneratedResolver }

// Address ...
func (r *GeneratedUserResolver) Address(ctx context.Context, obj *User) (res *Address, err error) {
	return r.Handlers.UserAddress(ctx, r.GeneratedResolver, obj)
}

// UserAddressHandler handler
func UserAddressHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res *Address, err error) {

	err = fmt.Errorf("Resolver handler for UserAddress not implemented")

	return
}

// Employers ...
func (r *GeneratedUserResolver) Employers(ctx context.Context, obj *User) (res []*Company, err error) {
	return r.Handlers.UserEmployers(ctx, r.GeneratedResolver, obj)
}

// UserEmployersHandler handler
func UserEmployersHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Company, err error) {

	if obj.EmployersPreloaded {
		res = obj.Employers
	} else {

		items := []*Company{}
		db := r.GetDB(ctx)
		if db == nil {
			db = r.DB.Query()
		}
		err = db.Model(obj).Related(&items, "Employers").Error
		res = items

	}

	return
}

// EmployersIds ...
func (r *GeneratedUserResolver) EmployersIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Company{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("companies")+".id").Related(&items, "Employers").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// EmployersConnection method
func (r *GeneratedUserResolver) EmployersConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*CompanySortType, filter *CompanyFilterType) (res *CompanyResultType, err error) {
	f := &CompanyFilterType{
		Employees: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &CompanyFilterType{
			And: []*CompanyFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryCompaniesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryCompanies(ctx, r.GeneratedResolver, opts)
}

// Tasks ...
func (r *GeneratedUserResolver) Tasks(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserTasks(ctx, r.GeneratedResolver, obj)
}

// UserTasksHandler handler
func UserTasksHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error) {

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "Tasks").Error
	res = items

	return
}

// TasksIds ...
func (r *GeneratedUserResolver) TasksIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("tasks")+".id").Related(&items, "Tasks").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// TasksConnection method
func (r *GeneratedUserResolver) TasksConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*TaskSortType, filter *TaskFilterType) (res *TaskResultType, err error) {
	f := &TaskFilterType{
		Assignee: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &TaskFilterType{
			And: []*TaskFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryTasksHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}

// CreatedTasks ...
func (r *GeneratedUserResolver) CreatedTasks(ctx context.Context, obj *User) (res []*Task, err error) {
	return r.Handlers.UserCreatedTasks(ctx, r.GeneratedResolver, obj)
}

// UserCreatedTasksHandler handler
func UserCreatedTasksHandler(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error) {

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Related(&items, "CreatedTasks").Error
	res = items

	return
}

// CreatedTasksIds ...
func (r *GeneratedUserResolver) CreatedTasksIds(ctx context.Context, obj *User) (ids []string, err error) {
	ids = []string{}

	items := []*Task{}
	db := r.GetDB(ctx)
	if db == nil {
		db = r.DB.Query()
	}
	err = db.Model(obj).Select(TableName("tasks")+".id").Related(&items, "CreatedTasks").Error

	for _, item := range items {
		ids = append(ids, item.ID)
	}

	return
}

// CreatedTasksConnection method
func (r *GeneratedUserResolver) CreatedTasksConnection(ctx context.Context, obj *User, offset *int, limit *int, q *string, sort []*TaskSortType, filter *TaskFilterType) (res *TaskResultType, err error) {
	f := &TaskFilterType{
		Owner: &UserFilterType{
			ID: &obj.ID,
		},
	}
	if filter == nil {
		filter = f
	} else {
		filter = &TaskFilterType{
			And: []*TaskFilterType{
				filter,
				f,
			},
		}
	}
	opts := QueryTasksHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryTasks(ctx, r.GeneratedResolver, opts)
}

// QueryPlainEntityHandlerOptions struct
type QueryPlainEntityHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *PlainEntityFilterType
}

// PlainEntity ...
func (r *GeneratedQueryResolver) PlainEntity(ctx context.Context, id *string, q *string, filter *PlainEntityFilterType) (*PlainEntity, error) {
	opts := QueryPlainEntityHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryPlainEntity(ctx, r.GeneratedResolver, opts)
}

// QueryPlainEntityHandler handler
func QueryPlainEntityHandler(ctx context.Context, r *GeneratedResolver, opts QueryPlainEntityHandlerOptions) (*PlainEntity, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := PlainEntityQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &PlainEntityResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("plain_entities")+".id = ?", *opts.ID)
	}

	var items []*PlainEntity
	giOpts := GetItemsOptions{
		Alias:      TableName("plain_entities"),
		Preloaders: []string{},
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

// QueryPlainEntitiesHandlerOptions struct
type QueryPlainEntitiesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*PlainEntitySortType
	Filter *PlainEntityFilterType
}

// PlainEntities ...
func (r *GeneratedQueryResolver) PlainEntities(ctx context.Context, offset *int, limit *int, q *string, sort []*PlainEntitySortType, filter *PlainEntityFilterType) (*PlainEntityResultType, error) {
	opts := QueryPlainEntitiesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryPlainEntities(ctx, r.GeneratedResolver, opts)
}

// QueryPlainEntitiesHandler handler
func QueryPlainEntitiesHandler(ctx context.Context, r *GeneratedResolver, opts QueryPlainEntitiesHandlerOptions) (*PlainEntityResultType, error) {
	query := PlainEntityQueryFilter{opts.Q}

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

	return &PlainEntityResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

// GeneratedPlainEntityResultTypeResolver struct
type GeneratedPlainEntityResultTypeResolver struct{ *GeneratedResolver }

// Items ...
func (r *GeneratedPlainEntityResultTypeResolver) Items(ctx context.Context, obj *PlainEntityResultType) (items []*PlainEntity, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("plain_entities"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*PlainEntity{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

// Count ...
func (r *GeneratedPlainEntityResultTypeResolver) Count(ctx context.Context, obj *PlainEntityResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("plain_entities"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &PlainEntity{})
}

// GeneratedPlainEntityResolver struct
type GeneratedPlainEntityResolver struct{ *GeneratedResolver }

// ShortText ...
func (r *GeneratedPlainEntityResolver) ShortText(ctx context.Context, obj *PlainEntity) (res string, err error) {
	return r.Handlers.PlainEntityShortText(ctx, r.GeneratedResolver, obj)
}

// PlainEntityShortTextHandler handler
func PlainEntityShortTextHandler(ctx context.Context, r *GeneratedResolver, obj *PlainEntity) (res string, err error) {

	err = fmt.Errorf("Resolver handler for PlainEntityShortText not implemented")

	return
}
