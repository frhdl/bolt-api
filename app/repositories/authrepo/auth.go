package authrepo

import (
	"fmt"

	"github.com/frhdl/bolt-api/app/context"
	"github.com/frhdl/bolt-api/app/core/domains"
	"github.com/frhdl/bolt-api/app/core/ports"
)

// AuthRepository contains objects for database communication.
type AuthRepository struct {
	Db    ports.Persistence
	Cache ports.Cache
}

// New create auth repository.
func New(db ports.Persistence, cache ports.Cache) *AuthRepository {
	return &AuthRepository{
		Db:    db,
		Cache: cache,
	}
}

// Login return user name, user email and user id.
func (r *AuthRepository) Login(ctx *context.Context, user *domains.User) (context.Result, int, string, string) {
	query := fmt.Sprintf(`SELECT id, name, email from users WHERE client_id = '%v' and client_secret = '%v'`, user.ClientID, user.ClientSecret)

	rows, err := r.Db.QueryRow(query)
	if err != nil {
		ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get user - ClientID: %v", user.ClientID)
	}

	defer rows.Close()
	var id int
	var name string
	var email string

	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			ctx.Logger.WithField("Error", err.Error()).Errorf("Error to get user - ClientID: %v", user.ClientID)

			return ctx.ResultError(2, err.Error()), 0, "", ""
		}
	}

	return ctx.ResultSuccess(), id, name, email
}
