package services

import (
	"context"

	"github.com/klabhisheky/user_service/model"
)

type User interface {
	Find(ctx context.Context, userId int) (*model.User, error)
	Create(ctx context.Context, model *model.User) (*model.User, error)
}
