package user

import (
	"context"

	"developer.zopsmart.com/go/backend/zs/types"
	"github.com/klabhisheky/user_service/model"
	"github.com/klabhisheky/user_service/store"
)

type userService struct {
	usrstore store.User
}

func New(usrstr store.User) *userService {
	return &userService{usrstore: usrstr}
}

func (usrsvc *userService) Find(ctx context.Context, userId int) (*model.User, error) {
	strRes, err := usrsvc.usrstore.Find(ctx, userId)
	if err != nil {
		return nil, err
	}

	return strRes, nil
}

func (usrsvc *userService) Create(ctx context.Context, usr *model.User) (*model.User, error) {

	if usr.UserId != 0 { //TODO: create constant with NULL=0, somewhere in design
		return nil, types.Error{Code: 400, Message: "Unnecessary parameter provided UserId"}
	}

	strRes, err := usrsvc.usrstore.Create(ctx, usr)
	if err != nil {
		return nil, err
	}

	return strRes, nil
}
