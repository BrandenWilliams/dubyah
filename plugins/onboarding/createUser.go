package plugin

import (
	"context"
	"fmt"

	"github.com/gdbu/jump/users"
)

func createUser(ctx context.Context, req CreateRequest) (userID string, err error) {
	defer func() {
		if err == nil {
			// No errors ocurred and no rollback needed, return
			return
		}

		// If error occurrs on user creation delete what was created
		deleteUser(context.Background(), userID)
	}()

	// Create user
	userID, _, err = p.Jump.CreateUser(req.Email, req.Password, req.GetGroups()...)
	switch err {
	case nil:
	case users.ErrEmailExists:
		err = fmt.Errorf("error creating user for <%s>, please contact customer support", req.Email)
		return
	default:
		return
	}

	return
}

// CreateRequest is the request used for new user onboarding
type CreateRequest struct {
	// Login credentials
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	RepeatPassword string `json:"repeatpassword" form:"repeatpassword"`
}

// GetGroups will get the groups for a given CreateRequest
func (c *CreateRequest) GetGroups() (groups []string) {
	// All users are both shoppers and merchants
	groups = []string{
		"customer",
		"tasker",
	}

	return
}
