package usersrepo

import (
	"fmt"

	"github.com/getchipman/bolt-api/app/context"
	"github.com/getchipman/bolt-api/app/core/domains"
	"github.com/getchipman/bolt-api/app/core/ports"
)

// UserRepository contains objects for database communication.
type UserRepository struct {
	Db    ports.Persistence
	Cache ports.Cache
}

//New create a User repository.
func New(db ports.Persistence, cache ports.Cache) *UserRepository {
	return &UserRepository{
		Db:    db,
		Cache: cache,
	}
}

// Create save a new user.
func (r *UserRepository) Create(ctx *context.Context, user *domains.User) context.Result {
	query := fmt.Sprintf(`
		INSERT INTO users(name, email, client_id, client_secret, create_at, update_at)
		VALUES ('%v', '%v', '%v', '%v', NOW(), NOW())`, user.Name, user.Email, user.ClientID, user.ClientSecret)

	_, err := r.Db.Exec(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to insert user - Query: %v", query)
		return ctx.ResultError(1, err.Error())
	}

	return ctx.ResultSuccess()
}

// Find find a user in database.
func (r *UserRepository) Find(ctx *context.Context, user *domains.User) (context.Result, string, string) {
	query := fmt.Sprintf(`
	SELECT 
		email, client_id 
	FROM 
		users 
	WHERE 
		email = '%v' 
	OR 
		client_id = '%v'
	`, user.Email, user.ClientID)

	rows, err := r.Db.QueryRow(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get user - Email: %v, ClientID: %v", user.Email, user.ClientID)
		return ctx.ResultError(1, err.Error()), "", ""
	}

	defer rows.Close()
	var email string
	var clientID string

	for rows.Next() {
		err := rows.Scan(&email, &clientID)
		if err != nil {
			ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get user - Email: %v, ClientID: %v", user.Email, user.ClientID)
			return ctx.ResultError(2, err.Error()), "", ""
		}
	}

	return ctx.ResultSuccess(), email, clientID
}
