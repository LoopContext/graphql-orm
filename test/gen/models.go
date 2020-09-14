package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

// TaskResultType struct
type TaskResultType struct {
	EntityResultType
}

// Task struct
type Task struct {
	ID           string     `json:"id" gorm:"column:id;primary_key"`
	Title        *string    `json:"title" gorm:"column:title"`
	Completed    *bool      `json:"completed" gorm:"column:completed;default:false"`
	State        *TaskState `json:"state" gorm:"column:state"`
	DueDate      *time.Time `json:"dueDate" gorm:"column:dueDate"`
	Metas        *string    `json:"metas" gorm:"column:metas;type:text"`
	Meta         *string    `json:"meta" gorm:"column:meta;type:text"`
	AssigneeID   *string    `json:"assigneeId" gorm:"column:assigneeId"`
	OwnerID      *string    `json:"ownerId" gorm:"column:ownerId"`
	ParentTaskID *string    `json:"parentTaskId" gorm:"column:parentTaskId"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy    *string    `json:"createdBy" gorm:"column:createdBy"`

	Assignee          *User `json:"assignee"`
	AssigneePreloaded bool  `gorm:"-"`

	Owner *User `json:"owner"`

	ParentTask *Task `json:"parentTask"`

	Subtasks []*Task `json:"subtasks" gorm:"foreignkey:ParentTaskID"`

	Categories []*TaskCategory `json:"categories" gorm:"many2many:taskCategory_tasks;jointable_foreignkey:taskId;association_jointable_foreignkey:categoryId"`
}

// IsEntity ...
func (m *Task) IsEntity() {}

// TaskChanges struct
type TaskChanges struct {
	ID           string
	Title        *string
	Completed    *bool
	State        *TaskState
	DueDate      *time.Time
	Metas        []*TaskMetaInput
	Meta         *TaskMetaInput
	AssigneeID   *string
	OwnerID      *string
	ParentTaskID *string
	UpdatedAt    *time.Time
	CreatedAt    time.Time
	UpdatedBy    *string
	CreatedBy    *string

	SubtasksIDs   []*string
	CategoriesIDs []*string
}

// TaskCategoryResultType struct
type TaskCategoryResultType struct {
	EntityResultType
}

// TaskCategory struct
type TaskCategory struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Name      *string    `json:"name" gorm:"column:name"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Tasks []*Task `json:"tasks" gorm:"many2many:taskCategory_tasks;jointable_foreignkey:categoryId;association_jointable_foreignkey:taskId"`
}

// IsEntity ...
func (m *TaskCategory) IsEntity() {}

// TaskCategoryChanges struct
type TaskCategoryChanges struct {
	ID        string
	Name      *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	TasksIDs []*string
}

// TaskCategoryTasks struct
type TaskCategoryTasks struct {
	CategoryID string
	TaskID     string
}

// TableName ...
func (TaskCategoryTasks) TableName() string {
	return TableName("taskCategory_tasks")
}

// CompanyResultType struct
type CompanyResultType struct {
	EntityResultType
}

// Company struct
type Company struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Name      *string    `json:"name" gorm:"column:name"`
	CountryID *string    `json:"countryId" gorm:"column:countryId"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`

	Employees []*User `json:"employees" gorm:"many2many:user_employers;jointable_foreignkey:employerId;association_jointable_foreignkey:employeeId"`
}

// IsEntity ...
func (m *Company) IsEntity() {}

// IsNamedEntity ...
func (m *Company) IsNamedEntity() {}

// CompanyChanges struct
type CompanyChanges struct {
	ID        string
	Name      *string
	CountryID *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string

	EmployeesIDs []*string
}

// UserResultType struct
type UserResultType struct {
	EntityResultType
}

// User struct
type User struct {
	ID         string     `json:"id" gorm:"column:id;primary_key"`
	Code       *int       `json:"code" gorm:"column:code"`
	Email      *string    `json:"email" gorm:"column:email;unique"`
	FirstName  *string    `json:"firstName" gorm:"column:firstName"`
	LastName   *string    `json:"lastName" gorm:"column:lastName"`
	AddressRaw *string    `json:"addressRaw" gorm:"column:addressRaw"`
	Salary     *int       `json:"salary" gorm:"column:salary"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy  *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy  *string    `json:"createdBy" gorm:"column:createdBy"`

	Employers          []*Company `json:"employers" gorm:"many2many:user_employers;jointable_foreignkey:employeeId;association_jointable_foreignkey:employerId"`
	EmployersPreloaded bool       `gorm:"-"`

	Tasks []*Task `json:"tasks" gorm:"foreignkey:AssigneeID"`

	CreatedTasks []*Task `json:"createdTasks" gorm:"foreignkey:OwnerID"`
}

// IsEntity ...
func (m *User) IsEntity() {}

// UserChanges struct
type UserChanges struct {
	ID         string
	Code       *int
	Email      *string
	FirstName  *string
	LastName   *string
	AddressRaw *string
	Salary     *int
	UpdatedAt  *time.Time
	CreatedAt  time.Time
	UpdatedBy  *string
	CreatedBy  *string

	EmployersIDs    []*string
	TasksIDs        []*string
	CreatedTasksIDs []*string
}

// UserEmployers struct
type UserEmployers struct {
	EmployeeID string
	EmployerID string
}

// TableName ...
func (UserEmployers) TableName() string {
	return TableName("user_employers")
}

// PlainEntityResultType struct
type PlainEntityResultType struct {
	EntityResultType
}

// PlainEntity struct
type PlainEntity struct {
	ID        string     `json:"id" gorm:"column:id;primary_key"`
	Date      *time.Time `json:"date" gorm:"column:date"`
	Text      *string    `json:"text" gorm:"column:text;type:text"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy *string    `json:"createdBy" gorm:"column:createdBy"`
}

// IsEntity ...
func (m *PlainEntity) IsEntity() {}

// PlainEntityChanges struct
type PlainEntityChanges struct {
	ID        string
	Date      *time.Time
	Text      *string
	UpdatedAt *time.Time
	CreatedAt time.Time
	UpdatedBy *string
	CreatedBy *string
}

// ApplyChanges used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
