package store

import (
	"context"

	"github.com/klabhisheky/user_service/model"
)

type User interface {
	Find(ctx context.Context, userId int) (*model.User, error)
	Create(ctx context.Context, usr *model.User) (*model.User, error)
}
