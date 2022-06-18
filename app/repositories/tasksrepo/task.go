package tasksrepo

import (
	"fmt"

	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/frhdl/bolt-api/app/core/ports"
)

// TaskRepository contains objects for database communication.
type TaskRepository struct {
	Db    ports.Persistence
	Cache ports.Cache
}

// New create task repository.
func New(db ports.Persistence, cache ports.Cache) *TaskRepository {
	return &TaskRepository{
		Db:    db,
		Cache: cache,
	}
}

// Create save a new task.
func (r *TaskRepository) Create(ctx *context.Context, task *domains.Task) context.Result {
	query := fmt.Sprintf(`
	INSERT INTO tasks(description, user_id, project_id, create_at, finish_at, done)
	VALUES ('%v', %v, %v, NOW(), NOW(), false)
	`, task.Description, ctx.LoggedUserID, task.ProjectID)

	_, err := r.Db.Exec(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to insert task - Query: %v", query)
		return ctx.ResultError(1, err.Error())
	}

	return ctx.ResultSuccess()
}

// GetAll list all tasks in the target project.
func (r *TaskRepository) GetAll(ctx *context.Context, projectID int, page int, limit int) (context.Result, []*domains.Task) {
	page = (page - 1) * limit

	query := fmt.Sprintf(`
	SELECT 
		id, description, user_id, project_id, create_at, finish_at, done
	FROM
		tasks
	WHERE
		project_id = %v
	AND
		user_id = %v
	OFFSET %v
	LIMIT %v	
	`, projectID, ctx.LoggedUserID, page, limit)

	rows, err := r.Db.QueryRow(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get tasks - Query: %v", query)
		return ctx.ResultError(1, err.Error()), nil
	}

	defer rows.Close()
	var tasks []*domains.Task

	for rows.Next() {
		var task domains.Task

		err := rows.Scan(&task.ID, &task.Description, &task.ProjectID, &task.UserID, &task.CreateAt, &task.FinishAt, &task.Done)
		if err != nil {
			ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get task - ProjectID %v", projectID)
			return ctx.ResultError(2, err.Error()), nil
		}

		tasks = append(tasks, &task)
	}

	if tasks == nil {
		return ctx.ResultSuccess(), []*domains.Task{}
	}

	return ctx.ResultSuccess(), tasks
}

// Update save a new data in the existing task.
func (r *TaskRepository) Update(ctx *context.Context, task *domains.Task) context.Result {
	query := fmt.Sprintf(`
	UPDATE tasks SET description = '%v'`, task.Description)

	if task.Done {
		query += ` , finish_at = NOW(), done = true `
	}

	query += fmt.Sprintf(`
	WHERE id = %v AND user_id = %v`, task.ID, ctx.LoggedUserID)

	_, err := r.Db.Exec(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to update task - Query: %v", query)
		return ctx.ResultError(1, err.Error())
	}

	return ctx.ResultSuccess()
}

// Delete delete an existing task.
func (r *TaskRepository) Delete(ctx *context.Context, taskID int) context.Result {
	query := fmt.Sprintf(`
	DELETE FROM
		tasks
	WHERE
		id = %v	
	AND
		user_id = %v
	`, taskID, ctx.LoggedUserID)

	_, err := r.Db.Exec(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to delete task - Query: %v", query)
		return ctx.ResultError(1, err.Error())
	}

	return ctx.ResultSuccess()
}
