package plugin

import (
	"context"

	"github.com/hatchify/errors"
)

func deleteUser(ctx context.Context, userID string) (err error) {
	if userID == "" {
		// No user ID is available, so we cannot remove any services, return
		return
	}

	// Users task delete

	var errs errors.ErrorList
	return errs.Err()
}
