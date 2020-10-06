package gen

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/loopcontext/go-graphql-orm/events"
)

// GeneratedMutationResolver struct
type GeneratedMutationResolver struct{ *GeneratedResolver }

// MutationEvents struct
type MutationEvents struct {
	Events []events.Event
}

// EnrichContextWithMutations method
func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}

// FinishMutationContext method
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}

// RollbackMutationContext method
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}

// GetMutationEventStore method
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}

// AddMutationEvent method
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

// CreateTask method
func (r *GeneratedMutationResolver) CreateTask(ctx context.Context, input map[string]interface{}) (item *Task, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateTask(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateTaskHandler handler
func CreateTaskHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Task, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Task{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Task",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		item.Title = changes.Title

		event.AddNewValue("title", changes.Title)
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		item.Completed = changes.Completed

		event.AddNewValue("completed", changes.Completed)
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		item.State = changes.State

		event.AddNewValue("state", changes.State)
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		item.DueDate = changes.DueDate

		event.AddNewValue("dueDate", changes.DueDate)
	}

	if _, ok := input["metas"]; ok {
		_value, _err := json.Marshal(changes.Metas)
		if _err != nil {
			err = _err
			return
		}
		strval := string(_value)
		value := &strval
		if item.Metas != value && (item.Metas == nil || value == nil || *item.Metas != *value) {
			item.Metas = value
			event.AddNewValue("metas", value)
		}
	}

	if _, ok := input["meta"]; ok {
		_value, _err := json.Marshal(changes.Meta)
		if _err != nil {
			err = _err
			return
		}
		strval := string(_value)
		value := &strval
		if item.Meta != value && (item.Meta == nil || value == nil || *item.Meta != *value) {
			item.Meta = value
			event.AddNewValue("meta", value)
		}
	}

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {
		item.AssigneeID = changes.AssigneeID

		event.AddNewValue("assigneeId", changes.AssigneeID)
	}

	if _, ok := input["ownerId"]; ok && (item.OwnerID != changes.OwnerID) && (item.OwnerID == nil || changes.OwnerID == nil || *item.OwnerID != *changes.OwnerID) {
		item.OwnerID = changes.OwnerID

		event.AddNewValue("ownerId", changes.OwnerID)
	}

	if _, ok := input["parentTaskId"]; ok && (item.ParentTaskID != changes.ParentTaskID) && (item.ParentTaskID == nil || changes.ParentTaskID == nil || *item.ParentTaskID != *changes.ParentTaskID) {
		item.ParentTaskID = changes.ParentTaskID

		event.AddNewValue("parentTaskId", changes.ParentTaskID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["subtasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Subtasks")
		association.Replace(items)
	}

	if ids, exists := input["categoriesIds"]; exists {
		items := []TaskCategory{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Categories")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateTask method
func (r *GeneratedMutationResolver) UpdateTask(ctx context.Context, id string, input map[string]interface{}) (item *Task, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateTask(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateTaskHandler handler
func UpdateTaskHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Task, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Task",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes TaskChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["title"]; ok && (item.Title != changes.Title) && (item.Title == nil || changes.Title == nil || *item.Title != *changes.Title) {
		event.AddOldValue("title", item.Title)
		event.AddNewValue("title", changes.Title)
		item.Title = changes.Title
	}

	if _, ok := input["completed"]; ok && (item.Completed != changes.Completed) && (item.Completed == nil || changes.Completed == nil || *item.Completed != *changes.Completed) {
		event.AddOldValue("completed", item.Completed)
		event.AddNewValue("completed", changes.Completed)
		item.Completed = changes.Completed
	}

	if _, ok := input["state"]; ok && (item.State != changes.State) && (item.State == nil || changes.State == nil || *item.State != *changes.State) {
		event.AddOldValue("state", item.State)
		event.AddNewValue("state", changes.State)
		item.State = changes.State
	}

	if _, ok := input["dueDate"]; ok && (item.DueDate != changes.DueDate) && (item.DueDate == nil || changes.DueDate == nil || *item.DueDate != *changes.DueDate) {
		event.AddOldValue("dueDate", item.DueDate)
		event.AddNewValue("dueDate", changes.DueDate)
		item.DueDate = changes.DueDate
	}

	if _, ok := input["metas"]; ok {
		_value, _err := json.Marshal(changes.Metas)
		if _err != nil {
			err = _err
			return
		}
		if _err != nil {
			err = _err
			return
		}
		strval := string(_value)
		value := &strval
		if item.Metas != value && (item.Metas == nil || value == nil || *item.Metas != *value) {
			event.AddOldValue("metas", item.Metas)
			event.AddNewValue("metas", value)
			item.Metas = value
		}
	}

	if _, ok := input["meta"]; ok {
		_value, _err := json.Marshal(changes.Meta)
		if _err != nil {
			err = _err
			return
		}
		if _err != nil {
			err = _err
			return
		}
		strval := string(_value)
		value := &strval
		if item.Meta != value && (item.Meta == nil || value == nil || *item.Meta != *value) {
			event.AddOldValue("meta", item.Meta)
			event.AddNewValue("meta", value)
			item.Meta = value
		}
	}

	if _, ok := input["assigneeId"]; ok && (item.AssigneeID != changes.AssigneeID) && (item.AssigneeID == nil || changes.AssigneeID == nil || *item.AssigneeID != *changes.AssigneeID) {
		event.AddOldValue("assigneeId", item.AssigneeID)
		event.AddNewValue("assigneeId", changes.AssigneeID)
		item.AssigneeID = changes.AssigneeID
	}

	if _, ok := input["ownerId"]; ok && (item.OwnerID != changes.OwnerID) && (item.OwnerID == nil || changes.OwnerID == nil || *item.OwnerID != *changes.OwnerID) {
		event.AddOldValue("ownerId", item.OwnerID)
		event.AddNewValue("ownerId", changes.OwnerID)
		item.OwnerID = changes.OwnerID
	}

	if _, ok := input["parentTaskId"]; ok && (item.ParentTaskID != changes.ParentTaskID) && (item.ParentTaskID == nil || changes.ParentTaskID == nil || *item.ParentTaskID != *changes.ParentTaskID) {
		event.AddOldValue("parentTaskId", item.ParentTaskID)
		event.AddNewValue("parentTaskId", changes.ParentTaskID)
		item.ParentTaskID = changes.ParentTaskID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["subtasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Subtasks")
		association.Replace(items)
	}

	if ids, exists := input["categoriesIds"]; exists {
		items := []TaskCategory{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Categories")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteTask method
func (r *GeneratedMutationResolver) DeleteTask(ctx context.Context, id string) (item *Task, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteTask(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteTaskHandler handler
func DeleteTaskHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Task, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Task{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Task",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("tasks")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllTasks method
func (r *GeneratedMutationResolver) DeleteAllTasks(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllTasks(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllTasksHandler handler
func DeleteAllTasksHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Task{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateTaskCategory method
func (r *GeneratedMutationResolver) CreateTaskCategory(ctx context.Context, input map[string]interface{}) (item *TaskCategory, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateTaskCategory(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateTaskCategoryHandler handler
func CreateTaskCategoryHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *TaskCategory, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &TaskCategory{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "TaskCategory",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes TaskCategoryChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["tasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateTaskCategory method
func (r *GeneratedMutationResolver) UpdateTaskCategory(ctx context.Context, id string, input map[string]interface{}) (item *TaskCategory, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateTaskCategory(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateTaskCategoryHandler handler
func UpdateTaskCategoryHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *TaskCategory, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &TaskCategory{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "TaskCategory",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes TaskCategoryChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["tasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteTaskCategory method
func (r *GeneratedMutationResolver) DeleteTaskCategory(ctx context.Context, id string) (item *TaskCategory, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteTaskCategory(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteTaskCategoryHandler handler
func DeleteTaskCategoryHandler(ctx context.Context, r *GeneratedResolver, id string) (item *TaskCategory, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &TaskCategory{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "TaskCategory",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("task_categories")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllTaskCategories method
func (r *GeneratedMutationResolver) DeleteAllTaskCategories(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllTaskCategories(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllTaskCategoriesHandler handler
func DeleteAllTaskCategoriesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&TaskCategory{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateCompany method
func (r *GeneratedMutationResolver) CreateCompany(ctx context.Context, input map[string]interface{}) (item *Company, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateCompany(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateCompanyHandler handler
func CreateCompanyHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &Company{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "Company",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CompanyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["countryId"]; ok && (item.CountryID != changes.CountryID) && (item.CountryID == nil || changes.CountryID == nil || *item.CountryID != *changes.CountryID) {
		item.CountryID = changes.CountryID

		event.AddNewValue("countryId", changes.CountryID)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["employeesIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateCompany method
func (r *GeneratedMutationResolver) UpdateCompany(ctx context.Context, id string, input map[string]interface{}) (item *Company, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateCompany(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateCompanyHandler handler
func UpdateCompanyHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Company{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "Company",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes CompanyChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["countryId"]; ok && (item.CountryID != changes.CountryID) && (item.CountryID == nil || changes.CountryID == nil || *item.CountryID != *changes.CountryID) {
		event.AddOldValue("countryId", item.CountryID)
		event.AddNewValue("countryId", changes.CountryID)
		item.CountryID = changes.CountryID
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["employeesIds"]; exists {
		items := []User{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employees")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteCompany method
func (r *GeneratedMutationResolver) DeleteCompany(ctx context.Context, id string) (item *Company, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteCompany(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteCompanyHandler handler
func DeleteCompanyHandler(ctx context.Context, r *GeneratedResolver, id string) (item *Company, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &Company{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "Company",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("companies")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllCompanies method
func (r *GeneratedMutationResolver) DeleteAllCompanies(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllCompanies(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllCompaniesHandler handler
func DeleteAllCompaniesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&Company{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreateUser method
func (r *GeneratedMutationResolver) CreateUser(ctx context.Context, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateUser(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreateUserHandler handler
func CreateUserHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &User{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "User",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["code"]; ok && (item.Code != changes.Code) && (item.Code == nil || changes.Code == nil || *item.Code != *changes.Code) {
		item.Code = changes.Code

		event.AddNewValue("code", changes.Code)
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		item.Email = changes.Email

		event.AddNewValue("email", changes.Email)
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		item.FirstName = changes.FirstName

		event.AddNewValue("firstName", changes.FirstName)
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		item.LastName = changes.LastName

		event.AddNewValue("lastName", changes.LastName)
	}

	if _, ok := input["addressRaw"]; ok && (item.AddressRaw != changes.AddressRaw) && (item.AddressRaw == nil || changes.AddressRaw == nil || *item.AddressRaw != *changes.AddressRaw) {
		item.AddressRaw = changes.AddressRaw

		event.AddNewValue("addressRaw", changes.AddressRaw)
	}

	if _, ok := input["salary"]; ok && (item.Salary != changes.Salary) && (item.Salary == nil || changes.Salary == nil || *item.Salary != *changes.Salary) {
		item.Salary = changes.Salary

		event.AddNewValue("salary", changes.Salary)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["employersIds"]; exists {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employers")
		association.Replace(items)
	}

	if ids, exists := input["tasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	if ids, exists := input["createdTasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("CreatedTasks")
		association.Replace(items)
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdateUser method
func (r *GeneratedMutationResolver) UpdateUser(ctx context.Context, id string, input map[string]interface{}) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateUser(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdateUserHandler handler
func UpdateUserHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &User{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "User",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes UserChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["code"]; ok && (item.Code != changes.Code) && (item.Code == nil || changes.Code == nil || *item.Code != *changes.Code) {
		event.AddOldValue("code", item.Code)
		event.AddNewValue("code", changes.Code)
		item.Code = changes.Code
	}

	if _, ok := input["email"]; ok && (item.Email != changes.Email) && (item.Email == nil || changes.Email == nil || *item.Email != *changes.Email) {
		event.AddOldValue("email", item.Email)
		event.AddNewValue("email", changes.Email)
		item.Email = changes.Email
	}

	if _, ok := input["firstName"]; ok && (item.FirstName != changes.FirstName) && (item.FirstName == nil || changes.FirstName == nil || *item.FirstName != *changes.FirstName) {
		event.AddOldValue("firstName", item.FirstName)
		event.AddNewValue("firstName", changes.FirstName)
		item.FirstName = changes.FirstName
	}

	if _, ok := input["lastName"]; ok && (item.LastName != changes.LastName) && (item.LastName == nil || changes.LastName == nil || *item.LastName != *changes.LastName) {
		event.AddOldValue("lastName", item.LastName)
		event.AddNewValue("lastName", changes.LastName)
		item.LastName = changes.LastName
	}

	if _, ok := input["addressRaw"]; ok && (item.AddressRaw != changes.AddressRaw) && (item.AddressRaw == nil || changes.AddressRaw == nil || *item.AddressRaw != *changes.AddressRaw) {
		event.AddOldValue("addressRaw", item.AddressRaw)
		event.AddNewValue("addressRaw", changes.AddressRaw)
		item.AddressRaw = changes.AddressRaw
	}

	if _, ok := input["salary"]; ok && (item.Salary != changes.Salary) && (item.Salary == nil || changes.Salary == nil || *item.Salary != *changes.Salary) {
		event.AddOldValue("salary", item.Salary)
		event.AddNewValue("salary", changes.Salary)
		item.Salary = changes.Salary
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if ids, exists := input["employersIds"]; exists {
		items := []Company{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Employers")
		association.Replace(items)
	}

	if ids, exists := input["tasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("Tasks")
		association.Replace(items)
	}

	if ids, exists := input["createdTasksIds"]; exists {
		items := []Task{}
		tx.Find(&items, "id IN (?)", ids)
		association := tx.Model(&item).Association("CreatedTasks")
		association.Replace(items)
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeleteUser method
func (r *GeneratedMutationResolver) DeleteUser(ctx context.Context, id string) (item *User, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteUser(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeleteUserHandler handler
func DeleteUserHandler(ctx context.Context, r *GeneratedResolver, id string) (item *User, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &User{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "User",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("users")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllUsers method
func (r *GeneratedMutationResolver) DeleteAllUsers(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllUsers(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllUsersHandler handler
func DeleteAllUsersHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&User{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}

// CreatePlainEntity method
func (r *GeneratedMutationResolver) CreatePlainEntity(ctx context.Context, input map[string]interface{}) (item *PlainEntity, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreatePlainEntity(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// CreatePlainEntityHandler handler
func CreatePlainEntityHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *PlainEntity, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &PlainEntity{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "PlainEntity",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PlainEntityChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["date"]; ok && (item.Date != changes.Date) && (item.Date == nil || changes.Date == nil || *item.Date != *changes.Date) {
		item.Date = changes.Date

		event.AddNewValue("date", changes.Date)
	}

	if _, ok := input["text"]; ok && (item.Text != changes.Text) && (item.Text == nil || changes.Text == nil || *item.Text != *changes.Text) {
		item.Text = changes.Text

		event.AddNewValue("text", changes.Text)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// UpdatePlainEntity method
func (r *GeneratedMutationResolver) UpdatePlainEntity(ctx context.Context, id string, input map[string]interface{}) (item *PlainEntity, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdatePlainEntity(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// UpdatePlainEntityHandler handler
func UpdatePlainEntityHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *PlainEntity, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &PlainEntity{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "PlainEntity",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes PlainEntityChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["date"]; ok && (item.Date != changes.Date) && (item.Date == nil || changes.Date == nil || *item.Date != *changes.Date) {
		event.AddOldValue("date", item.Date)
		event.AddNewValue("date", changes.Date)
		item.Date = changes.Date
	}

	if _, ok := input["text"]; ok && (item.Text != changes.Text) && (item.Text == nil || changes.Text == nil || *item.Text != *changes.Text) {
		event.AddOldValue("text", item.Text)
		event.AddNewValue("text", changes.Text)
		item.Text = changes.Text
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}

// DeletePlainEntity method
func (r *GeneratedMutationResolver) DeletePlainEntity(ctx context.Context, id string) (item *PlainEntity, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeletePlainEntity(ctx, r.GeneratedResolver, id)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}

// DeletePlainEntityHandler handler
func DeletePlainEntityHandler(ctx context.Context, r *GeneratedResolver, id string) (item *PlainEntity, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &PlainEntity{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "PlainEntity",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("plain_entities")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}

// DeleteAllPlainEntities method
func (r *GeneratedMutationResolver) DeleteAllPlainEntities(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllPlainEntities(ctx, r.GeneratedResolver)
	if err != nil {
		errRMC := RollbackMutationContext(ctx, r.GeneratedResolver)
		if errRMC != nil {
			err = fmt.Errorf("[Wrapped]: RollbackMutationContext error: %w\n[Original]: %q", errRMC, err)
		}
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}

// DeleteAllPlainEntitiesHandler handler
func DeleteAllPlainEntitiesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&PlainEntity{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
