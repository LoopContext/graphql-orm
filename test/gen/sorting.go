package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Apply method
func (s TaskSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("tasks"), sorts, joins)
}

// ApplyWithAlias method
func (s TaskSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Title != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("title"), Direction: s.Title.String()}
		*sorts = append(*sorts, sort)
	}

	if s.TitleMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("title") + ")", Direction: s.TitleMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.TitleMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("title") + ")", Direction: s.TitleMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Completed != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("completed"), Direction: s.Completed.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CompletedMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("completed") + ")", Direction: s.CompletedMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CompletedMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("completed") + ")", Direction: s.CompletedMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DueDate != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("dueDate"), Direction: s.DueDate.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DueDateMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("dueDate") + ")", Direction: s.DueDateMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DueDateMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("dueDate") + ")", Direction: s.DueDateMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AssigneeID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("assigneeId"), Direction: s.AssigneeID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AssigneeIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("assigneeId") + ")", Direction: s.AssigneeIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AssigneeIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("assigneeId") + ")", Direction: s.AssigneeIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.OwnerID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("ownerId"), Direction: s.OwnerID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.OwnerIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("ownerId") + ")", Direction: s.OwnerIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.OwnerIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("ownerId") + ")", Direction: s.OwnerIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ParentTaskID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("parentTaskId"), Direction: s.ParentTaskID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ParentTaskIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("parentTaskId") + ")", Direction: s.ParentTaskIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ParentTaskIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("parentTaskId") + ")", Direction: s.ParentTaskIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Assignee != nil {
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("assigneeId"))
		err := s.Assignee.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Owner != nil {
		_alias := alias + "_owner"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("ownerId"))
		err := s.Owner.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.ParentTask != nil {
		_alias := alias + "_parentTask"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("parentTaskId"))
		err := s.ParentTask.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Subtasks != nil {
		_alias := alias + "_subtasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("parentTaskId")+" = "+dialect.Quote(alias)+".id")
		err := s.Subtasks.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Categories != nil {
		_alias := alias + "_categories"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("taskCategory_tasks"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("taskId")+" LEFT JOIN "+dialect.Quote(TableName("task_categories"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("categoryId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Categories.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s TaskCategorySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("task_categories"), sorts, joins)
}

// ApplyWithAlias method
func (s TaskCategorySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("taskCategory_tasks"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("categoryId")+" LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("taskId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Tasks.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s CompanySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("companies"), sorts, joins)
}

// ApplyWithAlias method
func (s CompanySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CountryID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("countryId"), Direction: s.CountryID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CountryIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("countryId") + ")", Direction: s.CountryIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CountryIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("countryId") + ")", Direction: s.CountryIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Employees != nil {
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_employers"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employerId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employeeId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Employees.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s UserSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("users"), sorts, joins)
}

// ApplyWithAlias method
func (s UserSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Code != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("code"), Direction: s.Code.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CodeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CodeAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("code") + ")", Direction: s.CodeAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Email != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("email"), Direction: s.Email.String()}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("firstName"), Direction: s.FirstName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.FirstNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("firstName") + ")", Direction: s.FirstNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastName != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("lastName"), Direction: s.LastName.String()}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.LastNameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("lastName") + ")", Direction: s.LastNameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AddressRaw != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("addressRaw"), Direction: s.AddressRaw.String()}
		*sorts = append(*sorts, sort)
	}

	if s.AddressRawMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("addressRaw") + ")", Direction: s.AddressRawMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.AddressRawMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("addressRaw") + ")", Direction: s.AddressRawMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Salary != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("salary"), Direction: s.Salary.String()}
		*sorts = append(*sorts, sort)
	}

	if s.SalaryMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("salary") + ")", Direction: s.SalaryMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SalaryMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("salary") + ")", Direction: s.SalaryMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SalaryAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("salary") + ")", Direction: s.SalaryAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Employers != nil {
		_alias := alias + "_employers"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_employers"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employeeId")+" LEFT JOIN "+dialect.Quote(TableName("companies"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employerId")+" = "+dialect.Quote(_alias)+".id")
		err := s.Employers.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("assigneeId")+" = "+dialect.Quote(alias)+".id")
		err := s.Tasks.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	if s.CreatedTasks != nil {
		_alias := alias + "_createdTasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("ownerId")+" = "+dialect.Quote(alias)+".id")
		err := s.CreatedTasks.ApplyWithAlias(ctx, dialect, _alias, sorts, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Apply method
func (s PlainEntitySortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("plain_entities"), sorts, joins)
}

// ApplyWithAlias method
func (s PlainEntitySortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Date != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("date"), Direction: s.Date.String()}
		*sorts = append(*sorts, sort)
	}

	if s.DateMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("date") + ")", Direction: s.DateMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.DateMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("date") + ")", Direction: s.DateMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Text != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("text"), Direction: s.Text.String()}
		*sorts = append(*sorts, sort)
	}

	if s.TextMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("text") + ")", Direction: s.TextMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.TextMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("text") + ")", Direction: s.TextMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}
