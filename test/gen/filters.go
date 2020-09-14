package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// IsEmpty ...
func (f *TaskFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *TaskFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("tasks"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *TaskFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Assignee != nil {
		_alias := alias + "_assignee"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("assigneeId"))
		err := f.Assignee.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Owner != nil {
		_alias := alias + "_owner"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("ownerId"))
		err := f.Owner.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.ParentTask != nil {
		_alias := alias + "_parentTask"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+".id = "+alias+"."+dialect.Quote("parentTaskId"))
		err := f.ParentTask.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Subtasks != nil {
		_alias := alias + "_subtasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("parentTaskId")+" = "+dialect.Quote(alias)+".id")
		err := f.Subtasks.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Categories != nil {
		_alias := alias + "_categories"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("taskCategory_tasks"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("taskId")+" LEFT JOIN "+dialect.Quote(TableName("task_categories"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("categoryId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Categories.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *TaskFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Title != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" = ?")
		values = append(values, f.Title)
	}

	if f.TitleNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" != ?")
		values = append(values, f.TitleNe)
	}

	if f.TitleGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" > ?")
		values = append(values, f.TitleGt)
	}

	if f.TitleLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" < ?")
		values = append(values, f.TitleLt)
	}

	if f.TitleGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" >= ?")
		values = append(values, f.TitleGte)
	}

	if f.TitleLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" <= ?")
		values = append(values, f.TitleLte)
	}

	if f.TitleIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" IN (?)")
		values = append(values, f.TitleIn)
	}

	if f.TitleLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TitleLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TitlePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitlePrefix))
	}

	if f.TitleSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleSuffix))
	}

	if f.TitleNull != nil {
		if *f.TitleNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("title")+" IS NOT NULL")
		}
	}

	if f.Completed != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" = ?")
		values = append(values, f.Completed)
	}

	if f.CompletedNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" != ?")
		values = append(values, f.CompletedNe)
	}

	if f.CompletedGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" > ?")
		values = append(values, f.CompletedGt)
	}

	if f.CompletedLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" < ?")
		values = append(values, f.CompletedLt)
	}

	if f.CompletedGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" >= ?")
		values = append(values, f.CompletedGte)
	}

	if f.CompletedLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" <= ?")
		values = append(values, f.CompletedLte)
	}

	if f.CompletedIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IN (?)")
		values = append(values, f.CompletedIn)
	}

	if f.CompletedNull != nil {
		if *f.CompletedNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("completed")+" IS NOT NULL")
		}
	}

	if f.State != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" = ?")
		values = append(values, f.State)
	}

	if f.StateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" != ?")
		values = append(values, f.StateNe)
	}

	if f.StateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" > ?")
		values = append(values, f.StateGt)
	}

	if f.StateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" < ?")
		values = append(values, f.StateLt)
	}

	if f.StateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" >= ?")
		values = append(values, f.StateGte)
	}

	if f.StateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" <= ?")
		values = append(values, f.StateLte)
	}

	if f.StateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IN (?)")
		values = append(values, f.StateIn)
	}

	if f.StateNull != nil {
		if *f.StateNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("state")+" IS NOT NULL")
		}
	}

	if f.DueDate != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" = ?")
		values = append(values, f.DueDate)
	}

	if f.DueDateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" != ?")
		values = append(values, f.DueDateNe)
	}

	if f.DueDateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" > ?")
		values = append(values, f.DueDateGt)
	}

	if f.DueDateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" < ?")
		values = append(values, f.DueDateLt)
	}

	if f.DueDateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" >= ?")
		values = append(values, f.DueDateGte)
	}

	if f.DueDateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" <= ?")
		values = append(values, f.DueDateLte)
	}

	if f.DueDateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" IN (?)")
		values = append(values, f.DueDateIn)
	}

	if f.DueDateNull != nil {
		if *f.DueDateNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("dueDate")+" IS NOT NULL")
		}
	}

	if f.AssigneeID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" = ?")
		values = append(values, f.AssigneeID)
	}

	if f.AssigneeIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" != ?")
		values = append(values, f.AssigneeIDNe)
	}

	if f.AssigneeIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" > ?")
		values = append(values, f.AssigneeIDGt)
	}

	if f.AssigneeIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" < ?")
		values = append(values, f.AssigneeIDLt)
	}

	if f.AssigneeIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" >= ?")
		values = append(values, f.AssigneeIDGte)
	}

	if f.AssigneeIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" <= ?")
		values = append(values, f.AssigneeIDLte)
	}

	if f.AssigneeIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" IN (?)")
		values = append(values, f.AssigneeIDIn)
	}

	if f.AssigneeIDNull != nil {
		if *f.AssigneeIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("assigneeId")+" IS NOT NULL")
		}
	}

	if f.OwnerID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" = ?")
		values = append(values, f.OwnerID)
	}

	if f.OwnerIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" != ?")
		values = append(values, f.OwnerIDNe)
	}

	if f.OwnerIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" > ?")
		values = append(values, f.OwnerIDGt)
	}

	if f.OwnerIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" < ?")
		values = append(values, f.OwnerIDLt)
	}

	if f.OwnerIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" >= ?")
		values = append(values, f.OwnerIDGte)
	}

	if f.OwnerIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" <= ?")
		values = append(values, f.OwnerIDLte)
	}

	if f.OwnerIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" IN (?)")
		values = append(values, f.OwnerIDIn)
	}

	if f.OwnerIDNull != nil {
		if *f.OwnerIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("ownerId")+" IS NOT NULL")
		}
	}

	if f.ParentTaskID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" = ?")
		values = append(values, f.ParentTaskID)
	}

	if f.ParentTaskIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" != ?")
		values = append(values, f.ParentTaskIDNe)
	}

	if f.ParentTaskIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" > ?")
		values = append(values, f.ParentTaskIDGt)
	}

	if f.ParentTaskIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" < ?")
		values = append(values, f.ParentTaskIDLt)
	}

	if f.ParentTaskIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" >= ?")
		values = append(values, f.ParentTaskIDGte)
	}

	if f.ParentTaskIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" <= ?")
		values = append(values, f.ParentTaskIDLte)
	}

	if f.ParentTaskIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" IN (?)")
		values = append(values, f.ParentTaskIDIn)
	}

	if f.ParentTaskIDNull != nil {
		if *f.ParentTaskIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("parentTaskId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *TaskFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.TitleMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") = ?")
		values = append(values, f.TitleMin)
	}

	if f.TitleMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") = ?")
		values = append(values, f.TitleMax)
	}

	if f.TitleMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") != ?")
		values = append(values, f.TitleMinNe)
	}

	if f.TitleMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") != ?")
		values = append(values, f.TitleMaxNe)
	}

	if f.TitleMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") > ?")
		values = append(values, f.TitleMinGt)
	}

	if f.TitleMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") > ?")
		values = append(values, f.TitleMaxGt)
	}

	if f.TitleMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") < ?")
		values = append(values, f.TitleMinLt)
	}

	if f.TitleMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") < ?")
		values = append(values, f.TitleMaxLt)
	}

	if f.TitleMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") >= ?")
		values = append(values, f.TitleMinGte)
	}

	if f.TitleMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") >= ?")
		values = append(values, f.TitleMaxGte)
	}

	if f.TitleMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") <= ?")
		values = append(values, f.TitleMinLte)
	}

	if f.TitleMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") <= ?")
		values = append(values, f.TitleMaxLte)
	}

	if f.TitleMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") IN (?)")
		values = append(values, f.TitleMinIn)
	}

	if f.TitleMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") IN (?)")
		values = append(values, f.TitleMaxIn)
	}

	if f.TitleMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TitleMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TitleMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TitleMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TitleMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitleMinPrefix))
	}

	if f.TitleMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TitleMaxPrefix))
	}

	if f.TitleMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("title")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleMinSuffix))
	}

	if f.TitleMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("title")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TitleMaxSuffix))
	}

	if f.CompletedMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") = ?")
		values = append(values, f.CompletedMin)
	}

	if f.CompletedMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") = ?")
		values = append(values, f.CompletedMax)
	}

	if f.CompletedMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") != ?")
		values = append(values, f.CompletedMinNe)
	}

	if f.CompletedMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") != ?")
		values = append(values, f.CompletedMaxNe)
	}

	if f.CompletedMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") > ?")
		values = append(values, f.CompletedMinGt)
	}

	if f.CompletedMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") > ?")
		values = append(values, f.CompletedMaxGt)
	}

	if f.CompletedMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") < ?")
		values = append(values, f.CompletedMinLt)
	}

	if f.CompletedMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") < ?")
		values = append(values, f.CompletedMaxLt)
	}

	if f.CompletedMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") >= ?")
		values = append(values, f.CompletedMinGte)
	}

	if f.CompletedMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") >= ?")
		values = append(values, f.CompletedMaxGte)
	}

	if f.CompletedMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") <= ?")
		values = append(values, f.CompletedMinLte)
	}

	if f.CompletedMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") <= ?")
		values = append(values, f.CompletedMaxLte)
	}

	if f.CompletedMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("completed")+") IN (?)")
		values = append(values, f.CompletedMinIn)
	}

	if f.CompletedMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("completed")+") IN (?)")
		values = append(values, f.CompletedMaxIn)
	}

	if f.StateMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") = ?")
		values = append(values, f.StateMin)
	}

	if f.StateMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") = ?")
		values = append(values, f.StateMax)
	}

	if f.StateMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") != ?")
		values = append(values, f.StateMinNe)
	}

	if f.StateMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") != ?")
		values = append(values, f.StateMaxNe)
	}

	if f.StateMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") > ?")
		values = append(values, f.StateMinGt)
	}

	if f.StateMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") > ?")
		values = append(values, f.StateMaxGt)
	}

	if f.StateMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") < ?")
		values = append(values, f.StateMinLt)
	}

	if f.StateMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") < ?")
		values = append(values, f.StateMaxLt)
	}

	if f.StateMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") >= ?")
		values = append(values, f.StateMinGte)
	}

	if f.StateMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") >= ?")
		values = append(values, f.StateMaxGte)
	}

	if f.StateMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") <= ?")
		values = append(values, f.StateMinLte)
	}

	if f.StateMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") <= ?")
		values = append(values, f.StateMaxLte)
	}

	if f.StateMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("state")+") IN (?)")
		values = append(values, f.StateMinIn)
	}

	if f.StateMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("state")+") IN (?)")
		values = append(values, f.StateMaxIn)
	}

	if f.DueDateMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dueDate")+") = ?")
		values = append(values, f.DueDateMin)
	}

	if f.DueDateMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dueDate")+") = ?")
		values = append(values, f.DueDateMax)
	}

	if f.DueDateMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dueDate")+") != ?")
		values = append(values, f.DueDateMinNe)
	}

	if f.DueDateMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dueDate")+") != ?")
		values = append(values, f.DueDateMaxNe)
	}

	if f.DueDateMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dueDate")+") > ?")
		values = append(values, f.DueDateMinGt)
	}

	if f.DueDateMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dueDate")+") > ?")
		values = append(values, f.DueDateMaxGt)
	}

	if f.DueDateMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dueDate")+") < ?")
		values = append(values, f.DueDateMinLt)
	}

	if f.DueDateMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dueDate")+") < ?")
		values = append(values, f.DueDateMaxLt)
	}

	if f.DueDateMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dueDate")+") >= ?")
		values = append(values, f.DueDateMinGte)
	}

	if f.DueDateMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dueDate")+") >= ?")
		values = append(values, f.DueDateMaxGte)
	}

	if f.DueDateMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dueDate")+") <= ?")
		values = append(values, f.DueDateMinLte)
	}

	if f.DueDateMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dueDate")+") <= ?")
		values = append(values, f.DueDateMaxLte)
	}

	if f.DueDateMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("dueDate")+") IN (?)")
		values = append(values, f.DueDateMinIn)
	}

	if f.DueDateMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("dueDate")+") IN (?)")
		values = append(values, f.DueDateMaxIn)
	}

	if f.AssigneeIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("assigneeId")+") = ?")
		values = append(values, f.AssigneeIDMin)
	}

	if f.AssigneeIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("assigneeId")+") = ?")
		values = append(values, f.AssigneeIDMax)
	}

	if f.AssigneeIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("assigneeId")+") != ?")
		values = append(values, f.AssigneeIDMinNe)
	}

	if f.AssigneeIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("assigneeId")+") != ?")
		values = append(values, f.AssigneeIDMaxNe)
	}

	if f.AssigneeIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("assigneeId")+") > ?")
		values = append(values, f.AssigneeIDMinGt)
	}

	if f.AssigneeIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("assigneeId")+") > ?")
		values = append(values, f.AssigneeIDMaxGt)
	}

	if f.AssigneeIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("assigneeId")+") < ?")
		values = append(values, f.AssigneeIDMinLt)
	}

	if f.AssigneeIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("assigneeId")+") < ?")
		values = append(values, f.AssigneeIDMaxLt)
	}

	if f.AssigneeIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("assigneeId")+") >= ?")
		values = append(values, f.AssigneeIDMinGte)
	}

	if f.AssigneeIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("assigneeId")+") >= ?")
		values = append(values, f.AssigneeIDMaxGte)
	}

	if f.AssigneeIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("assigneeId")+") <= ?")
		values = append(values, f.AssigneeIDMinLte)
	}

	if f.AssigneeIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("assigneeId")+") <= ?")
		values = append(values, f.AssigneeIDMaxLte)
	}

	if f.AssigneeIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("assigneeId")+") IN (?)")
		values = append(values, f.AssigneeIDMinIn)
	}

	if f.AssigneeIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("assigneeId")+") IN (?)")
		values = append(values, f.AssigneeIDMaxIn)
	}

	if f.OwnerIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("ownerId")+") = ?")
		values = append(values, f.OwnerIDMin)
	}

	if f.OwnerIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("ownerId")+") = ?")
		values = append(values, f.OwnerIDMax)
	}

	if f.OwnerIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("ownerId")+") != ?")
		values = append(values, f.OwnerIDMinNe)
	}

	if f.OwnerIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("ownerId")+") != ?")
		values = append(values, f.OwnerIDMaxNe)
	}

	if f.OwnerIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("ownerId")+") > ?")
		values = append(values, f.OwnerIDMinGt)
	}

	if f.OwnerIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("ownerId")+") > ?")
		values = append(values, f.OwnerIDMaxGt)
	}

	if f.OwnerIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("ownerId")+") < ?")
		values = append(values, f.OwnerIDMinLt)
	}

	if f.OwnerIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("ownerId")+") < ?")
		values = append(values, f.OwnerIDMaxLt)
	}

	if f.OwnerIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("ownerId")+") >= ?")
		values = append(values, f.OwnerIDMinGte)
	}

	if f.OwnerIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("ownerId")+") >= ?")
		values = append(values, f.OwnerIDMaxGte)
	}

	if f.OwnerIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("ownerId")+") <= ?")
		values = append(values, f.OwnerIDMinLte)
	}

	if f.OwnerIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("ownerId")+") <= ?")
		values = append(values, f.OwnerIDMaxLte)
	}

	if f.OwnerIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("ownerId")+") IN (?)")
		values = append(values, f.OwnerIDMinIn)
	}

	if f.OwnerIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("ownerId")+") IN (?)")
		values = append(values, f.OwnerIDMaxIn)
	}

	if f.ParentTaskIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentTaskId")+") = ?")
		values = append(values, f.ParentTaskIDMin)
	}

	if f.ParentTaskIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentTaskId")+") = ?")
		values = append(values, f.ParentTaskIDMax)
	}

	if f.ParentTaskIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentTaskId")+") != ?")
		values = append(values, f.ParentTaskIDMinNe)
	}

	if f.ParentTaskIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentTaskId")+") != ?")
		values = append(values, f.ParentTaskIDMaxNe)
	}

	if f.ParentTaskIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentTaskId")+") > ?")
		values = append(values, f.ParentTaskIDMinGt)
	}

	if f.ParentTaskIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentTaskId")+") > ?")
		values = append(values, f.ParentTaskIDMaxGt)
	}

	if f.ParentTaskIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentTaskId")+") < ?")
		values = append(values, f.ParentTaskIDMinLt)
	}

	if f.ParentTaskIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentTaskId")+") < ?")
		values = append(values, f.ParentTaskIDMaxLt)
	}

	if f.ParentTaskIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentTaskId")+") >= ?")
		values = append(values, f.ParentTaskIDMinGte)
	}

	if f.ParentTaskIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentTaskId")+") >= ?")
		values = append(values, f.ParentTaskIDMaxGte)
	}

	if f.ParentTaskIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentTaskId")+") <= ?")
		values = append(values, f.ParentTaskIDMinLte)
	}

	if f.ParentTaskIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentTaskId")+") <= ?")
		values = append(values, f.ParentTaskIDMaxLte)
	}

	if f.ParentTaskIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("parentTaskId")+") IN (?)")
		values = append(values, f.ParentTaskIDMinIn)
	}

	if f.ParentTaskIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("parentTaskId")+") IN (?)")
		values = append(values, f.ParentTaskIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *TaskFilterType) AndWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TaskFilterType) OrWith(f2 ...*TaskFilterType) *TaskFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *TaskCategoryFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *TaskCategoryFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("task_categories"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *TaskCategoryFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("taskCategory_tasks"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("categoryId")+" LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("taskId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Tasks.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *TaskCategoryFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *TaskCategoryFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *TaskCategoryFilterType) AndWith(f2 ...*TaskCategoryFilterType) *TaskCategoryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskCategoryFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *TaskCategoryFilterType) OrWith(f2 ...*TaskCategoryFilterType) *TaskCategoryFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &TaskCategoryFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *CompanyFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *CompanyFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("companies"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *CompanyFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Employees != nil {
		_alias := alias + "_employees"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_employers"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employerId")+" LEFT JOIN "+dialect.Quote(TableName("users"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employeeId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Employees.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *CompanyFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.CountryID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" = ?")
		values = append(values, f.CountryID)
	}

	if f.CountryIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" != ?")
		values = append(values, f.CountryIDNe)
	}

	if f.CountryIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" > ?")
		values = append(values, f.CountryIDGt)
	}

	if f.CountryIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" < ?")
		values = append(values, f.CountryIDLt)
	}

	if f.CountryIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" >= ?")
		values = append(values, f.CountryIDGte)
	}

	if f.CountryIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" <= ?")
		values = append(values, f.CountryIDLte)
	}

	if f.CountryIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" IN (?)")
		values = append(values, f.CountryIDIn)
	}

	if f.CountryIDNull != nil {
		if *f.CountryIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("countryId")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *CompanyFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.CountryIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("countryId")+") = ?")
		values = append(values, f.CountryIDMin)
	}

	if f.CountryIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("countryId")+") = ?")
		values = append(values, f.CountryIDMax)
	}

	if f.CountryIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("countryId")+") != ?")
		values = append(values, f.CountryIDMinNe)
	}

	if f.CountryIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("countryId")+") != ?")
		values = append(values, f.CountryIDMaxNe)
	}

	if f.CountryIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("countryId")+") > ?")
		values = append(values, f.CountryIDMinGt)
	}

	if f.CountryIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("countryId")+") > ?")
		values = append(values, f.CountryIDMaxGt)
	}

	if f.CountryIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("countryId")+") < ?")
		values = append(values, f.CountryIDMinLt)
	}

	if f.CountryIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("countryId")+") < ?")
		values = append(values, f.CountryIDMaxLt)
	}

	if f.CountryIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("countryId")+") >= ?")
		values = append(values, f.CountryIDMinGte)
	}

	if f.CountryIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("countryId")+") >= ?")
		values = append(values, f.CountryIDMaxGte)
	}

	if f.CountryIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("countryId")+") <= ?")
		values = append(values, f.CountryIDMinLte)
	}

	if f.CountryIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("countryId")+") <= ?")
		values = append(values, f.CountryIDMaxLte)
	}

	if f.CountryIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("countryId")+") IN (?)")
		values = append(values, f.CountryIDMinIn)
	}

	if f.CountryIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("countryId")+") IN (?)")
		values = append(values, f.CountryIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *CompanyFilterType) AndWith(f2 ...*CompanyFilterType) *CompanyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CompanyFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *CompanyFilterType) OrWith(f2 ...*CompanyFilterType) *CompanyFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &CompanyFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *UserFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *UserFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("users"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *UserFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	if f.Employers != nil {
		_alias := alias + "_employers"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("user_employers"))+" "+dialect.Quote(_alias+"_jointable")+" ON "+dialect.Quote(alias)+".id = "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employeeId")+" LEFT JOIN "+dialect.Quote(TableName("companies"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias+"_jointable")+"."+dialect.Quote("employerId")+" = "+dialect.Quote(_alias)+".id")
		err := f.Employers.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.Tasks != nil {
		_alias := alias + "_tasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("assigneeId")+" = "+dialect.Quote(alias)+".id")
		err := f.Tasks.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	if f.CreatedTasks != nil {
		_alias := alias + "_createdTasks"
		*joins = append(*joins, "LEFT JOIN "+dialect.Quote(TableName("tasks"))+" "+dialect.Quote(_alias)+" ON "+dialect.Quote(_alias)+"."+dialect.Quote("ownerId")+" = "+dialect.Quote(alias)+".id")
		err := f.CreatedTasks.ApplyWithAlias(ctx, dialect, _alias, wheres, whereValues, havings, havingValues, joins)
		if err != nil {
			return err
		}
	}

	return nil
}

// WhereContent ...
func (f *UserFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Code != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" = ?")
		values = append(values, f.Code)
	}

	if f.CodeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" != ?")
		values = append(values, f.CodeNe)
	}

	if f.CodeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" > ?")
		values = append(values, f.CodeGt)
	}

	if f.CodeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" < ?")
		values = append(values, f.CodeLt)
	}

	if f.CodeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" >= ?")
		values = append(values, f.CodeGte)
	}

	if f.CodeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" <= ?")
		values = append(values, f.CodeLte)
	}

	if f.CodeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IN (?)")
		values = append(values, f.CodeIn)
	}

	if f.CodeNull != nil {
		if *f.CodeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("code")+" IS NOT NULL")
		}
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" = ?")
		values = append(values, f.Email)
	}

	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" != ?")
		values = append(values, f.EmailNe)
	}

	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" > ?")
		values = append(values, f.EmailGt)
	}

	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" < ?")
		values = append(values, f.EmailLt)
	}

	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" >= ?")
		values = append(values, f.EmailGte)
	}

	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" <= ?")
		values = append(values, f.EmailLte)
	}

	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}

	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}

	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.EmailNull != nil {
		if *f.EmailNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NOT NULL")
		}
	}

	if f.FirstName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" = ?")
		values = append(values, f.FirstName)
	}

	if f.FirstNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" != ?")
		values = append(values, f.FirstNameNe)
	}

	if f.FirstNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" > ?")
		values = append(values, f.FirstNameGt)
	}

	if f.FirstNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" < ?")
		values = append(values, f.FirstNameLt)
	}

	if f.FirstNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" >= ?")
		values = append(values, f.FirstNameGte)
	}

	if f.FirstNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" <= ?")
		values = append(values, f.FirstNameLte)
	}

	if f.FirstNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IN (?)")
		values = append(values, f.FirstNameIn)
	}

	if f.FirstNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNamePrefix))
	}

	if f.FirstNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameSuffix))
	}

	if f.FirstNameNull != nil {
		if *f.FirstNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("firstName")+" IS NOT NULL")
		}
	}

	if f.LastName != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" = ?")
		values = append(values, f.LastName)
	}

	if f.LastNameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" != ?")
		values = append(values, f.LastNameNe)
	}

	if f.LastNameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" > ?")
		values = append(values, f.LastNameGt)
	}

	if f.LastNameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" < ?")
		values = append(values, f.LastNameLt)
	}

	if f.LastNameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" >= ?")
		values = append(values, f.LastNameGte)
	}

	if f.LastNameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" <= ?")
		values = append(values, f.LastNameLte)
	}

	if f.LastNameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IN (?)")
		values = append(values, f.LastNameIn)
	}

	if f.LastNameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNamePrefix))
	}

	if f.LastNameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameSuffix))
	}

	if f.LastNameNull != nil {
		if *f.LastNameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("lastName")+" IS NOT NULL")
		}
	}

	if f.AddressRaw != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" = ?")
		values = append(values, f.AddressRaw)
	}

	if f.AddressRawNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" != ?")
		values = append(values, f.AddressRawNe)
	}

	if f.AddressRawGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" > ?")
		values = append(values, f.AddressRawGt)
	}

	if f.AddressRawLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" < ?")
		values = append(values, f.AddressRawLt)
	}

	if f.AddressRawGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" >= ?")
		values = append(values, f.AddressRawGte)
	}

	if f.AddressRawLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" <= ?")
		values = append(values, f.AddressRawLte)
	}

	if f.AddressRawIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" IN (?)")
		values = append(values, f.AddressRawIn)
	}

	if f.AddressRawLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AddressRawLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AddressRawPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AddressRawPrefix))
	}

	if f.AddressRawSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AddressRawSuffix))
	}

	if f.AddressRawNull != nil {
		if *f.AddressRawNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("addressRaw")+" IS NOT NULL")
		}
	}

	if f.Salary != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" = ?")
		values = append(values, f.Salary)
	}

	if f.SalaryNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" != ?")
		values = append(values, f.SalaryNe)
	}

	if f.SalaryGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" > ?")
		values = append(values, f.SalaryGt)
	}

	if f.SalaryLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" < ?")
		values = append(values, f.SalaryLt)
	}

	if f.SalaryGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" >= ?")
		values = append(values, f.SalaryGte)
	}

	if f.SalaryLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" <= ?")
		values = append(values, f.SalaryLte)
	}

	if f.SalaryIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" IN (?)")
		values = append(values, f.SalaryIn)
	}

	if f.SalaryNull != nil {
		if *f.SalaryNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("salary")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *UserFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.CodeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMin)
	}

	if f.CodeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeMax)
	}

	if f.CodeAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("code")+") = ?")
		values = append(values, f.CodeAvg)
	}

	if f.CodeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMinNe)
	}

	if f.CodeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeMaxNe)
	}

	if f.CodeAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("code")+") != ?")
		values = append(values, f.CodeAvgNe)
	}

	if f.CodeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMinGt)
	}

	if f.CodeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeMaxGt)
	}

	if f.CodeAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("code")+") > ?")
		values = append(values, f.CodeAvgGt)
	}

	if f.CodeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMinLt)
	}

	if f.CodeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeMaxLt)
	}

	if f.CodeAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("code")+") < ?")
		values = append(values, f.CodeAvgLt)
	}

	if f.CodeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMinGte)
	}

	if f.CodeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeMaxGte)
	}

	if f.CodeAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("code")+") >= ?")
		values = append(values, f.CodeAvgGte)
	}

	if f.CodeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMinLte)
	}

	if f.CodeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeMaxLte)
	}

	if f.CodeAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("code")+") <= ?")
		values = append(values, f.CodeAvgLte)
	}

	if f.CodeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMinIn)
	}

	if f.CodeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeMaxIn)
	}

	if f.CodeAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("code")+") IN (?)")
		values = append(values, f.CodeAvgIn)
	}

	if f.EmailMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMin)
	}

	if f.EmailMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMax)
	}

	if f.EmailMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMinNe)
	}

	if f.EmailMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMaxNe)
	}

	if f.EmailMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMinGt)
	}

	if f.EmailMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMaxGt)
	}

	if f.EmailMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMinLt)
	}

	if f.EmailMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMaxLt)
	}

	if f.EmailMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMinGte)
	}

	if f.EmailMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMaxGte)
	}

	if f.EmailMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMinLte)
	}

	if f.EmailMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMaxLte)
	}

	if f.EmailMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMinIn)
	}

	if f.EmailMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMaxIn)
	}

	if f.EmailMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMinPrefix))
	}

	if f.EmailMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMaxPrefix))
	}

	if f.EmailMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMinSuffix))
	}

	if f.EmailMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMaxSuffix))
	}

	if f.FirstNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMin)
	}

	if f.FirstNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") = ?")
		values = append(values, f.FirstNameMax)
	}

	if f.FirstNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMinNe)
	}

	if f.FirstNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") != ?")
		values = append(values, f.FirstNameMaxNe)
	}

	if f.FirstNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMinGt)
	}

	if f.FirstNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") > ?")
		values = append(values, f.FirstNameMaxGt)
	}

	if f.FirstNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMinLt)
	}

	if f.FirstNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") < ?")
		values = append(values, f.FirstNameMaxLt)
	}

	if f.FirstNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMinGte)
	}

	if f.FirstNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") >= ?")
		values = append(values, f.FirstNameMaxGte)
	}

	if f.FirstNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMinLte)
	}

	if f.FirstNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") <= ?")
		values = append(values, f.FirstNameMaxLte)
	}

	if f.FirstNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMinIn)
	}

	if f.FirstNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") IN (?)")
		values = append(values, f.FirstNameMaxIn)
	}

	if f.FirstNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.FirstNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.FirstNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMinPrefix))
	}

	if f.FirstNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.FirstNameMaxPrefix))
	}

	if f.FirstNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMinSuffix))
	}

	if f.FirstNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("firstName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.FirstNameMaxSuffix))
	}

	if f.LastNameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMin)
	}

	if f.LastNameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") = ?")
		values = append(values, f.LastNameMax)
	}

	if f.LastNameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMinNe)
	}

	if f.LastNameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") != ?")
		values = append(values, f.LastNameMaxNe)
	}

	if f.LastNameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMinGt)
	}

	if f.LastNameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") > ?")
		values = append(values, f.LastNameMaxGt)
	}

	if f.LastNameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMinLt)
	}

	if f.LastNameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") < ?")
		values = append(values, f.LastNameMaxLt)
	}

	if f.LastNameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMinGte)
	}

	if f.LastNameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") >= ?")
		values = append(values, f.LastNameMaxGte)
	}

	if f.LastNameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMinLte)
	}

	if f.LastNameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") <= ?")
		values = append(values, f.LastNameMaxLte)
	}

	if f.LastNameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMinIn)
	}

	if f.LastNameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") IN (?)")
		values = append(values, f.LastNameMaxIn)
	}

	if f.LastNameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.LastNameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.LastNameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMinPrefix))
	}

	if f.LastNameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.LastNameMaxPrefix))
	}

	if f.LastNameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMinSuffix))
	}

	if f.LastNameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("lastName")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.LastNameMaxSuffix))
	}

	if f.AddressRawMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") = ?")
		values = append(values, f.AddressRawMin)
	}

	if f.AddressRawMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") = ?")
		values = append(values, f.AddressRawMax)
	}

	if f.AddressRawMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") != ?")
		values = append(values, f.AddressRawMinNe)
	}

	if f.AddressRawMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") != ?")
		values = append(values, f.AddressRawMaxNe)
	}

	if f.AddressRawMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") > ?")
		values = append(values, f.AddressRawMinGt)
	}

	if f.AddressRawMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") > ?")
		values = append(values, f.AddressRawMaxGt)
	}

	if f.AddressRawMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") < ?")
		values = append(values, f.AddressRawMinLt)
	}

	if f.AddressRawMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") < ?")
		values = append(values, f.AddressRawMaxLt)
	}

	if f.AddressRawMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") >= ?")
		values = append(values, f.AddressRawMinGte)
	}

	if f.AddressRawMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") >= ?")
		values = append(values, f.AddressRawMaxGte)
	}

	if f.AddressRawMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") <= ?")
		values = append(values, f.AddressRawMinLte)
	}

	if f.AddressRawMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") <= ?")
		values = append(values, f.AddressRawMaxLte)
	}

	if f.AddressRawMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") IN (?)")
		values = append(values, f.AddressRawMinIn)
	}

	if f.AddressRawMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") IN (?)")
		values = append(values, f.AddressRawMaxIn)
	}

	if f.AddressRawMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AddressRawMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AddressRawMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.AddressRawMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.AddressRawMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AddressRawMinPrefix))
	}

	if f.AddressRawMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.AddressRawMaxPrefix))
	}

	if f.AddressRawMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("addressRaw")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AddressRawMinSuffix))
	}

	if f.AddressRawMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("addressRaw")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.AddressRawMaxSuffix))
	}

	if f.SalaryMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("salary")+") = ?")
		values = append(values, f.SalaryMin)
	}

	if f.SalaryMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("salary")+") = ?")
		values = append(values, f.SalaryMax)
	}

	if f.SalaryAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("salary")+") = ?")
		values = append(values, f.SalaryAvg)
	}

	if f.SalaryMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("salary")+") != ?")
		values = append(values, f.SalaryMinNe)
	}

	if f.SalaryMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("salary")+") != ?")
		values = append(values, f.SalaryMaxNe)
	}

	if f.SalaryAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("salary")+") != ?")
		values = append(values, f.SalaryAvgNe)
	}

	if f.SalaryMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("salary")+") > ?")
		values = append(values, f.SalaryMinGt)
	}

	if f.SalaryMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("salary")+") > ?")
		values = append(values, f.SalaryMaxGt)
	}

	if f.SalaryAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("salary")+") > ?")
		values = append(values, f.SalaryAvgGt)
	}

	if f.SalaryMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("salary")+") < ?")
		values = append(values, f.SalaryMinLt)
	}

	if f.SalaryMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("salary")+") < ?")
		values = append(values, f.SalaryMaxLt)
	}

	if f.SalaryAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("salary")+") < ?")
		values = append(values, f.SalaryAvgLt)
	}

	if f.SalaryMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("salary")+") >= ?")
		values = append(values, f.SalaryMinGte)
	}

	if f.SalaryMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("salary")+") >= ?")
		values = append(values, f.SalaryMaxGte)
	}

	if f.SalaryAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("salary")+") >= ?")
		values = append(values, f.SalaryAvgGte)
	}

	if f.SalaryMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("salary")+") <= ?")
		values = append(values, f.SalaryMinLte)
	}

	if f.SalaryMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("salary")+") <= ?")
		values = append(values, f.SalaryMaxLte)
	}

	if f.SalaryAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("salary")+") <= ?")
		values = append(values, f.SalaryAvgLte)
	}

	if f.SalaryMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("salary")+") IN (?)")
		values = append(values, f.SalaryMinIn)
	}

	if f.SalaryMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("salary")+") IN (?)")
		values = append(values, f.SalaryMaxIn)
	}

	if f.SalaryAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("salary")+") IN (?)")
		values = append(values, f.SalaryAvgIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *UserFilterType) AndWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *UserFilterType) OrWith(f2 ...*UserFilterType) *UserFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &UserFilterType{
		Or: append(_f2, f),
	}
}

