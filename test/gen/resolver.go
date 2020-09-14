package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/loopcontext/graphql-orm/events"
)

// ResolutionHandlers struct
type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateTask     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error)
	UpdateTask     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error)
	DeleteTask     func(ctx context.Context, r *GeneratedResolver, id string) (item *Task, err error)
	DeleteAllTasks func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryTask      func(ctx context.Context, r *GeneratedResolver, opts QueryTaskHandlerOptions) (*Task, error)
	QueryTasks     func(ctx context.Context, r *GeneratedResolver, opts QueryTasksHandlerOptions) (*TaskResultType, error)

	TaskMetas func(ctx context.Context, r *GeneratedResolver, obj *Task) (res []*TaskMeta, err error)

	TaskMeta func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *TaskMeta, err error)

	TaskAssignee func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error)

	TaskOwner func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *User, err error)

	TaskParentTask func(ctx context.Context, r *GeneratedResolver, obj *Task) (res *Task, err error)

	TaskSubtasks func(ctx context.Context, r *GeneratedResolver, obj *Task) (res []*Task, err error)

	TaskCategories func(ctx context.Context, r *GeneratedResolver, obj *Task) (res []*TaskCategory, err error)

	CreateTaskCategory      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *TaskCategory, err error)
	UpdateTaskCategory      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *TaskCategory, err error)
	DeleteTaskCategory      func(ctx context.Context, r *GeneratedResolver, id string) (item *TaskCategory, err error)
	DeleteAllTaskCategories func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryTaskCategory       func(ctx context.Context, r *GeneratedResolver, opts QueryTaskCategoryHandlerOptions) (*TaskCategory, error)
	QueryTaskCategories     func(ctx context.Context, r *GeneratedResolver, opts QueryTaskCategoriesHandlerOptions) (*TaskCategoryResultType, error)

	TaskCategoryTasks func(ctx context.Context, r *GeneratedResolver, obj *TaskCategory) (res []*Task, err error)

	CreateCompany      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Company, err error)
	UpdateCompany      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Company, err error)
	DeleteCompany      func(ctx context.Context, r *GeneratedResolver, id string) (item *Company, err error)
	DeleteAllCompanies func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryCompany       func(ctx context.Context, r *GeneratedResolver, opts QueryCompanyHandlerOptions) (*Company, error)
	QueryCompanies     func(ctx context.Context, r *GeneratedResolver, opts QueryCompaniesHandlerOptions) (*CompanyResultType, error)

	CompanyCountry func(ctx context.Context, r *GeneratedResolver, obj *Company) (res *Country, err error)

	CompanyReviews func(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*Review, err error)

	CompanyEmployees func(ctx context.Context, r *GeneratedResolver, obj *Company) (res []*User, err error)

	CreateUser     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error)
	UpdateUser     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error)
	DeleteUser     func(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error)
	DeleteAllUsers func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryUser      func(ctx context.Context, r *GeneratedResolver, opts QueryUserHandlerOptions) (*User, error)
	QueryUsers     func(ctx context.Context, r *GeneratedResolver, opts QueryUsersHandlerOptions) (*UserResultType, error)

	UserAddress func(ctx context.Context, r *GeneratedResolver, obj *User) (res *Address, err error)

	UserEmployers func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Company, err error)

	UserTasks func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error)

	UserCreatedTasks func(ctx context.Context, r *GeneratedResolver, obj *User) (res []*Task, err error)

	CreatePlainEntity      func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PlainEntity, err error)
	UpdatePlainEntity      func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PlainEntity, err error)
	DeletePlainEntity      func(ctx context.Context, r *GeneratedResolver, id string) (item *PlainEntity, err error)
	DeleteAllPlainEntities func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryPlainEntity       func(ctx context.Context, r *GeneratedResolver, opts QueryPlainEntityHandlerOptions) (*PlainEntity, error)
	QueryPlainEntities     func(ctx context.Context, r *GeneratedResolver, opts QueryPlainEntitiesHandlerOptions) (*PlainEntityResultType, error)

	PlainEntityShortText func(ctx context.Context, r *GeneratedResolver, obj *PlainEntity) (res string, err error)
}

// DefaultResolutionHandlers ...
func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateTask:     CreateTaskHandler,
		UpdateTask:     UpdateTaskHandler,
		DeleteTask:     DeleteTaskHandler,
		DeleteAllTasks: DeleteAllTasksHandler,
		QueryTask:      QueryTaskHandler,
		QueryTasks:     QueryTasksHandler,

		TaskMetas: TaskMetasHandler,

		TaskMeta: TaskMetaHandler,

		TaskAssignee: TaskAssigneeHandler,

		TaskOwner: TaskOwnerHandler,

		TaskParentTask: TaskParentTaskHandler,

		TaskSubtasks: TaskSubtasksHandler,

		TaskCategories: TaskCategoriesHandler,

		CreateTaskCategory:      CreateTaskCategoryHandler,
		UpdateTaskCategory:      UpdateTaskCategoryHandler,
		DeleteTaskCategory:      DeleteTaskCategoryHandler,
		DeleteAllTaskCategories: DeleteAllTaskCategoriesHandler,
		QueryTaskCategory:       QueryTaskCategoryHandler,
		QueryTaskCategories:     QueryTaskCategoriesHandler,

		TaskCategoryTasks: TaskCategoryTasksHandler,

		CreateCompany:      CreateCompanyHandler,
		UpdateCompany:      UpdateCompanyHandler,
		DeleteCompany:      DeleteCompanyHandler,
		DeleteAllCompanies: DeleteAllCompaniesHandler,
		QueryCompany:       QueryCompanyHandler,
		QueryCompanies:     QueryCompaniesHandler,

		CompanyCountry: CompanyCountryHandler,

		CompanyReviews: CompanyReviewsHandler,

		CompanyEmployees: CompanyEmployeesHandler,

		CreateUser:     CreateUserHandler,
		UpdateUser:     UpdateUserHandler,
		DeleteUser:     DeleteUserHandler,
		DeleteAllUsers: DeleteAllUsersHandler,
		QueryUser:      QueryUserHandler,
		QueryUsers:     QueryUsersHandler,

		UserAddress: UserAddressHandler,

		UserEmployers: UserEmployersHandler,

		UserTasks: UserTasksHandler,

		UserCreatedTasks: UserCreatedTasksHandler,

		CreatePlainEntity:      CreatePlainEntityHandler,
		UpdatePlainEntity:      UpdatePlainEntityHandler,
		DeletePlainEntity:      DeletePlainEntityHandler,
		DeleteAllPlainEntities: DeleteAllPlainEntitiesHandler,
		QueryPlainEntity:       QueryPlainEntityHandler,
		QueryPlainEntities:     QueryPlainEntitiesHandler,

		PlainEntityShortText: PlainEntityShortTextHandler,
	}
	return handlers
}

// GeneratedResolver struct
type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
