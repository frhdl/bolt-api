package projectsrepo

import (
	"fmt"

	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/frhdl/bolt-api/app/core/ports"
)

// ProjectRepository contains objects for database communication.
type ProjectRepository struct {
	Db    ports.Persistence
	Cache ports.Cache
}

// New create project repository.
func New(db ports.Persistence, cache ports.Cache) *ProjectRepository {
	return &ProjectRepository{
		Db:    db,
		Cache: cache,
	}
}

// Create save a new project.
func (r *ProjectRepository) Create(ctx *context.Context, project *domains.Project) context.Result {
	query := fmt.Sprintf(`
	INSERT INTO projects(name, user_id, create_at) 
	VALUES ('%v', %v, NOW())
	`, project.Name, ctx.LoggedUserID)

	_, err := r.Db.Exec(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to insert project - Query: %v", query)
		return ctx.ResultError(1, err.Error())
	}

	return ctx.ResultSuccess()
}

// GetAll list all projects in the target user.
func (r *ProjectRepository) GetAll(ctx *context.Context, page int, limit int) (context.Result, []*domains.Project) {
	page = (page - 1) * limit

	query := fmt.Sprintf(`
	SELECT 
		id, name, user_id
	FROM 
		projects 
	WHERE 
		user_id = %v 
	OFFSET %v 
	LIMIT %v 
	`, ctx.LoggedUserID, page, limit)

	rows, err := r.Db.QueryRow(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get projects - Query: %v", query)
		return ctx.ResultError(1, err.Error()), nil
	}

	defer rows.Close()
	var projects []*domains.Project

	for rows.Next() {
		var project domains.Project

		err := rows.Scan(&project.ID, &project.Name, &project.UserID)
		if err != nil {
			ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get project - UserID %v", ctx.LoggedUserID)
			return ctx.ResultError(2, err.Error()), nil
		}

		projects = append(projects, &project)
	}

	if projects == nil {
		return ctx.ResultSuccess(), []*domains.Project{}
	}

	return ctx.ResultSuccess(), projects
}

// Update save a new data in the existing project.
func (r *ProjectRepository) Update(ctx *context.Context, project *domains.Project) context.Result {
	query := fmt.Sprintf(`
		UPDATE
			projects
		SET
			name = '%v'
		WHERE 
			id = %v
		AND
			user_id = %v
	`, project.Name, project.ID, ctx.LoggedUserID)

	_, err := r.Db.Exec(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to update project - Query: %v", query)
		return ctx.ResultError(1, err.Error())
	}

	return ctx.ResultSuccess()
}

// Delete delete an existing project.
func (r *ProjectRepository) Delete(ctx *context.Context, projectID int) context.Result {
	query := fmt.Sprintf(`
	DELETE FROM
		projects
	WHERE
		id = %v
	AND
		user_id = %v
	`, projectID, ctx.LoggedUserID)

	_, err := r.Db.Exec(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to delete project - Query: %v", query)
		return ctx.ResultError(1, err.Error())
	}

	return ctx.ResultSuccess()
}

// Find find a projecet in database.
func (r *ProjectRepository) Find(ctx *context.Context, project *domains.Project) (context.Result, string) {
	query := fmt.Sprintf(`
	SELECT 
		name
	FROM
		projects
	WHERE 
		id = %v
	AND
		user_id = %v		
	`, project.ID, ctx.LoggedUserID)

	rows, err := r.Db.QueryRow(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get project - ProjectID: %v, UserID: %v", project.ID, ctx.LoggedUserID)
		return ctx.ResultError(1, err.Error()), ""
	}

	defer rows.Close()
	var name string

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get project - ProjectID: %v, UserID: %v", project.ID, ctx.LoggedUserID)
			return ctx.ResultError(1, err.Error()), ""
		}
	}

	return ctx.ResultSuccess(), name
}