// IsEmpty ...
func (f *PlainEntityFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}

// Apply method
func (f *PlainEntityFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("plain_entities"), wheres, whereValues, havings, havingValues, joins)
}

// ApplyWithAlias method
func (f *PlainEntityFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	return nil
}

// WhereContent ...
func (f *PlainEntityFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Date != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" = ?")
		values = append(values, f.Date)
	}

	if f.DateNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" != ?")
		values = append(values, f.DateNe)
	}

	if f.DateGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" > ?")
		values = append(values, f.DateGt)
	}

	if f.DateLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" < ?")
		values = append(values, f.DateLt)
	}

	if f.DateGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" >= ?")
		values = append(values, f.DateGte)
	}

	if f.DateLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" <= ?")
		values = append(values, f.DateLte)
	}

	if f.DateIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IN (?)")
		values = append(values, f.DateIn)
	}

	if f.DateNull != nil {
		if *f.DateNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("date")+" IS NOT NULL")
		}
	}

	if f.Text != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" = ?")
		values = append(values, f.Text)
	}

	if f.TextNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" != ?")
		values = append(values, f.TextNe)
	}

	if f.TextGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" > ?")
		values = append(values, f.TextGt)
	}

	if f.TextLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" < ?")
		values = append(values, f.TextLt)
	}

	if f.TextGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" >= ?")
		values = append(values, f.TextGte)
	}

	if f.TextLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" <= ?")
		values = append(values, f.TextLte)
	}

	if f.TextIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" IN (?)")
		values = append(values, f.TextIn)
	}

	if f.TextLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TextLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TextPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TextPrefix))
	}

	if f.TextSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TextSuffix))
	}

	if f.TextNull != nil {
		if *f.TextNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("text")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}

// HavingContent method
func (f *PlainEntityFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.DateMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") = ?")
		values = append(values, f.DateMin)
	}

	if f.DateMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") = ?")
		values = append(values, f.DateMax)
	}

	if f.DateMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") != ?")
		values = append(values, f.DateMinNe)
	}

	if f.DateMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") != ?")
		values = append(values, f.DateMaxNe)
	}

	if f.DateMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") > ?")
		values = append(values, f.DateMinGt)
	}

	if f.DateMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") > ?")
		values = append(values, f.DateMaxGt)
	}

	if f.DateMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") < ?")
		values = append(values, f.DateMinLt)
	}

	if f.DateMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") < ?")
		values = append(values, f.DateMaxLt)
	}

	if f.DateMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") >= ?")
		values = append(values, f.DateMinGte)
	}

	if f.DateMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") >= ?")
		values = append(values, f.DateMaxGte)
	}

	if f.DateMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") <= ?")
		values = append(values, f.DateMinLte)
	}

	if f.DateMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") <= ?")
		values = append(values, f.DateMaxLte)
	}

	if f.DateMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("date")+") IN (?)")
		values = append(values, f.DateMinIn)
	}

	if f.DateMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("date")+") IN (?)")
		values = append(values, f.DateMaxIn)
	}

	if f.TextMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") = ?")
		values = append(values, f.TextMin)
	}

	if f.TextMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") = ?")
		values = append(values, f.TextMax)
	}

	if f.TextMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") != ?")
		values = append(values, f.TextMinNe)
	}

	if f.TextMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") != ?")
		values = append(values, f.TextMaxNe)
	}

	if f.TextMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") > ?")
		values = append(values, f.TextMinGt)
	}

	if f.TextMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") > ?")
		values = append(values, f.TextMaxGt)
	}

	if f.TextMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") < ?")
		values = append(values, f.TextMinLt)
	}

	if f.TextMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") < ?")
		values = append(values, f.TextMaxLt)
	}

	if f.TextMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") >= ?")
		values = append(values, f.TextMinGte)
	}

	if f.TextMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") >= ?")
		values = append(values, f.TextMaxGte)
	}

	if f.TextMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") <= ?")
		values = append(values, f.TextMinLte)
	}

	if f.TextMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") <= ?")
		values = append(values, f.TextMaxLte)
	}

	if f.TextMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") IN (?)")
		values = append(values, f.TextMinIn)
	}

	if f.TextMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") IN (?)")
		values = append(values, f.TextMaxIn)
	}

	if f.TextMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TextMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TextMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.TextMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.TextMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TextMinPrefix))
	}

	if f.TextMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.TextMaxPrefix))
	}

	if f.TextMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TextMinSuffix))
	}

	if f.TextMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("text")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.TextMaxSuffix))
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *PlainEntityFilterType) AndWith(f2 ...*PlainEntityFilterType) *PlainEntityFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PlainEntityFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *PlainEntityFilterType) OrWith(f2 ...*PlainEntityFilterType) *PlainEntityFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &PlainEntityFilterType{
		Or: append(_f2, f),
	}
}
